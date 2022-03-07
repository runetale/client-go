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

type TURNConfig struct {
	Turns                []*Host
	CredentialsTTL       Duration
	Secret               string
	TimeBasedCredentials bool
}

type Host struct {
	URL      string
	Username *string
	Password *string
}

type JwtConfig struct {
	Aud          string
	Iss          string
	KeysLocation string
	Secret       string
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
	JwtConfig  JwtConfig
	TLSConfig  TLSConfig
}

func NewServerConfig(path, domain, certfile, certkey string) *ServerConfig {
	b, err := ioutil.ReadFile(path)

	switch {
	case errors.Is(err, os.ErrNotExist):
		return writeServerConfig(path, domain, certfile, certkey)
	case err != nil:
		log.Fatalf("failed to load config for server. because %s", err.Error())
		panic(err)
	default:
		var cfg ServerConfig
		if err := json.Unmarshal(b, &cfg); err != nil {
			log.Fatalf("failed to unmarshall server config file. becasue %s", err.Error())
		}
		return &cfg
	}
}

func writeServerConfig(path, domain, certfile, certkey string) *ServerConfig {
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
		log.Fatalf("failed to marshall indent server config file. because %s", err.Error())
	}

	if err = utils.AtomicWriteFile(path, b, 0600); err != nil {
		log.Fatalf("failed to write server config file. because %s", err.Error())
	}

	return &cfg
}
