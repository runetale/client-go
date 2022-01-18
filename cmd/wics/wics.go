package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

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
	flag.StringVar(&args.configpath, "config", paths.DefaultWicsFile(), "path of wics file")
	flag.Var(flagtype.PortValue(&args.port, flagtype.DefaultPort), "port", "specify the port of the wics server")
	flag.IntVar(&args.verbose, "verbose", 0, "0 is the default value, 1 is a redundant message")
	flag.StringVar(&args.storepath, "store", paths.DefaultWicsStateFile(), "path of wics store state file")
	flag.StringVar(&args.domain, "domain", "", "your domain")
	flag.StringVar(&args.certfile, "cert-file", "", "your cert")
	flag.StringVar(&args.certkey, "cert-key", "", "your cert key")
	flag.BoolVar(&args.version, "version", false, "print version")

	flag.Parse()
	if flag.NArg() > 0 {
		log.Fatalf("does not take non-flag arguments: %q", flag.Args())
	}

	if args.version {
		fmt.Println(version.String())
		os.Exit(0)
	}

	fs, err := store.NewFileStore(args.storepath)
	if err != nil {
 	    log.Fatal(err)
	}

	fmt.Println(fs)

	cfg := config.LoadConfig(args.configpath, args.domain, args.certfile, args.certkey)
	fmt.Println(cfg)

	account := store.NewAccount(fs)
	fmt.Println(account)

	grpcServer := grpc.NewServer()
	s, err :=  server.NewServer(cfg, account)
	if err != nil {
		fmt.Println("aaa")
 	    log.Fatal(err)
	}

	proto.RegisterPeerServiceServer(grpcServer, s)
	proto.RegisterUserServiceServer(grpcServer, s)
	log.Printf("started wics server: :%v", args.port)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", args.port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	go func() {
		if err = grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve grpc server: %v", err)
		}
	}()

	stopCh := make(chan struct{})
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			stopCh <- struct{}{}
		}
	}()
	<-stopCh
	log.Println("terminated wics server")

	grpcServer.Stop()
}
