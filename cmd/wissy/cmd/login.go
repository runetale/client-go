package cmd

import (
	"flag"
	"fmt"
	"log"

	"github.com/Notch-Technologies/wizy/cmd/wissy/client"
	"github.com/Notch-Technologies/wizy/paths"
	"github.com/Notch-Technologies/wizy/types/flagtype"
	"github.com/Notch-Technologies/wizy/wislog"
	"github.com/peterbourgon/ff/ffcli"
)

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


	conf := client.GetClientConfig(loginArgs.clientpath, loginArgs.wicshost, int(loginArgs.wicsport))
	fmt.Println(conf)
	fmt.Println(loginArgs.clientpath)

	return nil
}
