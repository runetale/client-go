package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/Notch-Technologies/wizy/cmd/wics/config"
	"github.com/Notch-Technologies/wizy/cmd/wics/proto"
	"github.com/Notch-Technologies/wizy/cmd/wics/server"
	"github.com/Notch-Technologies/wizy/paths"
	"github.com/Notch-Technologies/wizy/store"
	"github.com/Notch-Technologies/wizy/types/flagtype"
	"github.com/Notch-Technologies/wizy/version"
	"google.golang.org/grpc"
)

var args struct {
	configpath string
	port uint16
	verbose int
	storepath string
	domain string
	certfile string
	certkey string
	version bool
}

func main() {
	flag.StringVar(&args.configpath, "config", paths.DefaultWicsConfigFile(), "path of wics config file")
	flag.Var(flagtype.PortValue(&args.port, flagtype.DefaultPort), "port", "specify the port of the wics server")
	flag.IntVar(&args.verbose, "verbose", 0, "0 is the default value, 1 is a redundant message")
	flag.StringVar(&args.storepath, "store", paths.DefaultAccountStateFile(), "path of wics store state file")
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

	fs, err := store.NewFileStore(args.storepath)
	if err != nil {
 	    log.Fatal(err)
	}

	cfg := config.LoadConfig(args.configpath, args.domain, args.certfile, args.certkey)

	account := store.NewAccount(fs)

	grpcServer := grpc.NewServer()
	s, err :=  server.NewServer(cfg, account)
	if err != nil {
 	    log.Fatal(err)
	}

	proto.RegisterPeerServiceServer(grpcServer, s.PeerServiceServer)
	proto.RegisterUserServiceServer(grpcServer, s.UserServiceServer)
	log.Printf("started wics server: localhost:%v", args.port)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", args.port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

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
