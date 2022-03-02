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

	"github.com/Notch-Technologies/wizy/client"
	"github.com/Notch-Technologies/wizy/cmd/server/channel"
	"github.com/Notch-Technologies/wizy/cmd/server/config"
	"github.com/Notch-Technologies/wizy/cmd/server/database"
	"github.com/Notch-Technologies/wizy/cmd/server/pb/organization"
	"github.com/Notch-Technologies/wizy/cmd/server/pb/peer"
	"github.com/Notch-Technologies/wizy/cmd/server/pb/session"
	"github.com/Notch-Technologies/wizy/cmd/server/pb/user"
	server "github.com/Notch-Technologies/wizy/cmd/server/server"
	"github.com/Notch-Technologies/wizy/paths"
	"github.com/Notch-Technologies/wizy/store"
	"github.com/Notch-Technologies/wizy/types/flagtype"
	"github.com/Notch-Technologies/wizy/version"
	"github.com/Notch-Technologies/wizy/wislog"
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
	port       uint16
	verbose    int
	domain     string
	certfile   string
	certkey    string
	version    bool
	logFile    string
	logLevel   string
	dev        bool
}

func main() {
	flag.StringVar(&args.configpath, "config", paths.DefaultServerConfigFile(), "path of server config file")
	flag.Var(flagtype.PortValue(&args.port, flagtype.DefaultGrpcServerPort), "port", "specify the port of the server")
	flag.IntVar(&args.verbose, "verbose", 0, "0 is the default value, 1 is a redundant message")
	flag.StringVar(&args.domain, "domain", "", "your domain")
	flag.StringVar(&args.certfile, "cert-file", "", "your cert")
	flag.StringVar(&args.certkey, "cert-key", "", "your cert key")
	flag.BoolVar(&args.version, "version", false, "print version")
	flag.StringVar(&args.logFile, "logfile", paths.DefaultServerLogFile(), "set logfile path")
	flag.StringVar(&args.logLevel, "loglevel", wislog.DebugLevelStr, "set log level")
	flag.BoolVar(&args.dev, "dev", true, "is dev")

	flag.Parse()
	if flag.NArg() > 0 {
		log.Fatalf("does not take non-flag arguments: %q.", flag.Args())
	}

	if args.version {
		fmt.Println(version.String())
		os.Exit(0)
	}

	// initialize wissy logger
	//
	err := wislog.InitWisLog(args.logLevel, args.logFile, args.dev)
	if err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}

	wislog := wislog.NewWisLog("server")

	// initialize sqlite database
	//
	db := database.NewSqlite(wislog)
	err = db.MigrationUp()
	if err != nil {
		log.Fatal(err)
	}

	// initialize file store
	//
	sfs, err := store.NewFileStore(paths.DefaultWissyServerStateFile(), wislog)
	if err != nil {
		log.Fatal(err)
	}

	ss := store.NewServerStore(sfs)
	err = ss.WritePrivateKey()
	if err != nil {
		log.Fatal(err)
	}

	// load sever config
	//
	cfg := config.LoadConfig(args.configpath, args.domain, args.certfile, args.certkey)

	// initialize auth0 client
	//
	auth0Client := client.NewAuth0Client()

	// initialize peer update manager
	//
	peerUpdateManager := channel.NewPeersUpdateManager()

	// initialize server
	//
	s := server.NewServer(db, cfg, ss, auth0Client, peerUpdateManager)

	// configuration grpc server option
	//
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

	middleware := server.NewMiddlware(auth0Client)

	opts = append(opts, grpc.KeepaliveEnforcementPolicy(kaep), grpc.KeepaliveParams(kasp), grpc.UnaryInterceptor(grpc_auth.UnaryServerInterceptor(middleware.Authenticate)))
	grpcServer := grpc.NewServer(opts...)

	// initialize grpc server services
	//
	peer.RegisterPeerServiceServer(grpcServer, s.PeerServerService)
	user.RegisterUserServiceServer(grpcServer, s.UserServerService)
	session.RegisterSessionServiceServer(grpcServer, s.SessionServerService)
	organization.RegisterOrganizationServiceServer(grpcServer, s.OrganizationServerService)

	log.Printf("starting server: localhost:%v", args.port)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", args.port))
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
	log.Println("terminate server.")
	grpcServer.Stop()
}
