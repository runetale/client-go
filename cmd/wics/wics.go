package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Notch-Technologies/wizy/cmd/wics/config"
	"github.com/Notch-Technologies/wizy/cmd/wics/proto"
	"github.com/Notch-Technologies/wizy/cmd/wics/server"
	"github.com/Notch-Technologies/wizy/cmd/wics/server/api"
	"github.com/Notch-Technologies/wizy/cmd/wics/server/redis"
	"github.com/Notch-Technologies/wizy/paths"
	"github.com/Notch-Technologies/wizy/store"
	"github.com/Notch-Technologies/wizy/types/flagtype"
	"github.com/Notch-Technologies/wizy/version"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

var args struct {
	configpath	string
	wicsport	uint16
	port    	uint16
	verbose   	int
	domain   	string
	certfile   	string
	certkey    	string
	version    	bool
}

func main() {
	flag.StringVar(&args.configpath, "config", paths.DefaultWicsConfigFile(), "path of wics config file")
	flag.Var(flagtype.PortValue(&args.wicsport, flagtype.DefaultWicsPort), "wics-port", "specify the port of the wics server")
	flag.Var(flagtype.PortValue(&args.port, flagtype.DefaultApiPort), "port", "specify the port of the http server")
	flag.IntVar(&args.verbose, "verbose", 0, "0 is the default value, 1 is a redundant message")
	flag.StringVar(&args.domain, "domain", "", "your domain")
	flag.StringVar(&args.certfile, "cert-file", "", "your cert")
	flag.StringVar(&args.certkey, "cert-key", "", "your cert key")
	flag.BoolVar(&args.version, "version", false, "print version")

	flag.Parse()
	if flag.NArg() > 0 {
		log.Fatalf("does not take non-flag arguments: %q.", flag.Args())
	}

	if args.version {
		fmt.Println(version.String())
		os.Exit(0)
	}

	// login to redis
	p := os.Getenv("REDIS_PASSWORD")
	redisClient := redis.NewRedisClient(p)

	// create account store
	account := redis.NewAccountStore(redisClient)

	// create wics server state file
	sfs, err := store.NewFileStore(paths.DefaultWicsServerStateFile())
	if err != nil {
		log.Fatal(err)
	}

	ss := store.NewServerStore(sfs)
	err = ss.WritePrivateKey()
	if err != nil {
		log.Fatal(err)
	}

	// create or open wics config file
	cfg := config.LoadConfig(args.configpath, args.domain, args.certfile, args.certkey)

	// initialize new wics server
	s, err := server.NewServer(cfg, account, ss, redisClient)
	if err != nil {
		log.Fatal(err)
	}

	// launch grpc server
	var opts []grpc.ServerOption
	kaep := keepalive.EnforcementPolicy{
		MinTime:             15 * time.Second,
		PermitWithoutStream: true,
	}

	kasp := keepalive.ServerParameters{
		MaxConnectionIdle:     15 * time.Second,
		MaxConnectionAgeGrace: 5 * time.Second,
		Time:                  5 * time.Second,
		Timeout:               2 * time.Second,
	}

	opts = append(opts, grpc.KeepaliveEnforcementPolicy(kaep), grpc.KeepaliveParams(kasp))
	grpcServer := grpc.NewServer(opts...)

	proto.RegisterPeerServiceServer(grpcServer, s.PeerServiceServer)
	proto.RegisterUserServiceServer(grpcServer, s.UserServiceServer)
	log.Printf("started wics server: localhost:%v", args.wicsport)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", args.wicsport))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Start API Server
	httpServer := api.NewHTTPServer(args.port)
	log.Printf("started http server: localhost:%v", args.port)

	reflection.Register(grpcServer)
	go func() {
		if err = grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve grpc server: %v.", err)
		}
	}()

	stop := make(chan struct{})
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c,
			os.Interrupt,
			syscall.SIGKILL,
			syscall.SIGTERM,
			syscall.SIGINT,
		)
		select {
		case <-c:
			close(stop)
		}
	}()
	<-stop
	log.Println("terminate wics server.")
	grpcServer.Stop()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	httpServer.Shutdown(ctx)
}
