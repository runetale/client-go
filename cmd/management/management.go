package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/Notch-Technologies/wizy/paths"
	"github.com/Notch-Technologies/wizy/version"
)

// TODO: (shintard) move to other directory
type portValue struct{ n *uint16 }

const DefaultPort = 443

func PortValue(dst *uint16, defaultPort uint16) flag.Value {
	*dst = defaultPort
	return portValue{dst}
}

func (p portValue) String() string {
	if p.n == nil {
		return ""
	}
	return fmt.Sprint(*p.n)
}

func (p portValue) Set(v string) error {
	if v == "" {
		return errors.New("can't be the empty string")
	}
	if strings.Contains(v, ":") {
		return errors.New("expecting just a port number, without a colon")
	}
	n, err := strconv.ParseUint(v, 10, 64) // use 64 instead of 16 to return nicer error message
	if err != nil {
		return fmt.Errorf("not a valid number")
	}
	if n > math.MaxUint16 {
		return errors.New("out of range for port number")
	}
	*p.n = uint16(n)
	return nil
}

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
	// flag.Var(PortValue(&args.port, DefaultPort), "port", "specify the port of the management server")
	// flag.IntVar(&args.verbose, "verbose", 0, "0 is the default value, 1 is a redundant message")
	// flag.StringVar(&args.storepath, "store", paths.DefaultStoreStateFile(), "path of management store state file")
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

func loadConfig() {
	b, err := ioutil.ReadFile(args.configpath)
	switch {
 	case err != nil:
 	    log.Fatal(err)
 	    panic("unreachable")
 	case errors.Is(err, os.ErrNotExist):
 	//    return writeNewConfig()
		fmt.Println("ErrNotExist")
		fmt.Printf("check management config file: %s\n", paths.DefaultManagementFile())
		return
 	default:
		// var cfg config
 	    // if err := json.Unmarshal(b, &cfg); err != nil {
 	    //     log.Fatalf("config: %v", err)
 	    // }
 	    // return cfg
		fmt.Println("default")
		fmt.Println(b)
		return
 	}
}
