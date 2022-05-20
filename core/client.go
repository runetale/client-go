package core

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

type ClientCore struct {
	WgPrivateKey   string
	ServerHost     *url.URL
	SignalHost     *url.URL
	TunName        string
	PreSharedKey   string
	IfaceBlackList []string

	path string

	dotlog *dotlog.DotLog
}

func NewClientCore(
	path string,
	serverHost string, serverPort int,
	signalHost string, signalPort int,
	dl *dotlog.DotLog,
) (*ClientCore, error) {
	var serverHostURL *url.URL
	var signalHostURL *url.URL
	var err error

	if serverHost != "" {
		serverURL := serverHost + ":" + strconv.Itoa(serverPort)
		serverHostURL, err = url.Parse(serverURL)
		if err != nil {
			return nil, err
		}
	}

	if signalHost != "" {
		signalURL := signalHost + ":" + strconv.Itoa(signalPort)
		signalHostURL, err = url.Parse(signalURL)
		if err != nil {
			return nil, err
		}
	}

	return &ClientCore{
		ServerHost: serverHostURL,
		SignalHost: signalHostURL,

		path: path,

		dotlog: dl,
	}, nil
}

func (c *ClientCore) writeClientCore(
	wgPrivateKey, tunName string,
	serverHost, signalHost *url.URL,
	ifaceBlackList []string,
) *ClientCore {
	if err := os.MkdirAll(filepath.Dir(c.path), 0777); err != nil {
		c.dotlog.Logger.Fatalf("failed to create directory with %s. because %s", c.path, err.Error())
	}

	c.WgPrivateKey = wgPrivateKey
	c.TunName = tunName
	c.ServerHost = serverHost
	c.SignalHost = signalHost
	c.IfaceBlackList = ifaceBlackList

	b, err := json.MarshalIndent(*c, "", "\t")
	if err != nil {
		panic(err)
	}

	if err = utils.AtomicWriteFile(c.path, b, 0600); err != nil {
		panic(err)
	}

	return c
}

func (c *ClientCore) GetClientCore() *ClientCore {
	b, err := ioutil.ReadFile(c.path)
	switch {
	case errors.Is(err, os.ErrNotExist):
		privKey, err := key.NewGenerateKey()
		if err != nil {
			c.dotlog.Logger.Error("failed to generate key for wireguard")
			panic(err)
		}
		return c.writeClientCore(privKey, tun.TunName(), c.ServerHost, c.SignalHost, []string{tun.TunName(), "tun0"})
	case err != nil:
		c.dotlog.Logger.Errorf("%s could not be read. exception error: %s", c.path, err.Error())
		panic(err)
	default:
		var core ClientCore
		if err := json.Unmarshal(b, &core); err != nil {
			c.dotlog.Logger.Fatalf("can not read client config file. because %v", err)
		}
		return c.writeClientCore(core.WgPrivateKey, core.TunName, c.ServerHost, c.SignalHost, core.IfaceBlackList)
	}
}
