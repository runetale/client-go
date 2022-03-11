package cmd

import (
	"context"
	"flag"
	"log"
	"strconv"

	"github.com/Notch-Technologies/wizy/cmd/server/client"
	signaling "github.com/Notch-Technologies/wizy/cmd/signaling/client"
	"github.com/Notch-Technologies/wizy/core"
	"github.com/Notch-Technologies/wizy/daemon"
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
	daemon     bool
}

var upCmd = &ffcli.Command{
	Name:       "up",
	ShortUsage: "up [flags]",
	ShortHelp:  "startup a peer. you can use it after logging in",
	FlagSet: (func() *flag.FlagSet {
		fs := flag.NewFlagSet("up", flag.ExitOnError)
		fs.StringVar(&upArgs.clientPath, "path", paths.DefaultClientConfigFile(), "client default config file")
		fs.StringVar(&upArgs.serverHost, "server-host", "", "grpc server host url")
		fs.Int64Var(&upArgs.serverPort, "server-port", flagtype.DefaultGrpcServerPort, "grpc server port")
		fs.StringVar(&upArgs.signalHost, "signal-host", "", "signaling server host url")
		fs.Int64Var(&upArgs.signalPort, "signal-port", flagtype.DefaultSignalingServerPort, "signaling server port")
		fs.StringVar(&upArgs.setupKey, "key", "", "setup key issued by the grpc server")
		fs.StringVar(&upArgs.logFile, "logfile", paths.DefaultClientLogFile(), "set logfile path")
		fs.StringVar(&upArgs.logLevel, "loglevel", wislog.DebugLevelStr, "set log level")
		fs.BoolVar(&upArgs.dev, "dev", true, "is dev")
		fs.BoolVar(&upArgs.daemon, "daemon", false, "whether to run the daemon process")
		return fs
	})(),
	Exec: execUp,
}

func execUp(ctx context.Context, args []string) error {
	// initialize wissy logger
	//
	err := wislog.InitWisLog(upArgs.logLevel, upArgs.logFile, upArgs.dev)
	if err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}
	wislog := wislog.NewWisLog("up")

	// initialize file store
	//
	cfs, err := store.NewFileStore(paths.DefaultWicsClientStateFile(), wislog)
	if err != nil {
		wislog.Logger.Fatalf("failed to create clietnt state. because %v", err)
	}
	cs := store.NewClientStore(cfs, wislog)
	err = cs.WritePrivateKey()
	if err != nil {
		wislog.Logger.Fatalf("failed to write client state private key. because %v", err)
	}

	// initialize client Core
	//
	clientCore, err := core.NewClientCore(
		upArgs.clientPath,
		upArgs.serverHost, int(upArgs.serverPort),
		upArgs.signalHost, int(upArgs.signalPort),
		wislog)
	if err != nil {
		wislog.Logger.Fatalf("failed to initialize client core. because %v", err)
	}
	clientCore = clientCore.GetClientCore()

	// parse wireguard private key
	wgPrivateKey, err := wgtypes.ParseKey(clientCore.WgPrivateKey)
	if err != nil {
		wislog.Logger.Fatalf("failed to parse wg private key. because %v", err)
	}

	// initialize grpc client
	//
	gClient, err := client.NewGrpcClient(ctx, clientCore.ServerHost, wgPrivateKey, wislog)
	if err != nil {
		wislog.Logger.Fatalf("failed to connect server client. because %v", err)
	}

	wislog.Logger.Infof("connected to server %s", clientCore.ServerHost.String())

	// get server public key
	//
	serverPubKey, err := gClient.GetServerPublicKey()
	if err != nil {
		wislog.Logger.Fatalf("failed to get server public key. %v", err)
	}

	// start logging in
	//
	// TODO: (shintard) if setupkey is nil, use public key and wireguard public key of client machine key to obtain login information and start Engine.
	loginRes, err := gClient.Login(upArgs.setupKey, cs.GetPublicKey(), serverPubKey, wgPrivateKey)
	if err != nil {
		wislog.Logger.Fatalf("failed to login. because %v", err)
	}

	wislog.Logger.Infof("setup_key: [%s] was generated from [%s]", loginRes.SetupKey, loginRes.ClientPublicKey)

	// initialize signaling client
	//
	sClient, err := signaling.NewSignalingClient(ctx, clientCore.SignalHost, wgPrivateKey, wislog)
	// sClient, err := signaling.NewSignalingClient(ctx, login.SignalingHost, wgPrivateKey, wislog)
	if err != nil {
		wislog.Logger.Fatalf("failed to connect signaling client. %v", err)
		return err
	}

	wislog.Logger.Infof("connected to signaling server %s", loginRes.SignalingHost)

	// create wireguard interface
	//
	i := iface.NewIface(clientCore.TunName, clientCore.WgPrivateKey, loginRes.Ip, loginRes.Cidr, wislog)
	err = iface.CreateIface(i, loginRes.Ip, strconv.Itoa(int(loginRes.Cidr)))
	if err != nil {
		wislog.Logger.Errorf("failed creating Wireguard interface [%s]: %s", clientCore.TunName, err.Error())
		return err
	}

	// initialize engine
	//
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	engineConfig := engine.NewEngineConfig(wgPrivateKey, clientCore, loginRes.Ip, strconv.Itoa(int(loginRes.Cidr)))
	e := engine.NewEngine(wislog, i, gClient, sClient, cancel, ctx, engineConfig, cs.GetPublicKey(), wgPrivateKey)

	// setup daemon
	//
	if upArgs.daemon {
		d := daemon.NewDaemon(daemon.BinPath, daemon.ServiceName, daemon.DameonFilePath, daemon.SystemConfig, wislog)
		err = d.Install()
		if err != nil {
			wislog.Logger.Errorf("failed to install daemon: %s", err.Error())
			return err
		}
		return nil
	}

	// start engine
	//
	e.Start()

	select {
	case <-stopCh:
	case <-ctx.Done():
	}

	return nil
}
