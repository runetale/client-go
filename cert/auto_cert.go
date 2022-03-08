package cert

import (
	"os"

	"golang.org/x/crypto/acme/autocert"
)

type CertConfig struct {
	Dir string
	File string
	Key string
	Domain string
}

func NewCertConfig(certDir, domain, secret, file string) *CertConfig {
	return &CertConfig{
		Dir: certDir,
		File: file,
		Key: secret,
		Domain: domain,
	}
}

func (c *CertConfig) CreateCertManager() *autocert.Manager {
	if _, err := os.Stat(c.Dir); os.IsNotExist(err) {
		err = os.MkdirAll(c.Dir, os.ModeDir)
		if err != nil {
		}
	}

	certManager := &autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		Cache:      autocert.DirCache(c.Dir),
		HostPolicy: autocert.HostWhitelist(c.Domain),
	}

	return certManager
}
