package cmd

import (
	"context"
	"flag"
	"fmt"
	"log"

	wics "github.com/Notch-Technologies/wizy/cmd/wics/client"
	"github.com/Notch-Technologies/wizy/cmd/wissy/client"
	"github.com/Notch-Technologies/wizy/paths"
	"github.com/Notch-Technologies/wizy/store"
	"github.com/Notch-Technologies/wizy/types/flagtype"
	"github.com/Notch-Technologies/wizy/wislog"
	"github.com/peterbourgon/ff/ffcli"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

var loginArgs struct {
	clientpath string
	wicshost   string
	wicsport   int64
	setupkey   string
	logfile    string
	loglevel   string
	dev        bool
}

var loginCmd = &ffcli.Command{
	Name:      "login",
	Usage:     "login",
	ShortHelp: "login to wissy, start the management server and then run it",
	FlagSet: (func() *flag.FlagSet {
		fs := flag.NewFlagSet("login", flag.ExitOnError)
		fs.StringVar(&loginArgs.clientpath, "client", paths.DefaultClientConfigFile(), "client default config file")
		fs.StringVar(&loginArgs.wicshost, "host", "http://localhost", "wics server host url")
		fs.Int64Var(&loginArgs.wicsport, "port", flagtype.DefaultPort, "wics server host port")
		fs.StringVar(&loginArgs.setupkey, "key", "", "setup key issued by the wics server")
		fs.StringVar(&loginArgs.logfile, "logfile", paths.DefaultClientLogFile(), "set logfile path")
		fs.StringVar(&loginArgs.loglevel, "loglevel", wislog.DebugLevelStr, "set log level")
		fs.BoolVar(&loginArgs.dev, "dev", true, "is dev")
		return fs
	})(),
	Exec: execLogin,
}

func execLogin(args []string) error {
	err := wislog.InitWisLog(loginArgs.loglevel, loginArgs.logfile, loginArgs.dev)
	if err != nil {
		log.Fatalf("failed to initialize logger. %v", err)
	}
	_ = wislog.NewWisLog("login")

	// create client state key
	cfs, err := store.NewFileStore(paths.DefaultWicsClientStateFile())
	if err != nil {
		log.Fatalf("failed to create wics clietnt state. %v", err)
	}

	cs := store.NewClientStore(cfs)
	err = cs.WritePrivateKey()
	if err != nil {
		log.Fatalf("failed to write client state private key. %v", err)
	}

	// TOOD: (shintard) port to host to path ga kawatta baai ni taiou suru
	conf := client.GetClientConfig(loginArgs.clientpath, loginArgs.wicshost, int(loginArgs.wicsport))

	ctx := context.Background()

	privateKey, err := wgtypes.ParseKey(conf.WgPrivateKey)
	if err != nil {
		log.Fatalf("failed to parse wg private key. %v", err)
	}

	wicsClient, err := wics.NewWicsClient(ctx, conf.Host, int(loginArgs.wicsport), privateKey)
	if err != nil {
		log.Fatalf("failed to connect wics client. %v", err)
	}

	log.Printf("connected to wics server %s", conf.Host.String())

	serverPubKey, err := wicsClient.GetServerPublicKey()
	if err != nil {
		log.Fatalf("failed to get wics server public key. %v", err)
	}

	a, err := wicsClient.Login(loginArgs.setupkey, cs.GetPublicKey(), serverPubKey)
	if err != nil {
		log.Fatalf("failed to get wics server public key. %v", err)
	}
	fmt.Println(a)

	return nil
}
