package core

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"strconv"

	"github.com/Notch-Technologies/wizy/cmd/wissy/tun"
	"github.com/Notch-Technologies/wizy/types/key"
	"github.com/Notch-Technologies/wizy/utils"
	"github.com/Notch-Technologies/wizy/wislog"
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
	path string, serverHost string, serverPort int,
	signalHost string, signalPort int,
	wl *wislog.WisLog,
) (*ClientCore, error) {
	serverURL := serverHost + ":" + strconv.Itoa(serverPort)

	serverHostURL, err := url.Parse(serverURL)
	if err != nil {
		return nil, err
	}

	signalURL := signalHost + ":" + strconv.Itoa(signalPort)

	signalHostURL, err := url.Parse(signalURL)
	if err != nil {
		return nil, err
	}

	return &ClientCore{
		ServerHost:  serverHostURL,
		SignalHost:  signalHostURL,
		IgonoreTUNs: []string{},

		path: path,

		wislog: wl,
	}, nil
}

func (c *ClientCore) writeClientCore(
	path, wgPrivateKey, tunName string,
	ifaceBlackList []string,
) *ClientCore {
	if err := os.MkdirAll(filepath.Dir(path), 0777); err != nil {
		c.wislog.Logger.Fatalf("failed to create directory with %s. because %s", path, err.Error())
	}

	c.WgPrivateKey = wgPrivateKey
	c.TunName = tunName
	c.IfaceBlackList = ifaceBlackList

	b, err := json.MarshalIndent(*c, "", "\t")
	if err != nil {
		panic(err)
	}

	if err = utils.AtomicWriteFile(path, b, 0600); err != nil {
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
		return c.writeClientCore(c.path, privKey, tun.TunName(), []string{tun.TunName(), "tun0"})
	case err != nil:
		c.wislog.Logger.Errorf("%s could not be read. exception error: %s", c.path, err.Error())
		panic(err)
	default:
		var core ClientCore
		if err := json.Unmarshal(b, &core); err != nil {
			c.wislog.Logger.Fatalf("can not read client config file. because %v", err)
		}
		return c.writeClientCore(c.path, core.WgPrivateKey, core.TunName, core.IfaceBlackList)
	}
}
