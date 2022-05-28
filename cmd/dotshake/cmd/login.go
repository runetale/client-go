package cmd

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	grpc_client "github.com/Notch-Technologies/dotshake/client/grpc"
	"github.com/Notch-Technologies/dotshake/conf"
	"github.com/Notch-Technologies/dotshake/dotlog"
	"github.com/Notch-Technologies/dotshake/paths"
	"github.com/Notch-Technologies/dotshake/polymer"
	"github.com/Notch-Technologies/dotshake/polymer/conn"
	"github.com/Notch-Technologies/dotshake/store"
	"github.com/Notch-Technologies/dotshake/types/flagtype"
	"github.com/peterbourgon/ff/v2/ffcli"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
)

var loginArgs struct {
	clientPath string
	serverHost string
	serverPort int64
	signalHost string
	signalPort int64
	logFile    string
	logLevel   string
	dev        bool
}

var loginCmd = &ffcli.Command{
	Name:       "login",
	ShortUsage: "login [flags]",
	ShortHelp:  "login to dotshake, start the management server and then run it",
	FlagSet: (func() *flag.FlagSet {
		fs := flag.NewFlagSet("login", flag.ExitOnError)
		fs.StringVar(&loginArgs.clientPath, "path", paths.DefaultClientConfigFile(), "client default config file")
		fs.StringVar(&loginArgs.serverHost, "server-host", "http://127.0.0.1", "grpc server host url")
		fs.Int64Var(&loginArgs.serverPort, "server-port", flagtype.DefaultGrpcServerPort, "grpc server host port")
		fs.StringVar(&loginArgs.signalHost, "signal-host", "http://127.0.0.1", "signaling server host url")
		fs.Int64Var(&loginArgs.signalPort, "signal-port", flagtype.DefaultSignalingServerPort, "signaling server host port")
		fs.StringVar(&loginArgs.logFile, "logfile", paths.DefaultClientLogFile(), "set logfile path")
		fs.StringVar(&loginArgs.logLevel, "loglevel", dotlog.DebugLevelStr, "set log level")
		fs.BoolVar(&loginArgs.dev, "dev", true, "is dev")
		return fs
	})(),
	Exec: execLogin,
}

func execLogin(ctx context.Context, args []string) error {
	// initialize dotshake logger
	//
	err := dotlog.InitDotLog(loginArgs.logLevel, loginArgs.logFile, loginArgs.dev)
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

	fmt.Printf("client machine key: %s\n", cs.GetPublicKey())

	// initialize client conf
	//
	clientConf, err := conf.NewClientConf(
		loginArgs.clientPath,
		loginArgs.serverHost, int(loginArgs.serverPort),
		loginArgs.signalHost, int(loginArgs.signalPort),
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

	serverClient := grpc_client.NewServerClient(ctx, gconn)
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

	// ログインの処理はここまで
	// upに後で移行する

	// initialize polymer
	signalURL := res.SignalHost + ":" + strconv.Itoa(int(res.SignalPort))
	signalHostURL, err := url.Parse(signalURL)
	if err != nil {
		dotlog.Logger.Fatalf("failed to parsing signal host => [%s:%d]. because %v", res.SignalHost, res.SignalPort, err)
		return err
	}

	siconn, err := grpc.DialContext(
		clientCtx,
		signalHostURL.Host,
		option,
		grpc.WithBlock(),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:    10 * time.Second,
			Timeout: 10 * time.Second,
		}))
	if err != nil {
		dotlog.Logger.Fatalf("failed to connect server client. because %v", err)
	}

	connState := conn.NewConnectedState()
	signalClient := grpc_client.NewSignalClient(ctx, siconn, connState)

	pctx, cancel := context.WithCancel(ctx)

	polymer, err := polymer.NewPolymer(
		signalClient,
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

	// start polymer
	err = polymer.Start()
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
