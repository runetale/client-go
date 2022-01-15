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
	"github.com/Notch-Technologies/wizy/types/flagtype"
	"github.com/Notch-Technologies/wizy/version"
	"github.com/Notch-Technologies/wizy/utils"
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

	fs, err := loadFileStore()
	if err != nil {
 	    log.Fatal(err)
	}

	fmt.Println(fs)

	cfg := loadConfig()
	fmt.Println(cfg)
}

func loadConfig() config.Config {
	b, err := ioutil.ReadFile(args.configpath)
	switch {
 	case errors.Is(err, os.ErrNotExist):
 		return createNewConfig()
 	case err != nil:
 	    log.Fatal(err)
 	    panic("failed to load cofig")
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

	cfg := config.Config{
		StorePath: args.storepath,
		TLSConfig: config.TLSConfig{
			Domain: args.domain,
			Certfile: args.certfile,
			CertKey: args.certkey,
		},
	}

	b, err := json.MarshalIndent(cfg, "", "\t")
	if err != nil {
		log.Fatal(err)
	}

	if err = utils.AtomicWriteFile(args.configpath, b, 0600); err != nil {
		log.Fatal(err)
	}

	return cfg
}

// TODO: (shintard) consider whether to manage users in FileStore or in a different structure (json).
type FileStore struct {
	path string
}

func loadFileStore() (*FileStore, error) {
	b, err := ioutil.ReadFile(args.storepath)
	if err != nil {
		if os.IsNotExist(err) {
			os.MkdirAll(filepath.Dir(args.storepath), 0755)
			store := FileStore{args.storepath}

			b, err := json.MarshalIndent(store, "", "\t")
			if err != nil {
				log.Fatal(err)
			}

			if err = utils.AtomicWriteFile(args.storepath, b, 0600); err != nil {
				return nil, err
			}
			return &store, nil
		}
		return nil, err
	}

	var fs FileStore

	if err := json.Unmarshal(b, &fs); err != nil {
		return nil, err
	}

	return &fs, nil
}
