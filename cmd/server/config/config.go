package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/Notch-Technologies/wizy/utils"
)

type Protocol string

const (
	UDP   Protocol = "udp"
	DTLS  Protocol = "dtls"
	TCP   Protocol = "tcp"
	HTTP  Protocol = "http"
	HTTPS Protocol = "https"
)

type TURNConfig struct {
	TimeBasedCredentials bool
	CredentialsTTL       Duration
	Secret               string
	Turns                []*Host
}

type Host struct {
	URL      string
	Username string
	Password string
}

type Auth0Config struct {
	Audience     string
	Issuer       string
	KeysLocation string
}

type TLSConfig struct {
	Domain   string
	Certfile string
	CertKey  string
}

type ServerConfig struct {
	Stuns      []*Host
	TURNConfig *TURNConfig
	Signal     *Host
	AuthConfig Auth0Config
	TLSConfig  TLSConfig
}

func NewServerConfig(path, domain, certfile, certkey string) *ServerConfig {
	b, err := ioutil.ReadFile(path)
	switch {
	case errors.Is(err, os.ErrNotExist):
		return createServerConfig(path, domain, certfile, certkey)
	case err != nil:
		log.Fatal(err)
		panic("failed to load cofig")
	default:
		var cfg ServerConfig
		if err := json.Unmarshal(b, &cfg); err != nil {
			log.Fatalf("config: %v", err)
		}
		return &cfg
	}
}

func createServerConfig(path, domain, certfile, certkey string) *ServerConfig {
	if err := os.MkdirAll(filepath.Dir(path), 0777); err != nil {
		log.Fatal(err)
	}

	cfg := ServerConfig{
		TLSConfig: TLSConfig{
			Domain:   domain,
			Certfile: certfile,
			CertKey:  certkey,
		},
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
