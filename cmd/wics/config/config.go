package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/Notch-Technologies/wizy/utils"
)

type Duration struct {
	time.Duration
}

func (d Duration) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

func (d *Duration) UnmarshalJSON(b []byte) error {
	var v interface{}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	switch value := v.(type) {
	case float64:
		d.Duration = time.Duration(value)
		return nil
	case string:
		var err error
		d.Duration, err = time.ParseDuration(value)
		if err != nil {
			return err
		}
		return nil
	default:
		return errors.New("invalid duration")
	}
}

// TOOD: (shintard) Refactor Config Scheme
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
	Protocol Protocol
	URL      string
	Username string
	Password string
}

type AuthConfig struct {
	AuthAudience     string
	AuthIssuer       string
	AuthKeysLocation string
}

type TLSConfig struct {
	Domain   string
	Certfile string
	CertKey  string
}

type Config struct {
	Stuns      []*Host
	TURNConfig *TURNConfig
	Signal     *Host
	StorePath  string
	AuthConfig AuthConfig
	TLSConfig  TLSConfig
}

func LoadConfig(path, domain, certfile, certkey string) *Config {
	b, err := ioutil.ReadFile(path)
	switch {
	case errors.Is(err, os.ErrNotExist):
		return newConfig(path, domain, certfile, certkey)
	case err != nil:
		log.Fatal(err)
		panic("failed to load cofig")
	default:
		var cfg Config
		if err := json.Unmarshal(b, &cfg); err != nil {
			log.Fatalf("config: %v", err)
		}
		return &cfg
	}
}

func newConfig(path, domain, certfile, certkey string) *Config {
	if err := os.MkdirAll(filepath.Dir(path), 0777); err != nil {
		log.Fatal(err)
	}

	cfg := Config{
		StorePath: path,
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
