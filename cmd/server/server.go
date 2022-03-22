package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	client "github.com/Notch-Technologies/wizy/auth0"
	"github.com/Notch-Technologies/wizy/cert"
	"github.com/Notch-Technologies/wizy/cmd/server/channel"
	"github.com/Notch-Technologies/wizy/cmd/server/config"
	"github.com/Notch-Technologies/wizy/cmd/server/database"
	"github.com/Notch-Technologies/wizy/cmd/server/pb/admin_network"
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
	"google.golang.org/grpc/reflection"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

var serverArgs struct {
	// server
	configpath string
	port       uint16
	// tls
	letsencryptDir string
	domain         string
	certfile       string
	certkey        string
	// turn
	turnURL      string // format like this"turn:www.notchturn.net:3478"
	turnUserName string
	turnPassword string
	turnSecret   string
	ttl          string
	credentials  bool
	// stun
	stunURL      string // format like this"stun:www.notchturn.net:3478"
	stunUserName string
	stunPassword string
	// signal
	signalURL string // format "172.16.165.128:10000" or "yourdomain:10000"
	// log
	logFile  string
	logLevel string
	verbose  int

	version bool
	dev     bool
}

func main() {
	flag.StringVar(&serverArgs.configpath, "config", paths.DefaultServerConfigFile(), "path of server config file")
	flag.Var(flagtype.PortValue(&serverArgs.port, flagtype.DefaultGrpcServerPort), "port", "specify the port of the server")

	flag.StringVar(&serverArgs.letsencryptDir, "letsencypt-dir", paths.DefaultLetsEncryptDir(), "directory of letsencrypt")
	flag.StringVar(&serverArgs.domain, "domain", "", "your domain")
	flag.StringVar(&serverArgs.certfile, "cert-file", "", "your cert file")
	flag.StringVar(&serverArgs.certkey, "cert-key", "", "your cert key")

	flag.StringVar(&serverArgs.turnURL, "turn-url", "stun:www.notchturn.net:3478", "your turn url")
	flag.StringVar(&serverArgs.turnUserName, "turn-user", "", "your turn username")
	flag.StringVar(&serverArgs.turnPassword, "turn-pass", "", "your turn password")
	flag.StringVar(&serverArgs.turnSecret, "turn-secret", "", "your cert key")
	flag.StringVar(&serverArgs.ttl, "turn-ttl", "0", "ttl duration")
	flag.BoolVar(&serverArgs.credentials, "turn-credentials", false, "")

	flag.StringVar(&serverArgs.stunURL, "stun-url", "turn:www.notchturn.net:3478", "your stun url")
	flag.StringVar(&serverArgs.stunUserName, "stun-user", "", "your stun username")
	flag.StringVar(&serverArgs.stunPassword, "stun-pass", "", "your stun password")

	flag.StringVar(&serverArgs.signalURL, "signal-url", "172.16.165.128:10000", "your signal url")

	flag.StringVar(&serverArgs.logFile, "logfile", paths.DefaultServerLogFile(), "set logfile path")
	flag.StringVar(&serverArgs.logLevel, "loglevel", wislog.DebugLevelStr, "set log level")
	flag.IntVar(&serverArgs.verbose, "verbose", 0, "0 is the default value, 1 is a redundant message")

	flag.BoolVar(&serverArgs.version, "version", false, "print version")
	flag.BoolVar(&serverArgs.dev, "dev", true, "is dev")

	flag.Parse()
	if flag.NArg() > 0 {
		log.Fatalf("does not take non-flag arguments: %q.", flag.Args())
	}

	if serverArgs.version {
		fmt.Println(version.String())
		os.Exit(0)
	}

	// initialize wissy logger
	//
	err := wislog.InitWisLog(serverArgs.logLevel, serverArgs.logFile, serverArgs.dev)
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
		wislog.Logger.Error(err)
		log.Fatal(err)
	}

	ss := store.NewServerStore(sfs)
	err = ss.WritePrivateKey()
	if err != nil {
		wislog.Logger.Error(err)
		log.Fatal(err)
	}

	// new sever config
	//
	cfg := config.NewServerConfig(
		serverArgs.configpath, serverArgs.domain, serverArgs.certfile, serverArgs.certkey, serverArgs.turnSecret,
		serverArgs.stunURL, serverArgs.stunUserName, serverArgs.stunPassword,
		serverArgs.turnURL, serverArgs.turnUserName, serverArgs.turnPassword, serverArgs.ttl, serverArgs.credentials,
		serverArgs.signalURL,
	)

	// new cert store
	//
	certConfig := cert.NewCertConfig(
		serverArgs.letsencryptDir, cfg.TLSConfig.Domain, cfg.TURNConfig.Secret, cfg.TLSConfig.Certfile,
	)

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
	opts, err := server.NewGrpcServerOption(certConfig)
	if err != nil {
		wislog.Logger.Error(err)
		log.Fatal(err)
	}

	// grpc middleware config
	//
	middleware := server.NewMiddlware(auth0Client)

	opts = append(opts, grpc.UnaryInterceptor(grpc_auth.UnaryServerInterceptor(middleware.Authenticate)))
	grpcServer := grpc.NewServer(opts...)

	// initialize grpc server services
	//
	peer.RegisterPeerServiceServer(grpcServer, s.PeerServerService)
	user.RegisterUserServiceServer(grpcServer, s.UserServerService)
	session.RegisterSessionServiceServer(grpcServer, s.SessionServerService)
	admin_network.RegisterAdminNetworkServiceServer(grpcServer, s.AdminNetworkServerService)
	organization.RegisterOrganizationServiceServer(grpcServer, s.OrganizationServerService)

	wislog.Logger.Infof("starting server: localhost:%v", serverArgs.port)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", serverArgs.port))
	if err != nil {
		wislog.Logger.Errorf("failed to listen: %v", err)
		log.Fatalf("failed to listen: %v", err)
	}

	reflection.Register(grpcServer)
	go func() {
		if err = grpcServer.Serve(lis); err != nil {
			wislog.Logger.Errorf("failed to serve grpc server: %v.", err)
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
	wislog.Logger.Info("terminated server")
	grpcServer.Stop()
}
