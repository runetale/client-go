package cmd

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/Notch-Technologies/wizy/cmd/server/client"
	signaling "github.com/Notch-Technologies/wizy/cmd/signaling/client"
	"github.com/Notch-Technologies/wizy/core"
	"github.com/Notch-Technologies/wizy/iface"
	"github.com/Notch-Technologies/wizy/paths"
	"github.com/Notch-Technologies/wizy/polymer/engine"
	"github.com/Notch-Technologies/wizy/store"
	"github.com/Notch-Technologies/wizy/types/flagtype"
	"github.com/Notch-Technologies/wizy/wislog"
	"github.com/peterbourgon/ff/v2/ffcli"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

var upArgs struct {
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

var upCmd = &ffcli.Command{
	Name:       "up",
	ShortUsage: "up",
	ShortHelp:  "launch a logged-in peer",
	FlagSet: (func() *flag.FlagSet {
		fs := flag.NewFlagSet("up", flag.ExitOnError)
		fs.StringVar(&upArgs.clientPath, "path", paths.DefaultClientConfigFile(), "client default config file")
		fs.StringVar(&upArgs.serverHost, "server-host", "http://172.16.165.129", "grpc server host url")
		fs.Int64Var(&upArgs.serverPort, "server-port", flagtype.DefaultGrpcServerPort, "grpc server host port")
		fs.StringVar(&upArgs.signalHost, "signal-host", "http://172.16.165.129", "signaling server host url")
		fs.Int64Var(&upArgs.signalPort, "signal-port", flagtype.DefaultSignalingServerPort, "signaling server host port")
		fs.StringVar(&upArgs.setupKey, "key", "", "setup key issued by the grpc server")
		fs.StringVar(&upArgs.logFile, "logfile", paths.DefaultClientLogFile(), "set logfile path")
		fs.StringVar(&upArgs.logLevel, "loglevel", wislog.DebugLevelStr, "set log level")
		fs.BoolVar(&upArgs.dev, "dev", true, "is dev")
		return fs
	})(),
	Exec: execUp,
}

func execUp(ctx context.Context, args []string) error {
	fmt.Println("exec up")
	// initialize wissy logger
	//
	err := wislog.InitWisLog(upArgs.logLevel, upArgs.logFile, upArgs.dev)
	if err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}

	wislog := wislog.NewWisLog("up")

	// initialize client Core
	//
	clientCore, err := core.NewClientCore(
		upArgs.clientPath, upArgs.serverHost, int(upArgs.serverPort),
		upArgs.signalHost, int(upArgs.signalPort), wislog)
	if err != nil {
		wislog.Logger.Fatalf("failed to initialize client core: %v", err)
	}
	clientCore = clientCore.GetClientCore()

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

	// parse wireguard privatekey
	//
	wgPrivateKey, err := wgtypes.ParseKey(clientCore.WgPrivateKey)
	if err != nil {
		wislog.Logger.Fatalf("failed to parse wg private key. %v", err)
	}

	// initialize grpc client
	//
	// ctx := context.Background()
	gClient, err := client.NewGrpcClient(ctx, clientCore.ServerHost, int(upArgs.serverPort), wgPrivateKey)
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

	err = iface.CreateIface(clientCore.TUNName, clientCore.WgPrivateKey, "10.0.0.2/24")
	if err != nil {
		wislog.Logger.Errorf("failed creating Wireguard interface [%s]: %s", clientCore.TUNName, err.Error())
		return err
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	engineConfig := engine.NewEngineConfig(wgPrivateKey, clientCore, "10.0.0.2/24")

	e := engine.NewEngine(wislog, gClient, sClient, cancel, ctx, engineConfig, cs.GetPublicKey(), wgPrivateKey)
	e.Start()

	select {
	case <-stopCh:
	case <-ctx.Done():
	}

	return nil
}
