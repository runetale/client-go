package cmd

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/Notch-Technologies/wizy/cmd/server/client"
	"github.com/Notch-Technologies/wizy/core"
	"github.com/Notch-Technologies/wizy/paths"
	"github.com/Notch-Technologies/wizy/store"
	"github.com/Notch-Technologies/wizy/types/flagtype"
	"github.com/Notch-Technologies/wizy/wislog"
	"github.com/peterbourgon/ff/v2/ffcli"
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
	Name:       "login",
	ShortUsage: "login [flags]",
	ShortHelp:  "login to wissy, start the management server and then run it",
	FlagSet: (func() *flag.FlagSet {
		fs := flag.NewFlagSet("login", flag.ExitOnError)
		fs.StringVar(&loginArgs.clientPath, "path", paths.DefaultClientConfigFile(), "client default config file")
		fs.StringVar(&loginArgs.serverHost, "server-host", "", "grpc server host url")
		fs.Int64Var(&loginArgs.serverPort, "server-port", flagtype.DefaultGrpcServerPort, "grpc server host port")
		fs.StringVar(&loginArgs.signalHost, "signal-host", "", "signaling server host url")
		fs.Int64Var(&loginArgs.signalPort, "signal-port", flagtype.DefaultSignalingServerPort, "signaling server host port")
		fs.StringVar(&loginArgs.setupKey, "key", "", "setup key issued by the grpc server")
		fs.StringVar(&loginArgs.logFile, "logfile", paths.DefaultClientLogFile(), "set logfile path")
		fs.StringVar(&loginArgs.logLevel, "loglevel", wislog.DebugLevelStr, "set log level")
		fs.BoolVar(&loginArgs.dev, "dev", true, "is dev")
		return fs
	})(),
	Exec: execLogin,
}

func execLogin(ctx context.Context, args []string) error {
	// initialize wissy logger
	//
	err := wislog.InitWisLog(loginArgs.logLevel, loginArgs.logFile, loginArgs.dev)
	if err != nil {
		log.Fatalf("failed to initialize logger. because %v", err)
	}

	wislog := wislog.NewWisLog("login")

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
		loginArgs.clientPath,
		loginArgs.serverHost, int(loginArgs.serverPort),
		loginArgs.signalHost, int(loginArgs.signalPort),
		wislog,
	)
	if err != nil {
		wislog.Logger.Fatalf("failed to initialize client core. because %v", err)
	}
	clientCore = clientCore.GetClientCore()

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
	login, err := gClient.Login(loginArgs.setupKey, cs.GetPublicKey(), serverPubKey, wgPrivateKey)
	if err != nil {
		wislog.Logger.Fatalf("failed to login. because %v", err)
	}

	wislog.Logger.Infof("setup_key: [%s] was generated from [%s]", login.SetupKey, login.ClientPublicKey)

	fmt.Println("login succeded.")
	fmt.Printf("your ip [%s/%d]", login.Ip, login.Cidr)
	fmt.Println("type the following command to activate.")
	fmt.Printf("$ sudo wissy up -key %s", login.SetupKey)

	return nil
}
