package config

import (
	"time"
	"encoding/json"
	"errors"
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

type Protocol string

const (
	UDP   Protocol = "udp"
	DTLS  Protocol = "dtls"
	TCP   Protocol = "tcp"
	HTTP  Protocol = "http"
	HTTPS Protocol = "https"
)

type Config struct {
	Stuns []*Host
	TURNConfig *TURNConfig
	Signal *Host
	StorePath string
	Auth Auth
	TLSConfig  TLSConfig
}

type TURNConfig  struct {
	TimeBasedCredentials bool
	CredentialsTTL Duration
	Secret string
	Turns []*Host
}

type Host struct {
    Protocol Protocol
    URL      string
    Username string
    Password string
}

type Auth struct {
	AuthAudience string
	AuthIssuer string
	AuthKeysLocation string
}

type TLSConfig struct {
	Domain string
	Certfile string
	CertKey string
}
