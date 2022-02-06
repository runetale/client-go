package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Notch-Technologies/wizy/cmd/server/config"
	"github.com/Notch-Technologies/wizy/cmd/server/database"
	server "github.com/Notch-Technologies/wizy/cmd/server/grpc_server"
	"github.com/Notch-Technologies/wizy/cmd/server/pb/peer"
	"github.com/Notch-Technologies/wizy/cmd/server/pb/session"
	"github.com/Notch-Technologies/wizy/cmd/server/pb/user"
	"github.com/Notch-Technologies/wizy/paths"
	"github.com/Notch-Technologies/wizy/store"
	"github.com/Notch-Technologies/wizy/types/flagtype"
	"github.com/Notch-Technologies/wizy/version"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
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
	configpath string
	wicsport   uint16
	port       uint16
	verbose    int
	domain     string
	certfile   string
	certkey    string
	version    bool
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

	db := database.NewSqlite()

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
	s, err := server.NewServer(db, cfg, ss)
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

	middleware := server.NewMiddlware()

	opts = append(opts, grpc.KeepaliveEnforcementPolicy(kaep), grpc.KeepaliveParams(kasp), grpc.UnaryInterceptor(grpc_auth.UnaryServerInterceptor(middleware.Authenticate)),)
	grpcServer := grpc.NewServer(opts...)

	peer.RegisterPeerServiceServer(grpcServer, s.PeerServiceServer)
	user.RegisterUserServiceServer(grpcServer, s.UserServiceServer)
	session.RegisterSessionServiceServer(grpcServer, s.SessionServiceServer)
	log.Printf("started wics server: localhost:%v", args.wicsport)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", args.wicsport))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

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
}
