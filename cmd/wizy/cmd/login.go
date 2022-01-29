package cmd

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strconv"

	"github.com/Notch-Technologies/wizy/cmd/wizy/tun"
	"github.com/Notch-Technologies/wizy/paths"
	"github.com/Notch-Technologies/wizy/types/flagtype"
	"github.com/Notch-Technologies/wizy/types/key"
	"github.com/Notch-Technologies/wizy/utils"
	"github.com/Notch-Technologies/wizy/wislog"
	"github.com/peterbourgon/ff/ffcli"
)

type Config struct {
	WgPrivateKey string
	Host *url.URL
	IgonoreTUNs []string
	TUNName string
	PreSharedKey string
}

func newClientConfig() *Config {
	if err := os.MkdirAll(filepath.Dir(loginArgs.clientpath), 0777); err != nil {
		log.Fatal(err)
	}

	privKey, err := key.NewGenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	scheme := loginArgs.wicshost + ":" + strconv.Itoa(int(loginArgs.wicsport))

	host, err := url.Parse(scheme)
	if err != nil {
		log.Fatal(err)
	}


	cfg := Config{
		WgPrivateKey: privKey,
		Host: host,
		TUNName: tun.TunName(),
		IgonoreTUNs: []string{},
	}

	b, err := json.MarshalIndent(cfg, "", "\t")
	if err != nil {
		log.Fatal(err)
	}

	if err = utils.AtomicWriteFile(loginArgs.clientpath, b, 0600); err != nil {
		log.Fatal(err)
	}

	return &cfg
}

func getClientConfig() *Config {
	b, err := ioutil.ReadFile(loginArgs.clientpath)
	switch {
	case errors.Is(err, os.ErrNotExist):
		return newClientConfig()
	case err != nil:
		log.Fatal(err)
		panic(err)
	default:
		var cfg Config
		if err := json.Unmarshal(b, &cfg); err != nil {
			log.Fatalf("can not read client config file. %v", err)
		}
		return &cfg
	}
}

var loginArgs struct {
	clientpath string
	wicshost string
	wicsport int64
	setupkey string
	logfile string
	loglevel string
	dev bool
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
	Exec:      execLogin,
}

func execLogin(args []string) error {
	err := wislog.InitWisLog(loginArgs.loglevel, loginArgs.logfile, loginArgs.dev)
	if err != nil {
		log.Fatalf("failed to initialize logger. %v", err)
	}

	l := wislog.NewWisLog("login")

	fmt.Println(l)

	return nil
}
