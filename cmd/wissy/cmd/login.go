package cmd

import (
	"context"
	"flag"
	"fmt"
	"log"

	grpc_client "github.com/Notch-Technologies/wizy/cmd/server/grpc_client"
	"github.com/Notch-Technologies/wizy/cmd/wissy/client"
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
		fs.StringVar(&loginArgs.serverHost, "host", "http://172.16.165.129", "grpc server host url")
		fs.Int64Var(&loginArgs.serverPort, "port", flagtype.DefaultGrpcServerPort, "grpc server host port")
		fs.StringVar(&loginArgs.setupKey, "key", "", "setup key issued by the grpc server")
		fs.StringVar(&loginArgs.logFile, "logfile", paths.DefaultClientLogFile(), "set logfile path")
		fs.StringVar(&loginArgs.logLevel, "loglevel", wislog.DebugLevelStr, "set log level")
		fs.BoolVar(&loginArgs.dev, "dev", true, "is dev")
		return fs
	})(),
	Exec: execLogin,
}

func execLogin(args []string) error {
	err := wislog.InitWisLog(loginArgs.logLevel, loginArgs.logFile, loginArgs.dev)
	if err != nil {
		log.Fatalf("failed to initialize logger. %v", err)
	}
	wisLog := wislog.NewWisLog("login")

	// create client state key
	cfs, err := store.NewFileStore(paths.DefaultWicsClientStateFile())
	if err != nil {
		log.Fatalf("failed to create clietnt state. %v", err)
	}

	cs := store.NewClientStore(cfs)
	err = cs.WritePrivateKey()
	if err != nil {
		log.Fatalf("failed to write client state private key. %v", err)
	}

	conf := client.GetClientConfig(loginArgs.clientPath, loginArgs.serverHost, int(loginArgs.serverPort))

	ctx := context.Background()

	wgPrivateKey, err := wgtypes.ParseKey(conf.WgPrivateKey)
	if err != nil {
		log.Fatalf("failed to parse wg private key. %v", err)
	}

	client, err := grpc_client.NewGrpcClient(ctx, conf.ServerHost, int(loginArgs.serverPort), wgPrivateKey)
	if err != nil {
		log.Fatalf("failed to connect client. %v", err)
	}

	log.Printf("connected to server %s", conf.ServerHost.String())

	serverPubKey, err := client.GetServerPublicKey()
	if err != nil {
		log.Fatalf("failed to get server public key. %v", err)
	}

	_, err = client.Login(loginArgs.setupKey, cs.GetPublicKey(), serverPubKey, "10.0.0.1", wgPrivateKey)
	if err != nil {
		log.Fatalf("failed to login. %v", err)
	}

	// TOOD: upCmd Exec

	err = iface.CreateIface(conf.TUNName, conf.WgPrivateKey, "10.0.0.1/24")
	if err != nil {
		fmt.Printf("failed creating Wireguard interface [%s]: %s", conf.TUNName, err.Error())
		return err
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	engineConfig := polymer.NewEngineConfig(wgPrivateKey, conf, "10.0.0.1/24")

	e := polymer.NewEngine(wisLog, client, cancel, ctx, engineConfig, cs.GetPublicKey(), wgPrivateKey)
	e.Start()

	select {
	case <-stopCh:
	case <-ctx.Done():
	}

	return nil
}
