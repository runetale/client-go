package client

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strconv"

	"github.com/Notch-Technologies/wizy/cmd/wissy/tun"
	"github.com/Notch-Technologies/wizy/types/key"
	"github.com/Notch-Technologies/wizy/utils"
)

type Config struct {
	WgPrivateKey string
	Host *url.URL
	IgonoreTUNs []string
	TUNName string
	PreSharedKey string
}

func newClientConfig(path string, host string, port int) *Config {
	if err := os.MkdirAll(filepath.Dir(path), 0777); err != nil {
		log.Fatal(err)
	}

	privKey, err := key.NewGenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	scheme := host + ":" + strconv.Itoa(port)

	h, err := url.Parse(scheme)
	if err != nil {
		log.Fatal(err)
	}

	cfg := Config{
		WgPrivateKey: privKey,
		Host: h,
		TUNName: tun.TunName(),
		IgonoreTUNs: []string{},
	}

	b, err := json.MarshalIndent(cfg, "", "\t")
	if err != nil {
		log.Fatal(err)
	}

	if err = utils.AtomicWriteFile(path, b, 0600); err != nil {
		log.Fatal(err)
	}

	return &cfg
}

func GetClientConfig(path string, host string, port int) *Config {
	b, err := ioutil.ReadFile(path)
	switch {
	case errors.Is(err, os.ErrNotExist):
		return newClientConfig(path, host, port)
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

