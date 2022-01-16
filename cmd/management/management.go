package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/Notch-Technologies/wizy/cmd/management/config"
	"github.com/Notch-Technologies/wizy/paths"
	"github.com/Notch-Technologies/wizy/types/flagtype"
	"github.com/Notch-Technologies/wizy/version"
	"github.com/Notch-Technologies/wizy/store"
)

var args struct {
	configpath string
	port uint16
	verbose int
	storepath string
	domain string
	certfile string
	certkey string
	version bool
}

func main() {
	flag.StringVar(&args.configpath, "config", paths.DefaultManagementFile(), "path of mangement file")
	flag.Var(flagtype.PortValue(&args.port, flagtype.DefaultPort), "port", "specify the port of the management server")
	flag.IntVar(&args.verbose, "verbose", 0, "0 is the default value, 1 is a redundant message")
	flag.StringVar(&args.storepath, "store", paths.DefaultStoreStateFile(), "path of management store state file")
	flag.StringVar(&args.domain, "domain", "", "path of mangement file")
	flag.StringVar(&args.certfile, "cert-file", "", "path of mangement file")
	flag.StringVar(&args.certkey, "cert-key", "", "path of mangement file")
	flag.BoolVar(&args.version, "version", false, "path of mangement file")

	flag.Parse()
	if flag.NArg() > 0 {
		log.Fatalf("does not take non-flag arguments: %q", flag.Args())
	}

	if args.version {
		fmt.Println(version.String())
		os.Exit(0)
	}

	fs, err := store.LoadFileStore(args.storepath)
	if err != nil {
 	    log.Fatal(err)
	}

	fmt.Println(fs)

	cfg := config.LoadConfig(args.configpath, args.domain, args.certfile, args.certkey)
	fmt.Println(cfg)
}

