package conf

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/Notch-Technologies/dotshake/dotlog"
	"github.com/Notch-Technologies/dotshake/tun"
	"github.com/Notch-Technologies/dotshake/types/key"
	"github.com/Notch-Technologies/dotshake/utils"
)

type ClientConf struct {
	WgPrivateKey string   `json:"wg_private_key"`
	ServerHost   string   `json:"server_host"`
	ServerPort   uint     `json:"server_port"`
	TunName      string   `json:"tun"`
	PreSharedKey string   `json:"preshared_key"`
	BlackList    []string `json:"blacklist"`

	path string

	dotlog *dotlog.DotLog
}

func NewClientConf(
	path string,
	serverHost string, serverPort uint,
	dl *dotlog.DotLog,
) (*ClientConf, error) {

	return &ClientConf{
		ServerHost: serverHost,
		ServerPort: serverPort,

		path: path,

		dotlog: dl,
	}, nil
}

func (c *ClientConf) writeClientConf(
	wgPrivateKey, tunName string,
	serverHost string,
	serverPort uint,
	blackList []string,
	presharedKey string,
) *ClientConf {
	if err := os.MkdirAll(filepath.Dir(c.path), 0777); err != nil {
		c.dotlog.Logger.Fatalf("failed to create directory with %s. because %s", c.path, err.Error())
	}

	c.ServerHost = serverHost
	c.ServerPort = serverPort
	c.WgPrivateKey = wgPrivateKey
	c.TunName = tunName
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

		return c.writeClientConf(
			privKey,
			tun.TunName(),
			c.ServerHost,
			c.ServerPort,
			[]string{tun.TunName(), "tun0"},
			"",
		)
	case err != nil:
		c.dotlog.Logger.Errorf("%s could not be read. exception error: %s", c.path, err.Error())
		panic(err)
	default:
		var core ClientConf
		if err := json.Unmarshal(b, &core); err != nil {
			c.dotlog.Logger.Fatalf("can not read client config file. because %v", err)
		}

		return c.writeClientConf(
			core.WgPrivateKey,
			core.TunName,
			c.ServerHost,
			c.ServerPort,
			core.BlackList,
			"",
		)
	}
}

// format like this => 127.0.0.1:443
//
func (c *ClientConf) GetServerHost() string {
	port := strconv.Itoa(int(c.ServerPort))
	host := strings.Replace(c.ServerHost, "http://", "", -1)
	return host + ":" + port
}
