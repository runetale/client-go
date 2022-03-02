package cmd

import (
	"context"
	"flag"
	"log"

	"github.com/Notch-Technologies/wizy/cmd/server/client"
	signaling "github.com/Notch-Technologies/wizy/cmd/signaling/client"
	"github.com/Notch-Technologies/wizy/core"
	"github.com/Notch-Technologies/wizy/iface"
	"github.com/Notch-Technologies/wizy/paths"
	"github.com/Notch-Technologies/wizy/polymer"
	"github.com/Notch-Technologies/wizy/store"
	"github.com/Notch-Technologies/wizy/types/flagtype"
	"github.com/Notch-Technologies/wizy/wislog"
	"github.com/peterbourgon/ff/ffcli"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

var (
	stopCh chan int
)

func init() {
	stopCh = make(chan int)
}

var loginArgs struct {
	clientPath string
	serverHost string
	serverPort int64
	signalHost string
	signalPort int64
	setupKey   string
	logFile    string
	logLevel   string
	dev        bool
}

var loginCmd = &ffcli.Command{
	Name:      "login",
	Usage:     "login",
	ShortHelp: "login to wissy, start the management server and then run it",
	FlagSet: (func() *flag.FlagSet {
		fs := flag.NewFlagSet("login", flag.ExitOnError)
		fs.StringVar(&loginArgs.clientPath, "path", paths.DefaultClientConfigFile(), "client default config file")
		fs.StringVar(&loginArgs.serverHost, "server-host", "http://172.16.165.129", "grpc server host url")
		fs.Int64Var(&loginArgs.serverPort, "server-port", flagtype.DefaultGrpcServerPort, "grpc server host port")
		fs.StringVar(&loginArgs.signalHost, "signal-host", "http://172.16.165.129", "signaling server host url")
		fs.Int64Var(&loginArgs.signalPort, "signal-port", flagtype.DefaultSignalingServerPort, "signaling server host port")
		fs.StringVar(&loginArgs.setupKey, "key", "", "setup key issued by the grpc server")
		fs.StringVar(&loginArgs.logFile, "logfile", paths.DefaultClientLogFile(), "set logfile path")
		fs.StringVar(&loginArgs.logLevel, "loglevel", wislog.DebugLevelStr, "set log level")
		fs.BoolVar(&loginArgs.dev, "dev", true, "is dev")
		return fs
	})(),
	Exec: execLogin,
}

func execLogin(args []string) error {
	// initialize wissy logger
	//
	err := wislog.InitWisLog(loginArgs.logLevel, loginArgs.logFile, loginArgs.dev)
	if err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}

	wislog := wislog.NewWisLog("login")

	// initialize file store
	//
	cfs, err := store.NewFileStore(paths.DefaultWicsClientStateFile(), wislog)
	if err != nil {
		wislog.Logger.Fatalf("failed to create clietnt state: %v", err)
	}

	cs := store.NewClientStore(cfs, wislog)
	err = cs.WritePrivateKey()
	if err != nil {
		wislog.Logger.Fatalf("failed to write client state private key: %v", err)
	}

	// initialize client Core
	//
	clientCore, err := core.NewClientCore(
		loginArgs.clientPath, loginArgs.serverHost, int(loginArgs.serverPort),
		loginArgs.signalHost, int(loginArgs.signalPort), wislog)
	if err != nil {
		wislog.Logger.Fatalf("failed to initialize client core: %v", err)
	}
	clientCore = clientCore.GetClientCore()

	ctx := context.Background()

	wgPrivateKey, err := wgtypes.ParseKey(clientCore.WgPrivateKey)
	if err != nil {
		wislog.Logger.Fatalf("failed to parse wg private key. %v", err)
	}

	// initialize grpc client
	//
	gClient, err := client.NewGrpcClient(ctx, clientCore.ServerHost, int(loginArgs.serverPort), wgPrivateKey)
	if err != nil {
		wislog.Logger.Fatalf("failed to connect server client. %v", err)
	}

	wislog.Logger.Infof("connected to server %s", clientCore.ServerHost.String())

	// initialize signaling client
	//
	sClient, err := signaling.NewSignalingClient(ctx, clientCore.SignalHost, wgPrivateKey)
	if err != nil {
		wislog.Logger.Fatalf("failed to connect signaling client. %v", err)
	}

	wislog.Logger.Infof("connected to signaling server %s", clientCore.SignalHost.String())

	serverPubKey, err := gClient.GetServerPublicKey()
	if err != nil {
		wislog.Logger.Fatalf("failed to get server public key. %v", err)
	}

	_, err = gClient.Login(loginArgs.setupKey, cs.GetPublicKey(), serverPubKey, "10.0.0.2", wgPrivateKey)
	if err != nil {
		wislog.Logger.Fatalf("failed to login. %v", err)
	}

	// TODO: (shintard) ここからはupCmd
	err = iface.CreateIface(clientCore.TUNName, clientCore.WgPrivateKey, "10.0.0.2/24")
	if err != nil {
		wislog.Logger.Errorf("failed creating Wireguard interface [%s]: %s", clientCore.TUNName, err.Error())
		return err
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	engineConfig := polymer.NewEngineConfig(wgPrivateKey, clientCore, "10.0.0.2/24")

	e := polymer.NewEngine(wislog, gClient, sClient, cancel, ctx, engineConfig, cs.GetPublicKey(), wgPrivateKey)
	e.Start()

	select {
	case <-stopCh:
	case <-ctx.Done():
	}

	return nil
}
