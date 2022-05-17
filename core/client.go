package core

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"strconv"

	"github.com/Notch-Technologies/dotshake/cmd/dotshake/tun"
	"github.com/Notch-Technologies/dotshake/types/key"
	"github.com/Notch-Technologies/dotshake/utils"
	"github.com/Notch-Technologies/dotshake/wislog"
)

type ClientCore struct {
	WgPrivateKey   string
	ServerHost     *url.URL
	SignalHost     *url.URL
	IgonoreTUNs    []string
	TunName        string
	PreSharedKey   string
	IfaceBlackList []string

	path string

	wislog *wislog.WisLog
}

func NewClientCore(
	path string,
	serverHost string, serverPort int,
	signalHost string, signalPort int,
	wl *wislog.WisLog,
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

		wislog: wl,
	}, nil
}

func (c *ClientCore) writeClientCore(
	wgPrivateKey, tunName string,
	serverHost, signalHost *url.URL,
	igonoreTUNs, ifaceBlackList []string,
) *ClientCore {
	if err := os.MkdirAll(filepath.Dir(c.path), 0777); err != nil {
		c.wislog.Logger.Fatalf("failed to create directory with %s. because %s", c.path, err.Error())
	}

	c.WgPrivateKey = wgPrivateKey
	c.TunName = tunName
	c.ServerHost = serverHost
	c.SignalHost = signalHost
	c.IgonoreTUNs = igonoreTUNs
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
			c.wislog.Logger.Error("failed to generate key for wireguard")
			panic(err)
		}
		return c.writeClientCore(privKey, tun.TunName(), c.ServerHost, c.SignalHost, []string{}, []string{tun.TunName(), "tun0"})
	case err != nil:
		c.wislog.Logger.Errorf("%s could not be read. exception error: %s", c.path, err.Error())
		panic(err)
	default:
		var core ClientCore
		if err := json.Unmarshal(b, &core); err != nil {
			c.wislog.Logger.Fatalf("can not read client config file. because %v", err)
		}
		return c.writeClientCore(core.WgPrivateKey, core.TunName, core.ServerHost, core.SignalHost, core.IgonoreTUNs, core.IfaceBlackList)
	}
}
