package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/Notch-Technologies/wizy/cmd/management/config"
	"github.com/Notch-Technologies/wizy/paths"
	//"github.com/Notch-Technologies/wizy/types/flagtype"
	"github.com/Notch-Technologies/wizy/version"
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
	// flag.Var(flagtype.PortValue(&args.port, flagtype.DefaultPort), "port", "specify the port of the management server")
	// flag.IntVar(&args.verbose, "verbose", 0, "0 is the default value, 1 is a redundant message")
	flag.StringVar(&args.storepath, "store", paths.DefaultStoreStateFile(), "path of management store state file")
	// flag.StringVar(&args.domain, "domain", "", "path of mangement file")
	// flag.StringVar(&args.certfile, "cert-file", "", "path of mangement file")
	// flag.StringVar(&args.certkey, "cert-key", "", "path of mangement file")
	// flag.BoolVar(&args.version, "version", false, "path of mangement file")

	flag.Parse()
	if flag.NArg() > 0 {
		log.Fatalf("does not take non-flag arguments: %q", flag.Args())
	}

	if args.version {
		fmt.Println(version.String())
		os.Exit(0)
	}

	loadConfig()
}

func loadConfig() config.Config {
	b, err := ioutil.ReadFile(args.configpath)
	switch {
 	case errors.Is(err, os.ErrNotExist):
 		return createNewConfig()
 	case err != nil:
 	    log.Fatal(err)
 	    panic("unreachable")
 	default:
		var cfg config.Config
 	    if err := json.Unmarshal(b, &cfg); err != nil {
 	        log.Fatalf("config: %v", err)
 	    }
 	    return cfg
 	}
}

func createNewConfig() config.Config {
	if err := os.MkdirAll(filepath.Dir(args.configpath), 0777); err != nil {
		log.Fatal(err)
	}

	cfg := config.Config{StorePath: args.storepath}

	b, err := json.MarshalIndent(cfg, "", "\t")
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(args.configpath, b, 0600)
	if err != nil {
		log.Fatal(err)
	}

	return cfg
}
