package cmd

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	grpc_client "github.com/Notch-Technologies/dotshake/client/grpc"
	"github.com/Notch-Technologies/dotshake/conf"
	"github.com/Notch-Technologies/dotshake/dotengine"
	"github.com/Notch-Technologies/dotshake/dotlog"
	"github.com/Notch-Technologies/dotshake/paths"
	"github.com/Notch-Technologies/dotshake/store"
	"github.com/Notch-Technologies/dotshake/types/flagtype"
	"github.com/peterbourgon/ff/v2/ffcli"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
)

var upArgs struct {
	clientPath string
	serverHost string
	serverPort int64
	logFile    string
	logLevel   string
	dev        bool
}

var upCmd = &ffcli.Command{
	Name:       "up",
	ShortUsage: "up [flags]",
	ShortHelp:  "up to dotshake, communication client of dotshake",
	FlagSet: (func() *flag.FlagSet {
		fs := flag.NewFlagSet("up", flag.ExitOnError)
		fs.StringVar(&upArgs.clientPath, "path", paths.DefaultClientConfigFile(), "client default config file")
		fs.StringVar(&upArgs.serverHost, "server-host", "http://127.0.0.1", "grpc server host url")
		fs.Int64Var(&upArgs.serverPort, "server-port", flagtype.DefaultGrpcServerPort, "grpc server host port")
		fs.StringVar(&upArgs.logFile, "logfile", paths.DefaultClientLogFile(), "set logfile path")
		fs.StringVar(&upArgs.logLevel, "loglevel", dotlog.DebugLevelStr, "set log level")
		fs.BoolVar(&upArgs.dev, "dev", true, "is dev")
		return fs
	})(),
	Exec: execUp,
}

func execUp(ctx context.Context, args []string) error {
	// initialize dotshake logger
	//
	err := dotlog.InitDotLog(upArgs.logLevel, upArgs.logFile, upArgs.dev)
	if err != nil {
		log.Fatalf("failed to initialize logger. because %v", err)
	}

	dotlog := dotlog.NewDotLog("login")

	// initialize file store
	//
	cfs, err := store.NewFileStore(paths.DefaultDotshakeClientStateFile(), dotlog)
	if err != nil {
		dotlog.Logger.Fatalf("failed to create clietnt state. because %v", err)
	}

	// initialize client store
	//
	cs := store.NewClientStore(cfs, dotlog)
	err = cs.WritePrivateKey()
	if err != nil {
		dotlog.Logger.Fatalf("failed to write client state private key. because %v", err)
	}

	dotlog.Logger.Debugf("client machine key: %s", cs.GetPublicKey())

	// initialize client conf
	//
	clientConf, err := conf.NewClientConf(
		upArgs.clientPath,
		upArgs.serverHost, int(upArgs.serverPort),
		// upArgs.signalHost, int(upArgs.signalPort),
		dotlog,
	)
	if err != nil {
		fmt.Println(err)
		dotlog.Logger.Fatalf("failed to initialize client core. because %v", err)
	}

	clientConf = clientConf.GetClientConf()

	// initialize grpc client
	//
	clientCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	option := grpc.WithTransportCredentials(insecure.NewCredentials())
	gconn, err := grpc.DialContext(
		clientCtx,
		clientConf.ServerHost.Host,
		option,
		grpc.WithBlock(),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:    10 * time.Second,
			Timeout: 10 * time.Second,
		}))

	serverClient := grpc_client.NewServerClient(ctx, gconn, dotlog)
	if err != nil {
		dotlog.Logger.Fatalf("failed to connect server client. because %v", err)
	}

	res, err := serverClient.GetMachine(cs.GetPublicKey())
	if err != nil {
		return err
	}

	if !res.IsRegistered {
		fmt.Printf("please log in via this link => %s\n", res.LoginUrl)
		msg, err := serverClient.ConnectStreamPeerLoginSession(cs.GetPublicKey())
		if err != nil {
			return err
		}
		res.Ip = msg.Ip
		res.Cidr = msg.Cidr
	}

	pctx, cancel := context.WithCancel(ctx)

	engine, err := dotengine.NewDotEngine(
		serverClient,
		dotlog,
		clientConf.TunName,
		cs.GetPublicKey(),
		res.Ip,
		res.Cidr,
		clientConf.WgPrivateKey,
		clientConf.BlackList,
		pctx,
		cancel,
	)
	if err != nil {
		dotlog.Logger.Fatalf("failed to connect signal client. because %v", err)
		return err
	}

	// start engine
	err = engine.Start()
	if err != nil {
		dotlog.Logger.Fatalf("failed to start polymer. because %v", err)
		return err
	}

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

	return nil
}
