package cmd

import (
	"context"
	"flag"
	"log"
	"time"

	grpc_client "github.com/Notch-Technologies/dotshake/client/grpc"
	"github.com/Notch-Technologies/dotshake/core"
	"github.com/Notch-Technologies/dotshake/daemon"
	"github.com/Notch-Technologies/dotshake/dotlog"
	"github.com/Notch-Technologies/dotshake/paths"
	"github.com/Notch-Technologies/dotshake/polymer/conn"
	"github.com/Notch-Technologies/dotshake/store"
	"github.com/Notch-Technologies/dotshake/types/flagtype"
	"github.com/peterbourgon/ff/v2/ffcli"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
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
		fs.StringVar(&upArgs.logLevel, "loglevel", dotlog.DebugLevelStr, "set log level")
		fs.BoolVar(&upArgs.dev, "dev", true, "is dev")
		fs.BoolVar(&upArgs.daemon, "daemon", false, "whether to run the daemon process")
		return fs
	})(),
	Exec: execUp,
}

func execUp(ctx context.Context, args []string) error {
	// initialize dotshake logger
	//
	err := dotlog.InitDotLog(upArgs.logLevel, upArgs.logFile, upArgs.dev)
	if err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}
	dotlog := dotlog.NewDotLog("up")

	// initialize file store
	//
	cfs, err := store.NewFileStore(paths.DefaultWicsClientStateFile(), dotlog)
	if err != nil {
		dotlog.Logger.Fatalf("failed to create clietnt state. because %v", err)
	}
	cs := store.NewClientStore(cfs, dotlog)
	err = cs.WritePrivateKey()
	if err != nil {
		dotlog.Logger.Fatalf("failed to write client state private key. because %v", err)
	}

	// initialize client Core
	//
	clientCore, err := core.NewClientCore(
		upArgs.clientPath,
		upArgs.serverHost, int(upArgs.serverPort),
		upArgs.signalHost, int(upArgs.signalPort),
		dotlog)
	if err != nil {
		dotlog.Logger.Fatalf("failed to initialize client core. because %v", err)
	}
	clientCore = clientCore.GetClientCore()

	// parse wireguard private key
	wgPrivateKey, err := wgtypes.ParseKey(clientCore.WgPrivateKey)
	if err != nil {
		dotlog.Logger.Fatalf("failed to parse wg private key. because %v", err)
	}

	// TODO:(shintard) impl grpc server client
	// initialize server client
	//

	// gClient, err := client.NewGrpcClient(ctx, clientCore.ServerHost, wgPrivateKey, dotlog)
	// if err != nil {
	// 	dotlog.Logger.Fatalf("failed to connect server client. because %v", err)
	// }

	// dotlog.Logger.Infof("connected to server %s", clientCore.ServerHost.String())

	// // get server public key
	// //
	// serverPubKey, err := gClient.GetServerPublicKey()
	// if err != nil {
	// 	dotlog.Logger.Fatalf("failed to get server public key. %v", err)
	// }

	// // start logging in
	// //
	// // TODO: (shintard) if setupkey is nil, use public key and wireguard public key of client machine key to obtain login information and start Engine.
	// loginRes, err := gClient.Login(upArgs.setupKey, cs.GetPublicKey(), serverPubKey, wgPrivateKey)
	// if err != nil {
	// 	dotlog.Logger.Fatalf("failed to login. because %v", err)
	// }

	// dotlog.Logger.Infof("setup_key: [%s] was generated from [%s]", loginRes.SetupKey, loginRes.ClientPublicKey)

	clientCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// initialize signaling client
	//
	option := grpc.WithTransportCredentials(insecure.NewCredentials())
	gconn, err := grpc.DialContext(
		clientCtx,
		clientCore.SignalHost.Host,
		option,
		grpc.WithBlock(),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:    10 * time.Second,
			Timeout: 10 * time.Second,
		}))

	// TODO: (shintard) Engineに渡してあげる
	_ = grpc_client.NewSignalClient(ctx, gconn, conn.NewConnectedState())
	if err != nil {
		dotlog.Logger.Fatalf("failed to connect server client. because %v", err)
	}

	// create wireguard interface
	//
	// i := iface.NewIface(clientCore.TunName, clientCore.WgPrivateKey, loginRes.Ip, loginRes.Cidr, dotlog)
	// err = iface.CreateIface(i, loginRes.Ip, strconv.Itoa(int(loginRes.Cidr)), dotlog)

	// dotlog.Logger.Infof("ip: %s. cidr: %d", loginRes.Ip, loginRes.Cidr)

	if err != nil {
		dotlog.Logger.Errorf("failed creating Wireguard interface [%s]: %s", clientCore.TunName, err.Error())
		return err
	}

	// initialize engine
	//

	// engineConfig := engine.NewEngineConfig(wgPrivateKey, clientCore, loginRes.Ip, strconv.Itoa(int(loginRes.Cidr)))
	// e := engine.NewEngine(dotlog, i, gClient, sClient, cancel, ctx, engineConfig, cs.GetPublicKey(), wgPrivateKey)

	// setup daemon
	//
	if upArgs.daemon {
		d := daemon.NewDaemon(daemon.BinPath, daemon.ServiceName, daemon.DaemonFilePath, daemon.SystemConfig, dotlog)
		err = d.Install()
		if err != nil {
			dotlog.Logger.Errorf("failed to install daemon: %s", err.Error())
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
