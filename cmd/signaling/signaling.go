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

	"github.com/Notch-Technologies/wizy/cmd/signaling/pb/negotiation"
	"github.com/Notch-Technologies/wizy/cmd/signaling/server"
	"github.com/Notch-Technologies/wizy/paths"
	"github.com/Notch-Technologies/wizy/types/flagtype"
	"github.com/Notch-Technologies/wizy/version"
	"github.com/Notch-Technologies/wizy/wislog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
)

var signalingArgs struct {
	port     uint16
	verbose  int
	domain   string
	certfile string
	certkey  string
	version  bool
	logFile  string
	logLevel string
	dev      bool
}

func main() {
	flag.Var(flagtype.PortValue(&signalingArgs.port, flagtype.DefaultSignalingServerPort), "port", "specify the port of the signaling server")
	flag.IntVar(&signalingArgs.verbose, "verbose", 0, "0 is the default value, 1 is a redundant message")
	flag.StringVar(&signalingArgs.domain, "domain", "", "your domain")
	flag.StringVar(&signalingArgs.certfile, "cert-file", "", "your cert")
	flag.StringVar(&signalingArgs.certkey, "cert-key", "", "your cert key")
	flag.BoolVar(&signalingArgs.version, "version", false, "print version")
	flag.StringVar(&signalingArgs.logFile, "logfile", paths.DefaultSignalingLogFile(), "set logfile path")
	flag.StringVar(&signalingArgs.logLevel, "loglevel", wislog.DebugLevelStr, "set log level")
	flag.BoolVar(&signalingArgs.dev, "dev", true, "is dev")

	flag.Parse()
	if flag.NArg() > 0 {
		log.Fatalf("does not take non-flag arguments: %q.", flag.Args())
	}

	if signalingArgs.version {
		fmt.Println(version.String())
		os.Exit(0)
	}

	// initialize wissy logger
	//
	err := wislog.InitWisLog(signalingArgs.logLevel, signalingArgs.logFile, signalingArgs.dev)
	if err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}

	wislog := wislog.NewWisLog("signaling")

	// initialize server
	//
	s := server.NewServer(wislog)

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

	opts = append(opts, grpc.KeepaliveEnforcementPolicy(kaep), grpc.KeepaliveParams(kasp))
	grpcServer := grpc.NewServer(opts...)

	negotiation.RegisterNegotiationServer(grpcServer, s.NegotiationServer)

	log.Printf("starting signaling server: localhost:%v", signalingArgs.port)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", signalingArgs.port))
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
