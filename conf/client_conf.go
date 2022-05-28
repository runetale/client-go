package conf

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"strconv"

	"github.com/Notch-Technologies/dotshake/dotlog"
	"github.com/Notch-Technologies/dotshake/tun"
	"github.com/Notch-Technologies/dotshake/types/key"
	"github.com/Notch-Technologies/dotshake/utils"
)

type ClientConf struct {
	WgPrivateKey string
	ServerHost   *url.URL
	TunName      string
	PreSharedKey string
	BlackList    []string

	path string

	dotlog *dotlog.DotLog
}

func NewClientConf(
	path string,
	serverHost string, serverPort int,
	dl *dotlog.DotLog,
) (*ClientConf, error) {
	var serverHostURL *url.URL
	var err error

	if serverHost != "" {
		serverURL := serverHost + ":" + strconv.Itoa(serverPort)
		serverHostURL, err = url.Parse(serverURL)
		if err != nil {
			return nil, err
		}
	}

	return &ClientConf{
		ServerHost: serverHostURL,

		path: path,

		dotlog: dl,
	}, nil
}

func (c *ClientConf) writeClientConf(
	wgPrivateKey, tunName string,
	serverHost *url.URL,
	blackList []string,
) *ClientConf {
	if err := os.MkdirAll(filepath.Dir(c.path), 0777); err != nil {
		c.dotlog.Logger.Fatalf("failed to create directory with %s. because %s", c.path, err.Error())
	}

	c.WgPrivateKey = wgPrivateKey
	c.TunName = tunName
	c.ServerHost = serverHost
	c.BlackList = blackList

	b, err := json.MarshalIndent(*c, "", "\t")
	if err != nil {
		panic(err)
	}

	if err = utils.AtomicWriteFile(c.path, b, 0600); err != nil {
		panic(err)
	}

	return c
}

func (c *ClientConf) GetClientConf() *ClientConf {
	b, err := ioutil.ReadFile(c.path)
	switch {
	case errors.Is(err, os.ErrNotExist):
		privKey, err := key.NewGenerateKey()
		if err != nil {
			c.dotlog.Logger.Error("failed to generate key for wireguard")
			panic(err)
		}
		return c.writeClientConf(privKey, tun.TunName(), c.ServerHost, []string{tun.TunName(), "tun0"})
	case err != nil:
		c.dotlog.Logger.Errorf("%s could not be read. exception error: %s", c.path, err.Error())
		panic(err)
	default:
		var core ClientConf
		if err := json.Unmarshal(b, &core); err != nil {
			c.dotlog.Logger.Fatalf("can not read client config file. because %v", err)
		}
		return c.writeClientConf(core.WgPrivateKey, core.TunName, c.ServerHost, core.BlackList)
	}
}
