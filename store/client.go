package store

import (
	"fmt"
	"log"
	"sync"

	"github.com/Notch-Technologies/dotshake/dotlog"
	"github.com/Notch-Technologies/dotshake/types/key"
)

type ClientManager interface {
	GetPrivateKey() string
	GetPublicKey() string
}

type ClientStore struct {
	storeManager FileStoreManager
	privateKey   key.WicsClientPrivateState
	dotlog       *dotlog.DotLog

	mu sync.Mutex
}

// client Store initialization method.
//
func NewClientStore(f FileStoreManager, dotlog *dotlog.DotLog) *ClientStore {
	return &ClientStore{
		storeManager: f,
		dotlog:       dotlog,

		mu: sync.Mutex{},
	}
}

// read the PrivateKey from the Client State, and if it does not exist, write a new one.
//
func (c *ClientStore) WritePrivateKey() error {
	stateKey, err := c.storeManager.ReadState(ClientPrivateKeyStateKey)
	if err == nil {
		if err := c.privateKey.UnmarshalText(stateKey); err != nil {
			return fmt.Errorf("unable to unmarshal %s. %v", ClientPrivateKeyStateKey, err)
		}

		c.dotlog.Logger.Debugf("client private key already exists")
		return nil
	}

	// create new client private key
	k, err := key.NewClientPrivateKey()
	if err != nil {
		return err
	}

	ke, err := k.MarshalText()
	if err != nil {
		log.Fatal(err)
		return err
	}

	// write new client private key
	if err := c.storeManager.WriteState(ClientPrivateKeyStateKey, ke); err != nil {
		c.dotlog.Logger.Errorf("error writing client private key to store: %v.", err)
		return err
	}

	c.privateKey = k
	c.dotlog.Logger.Debugf("write new client private key")

	return nil
}

func (c *ClientStore) GetPublicKey() string {
	return c.privateKey.PublicKey()
}

func (c *ClientStore) GetPrivateKey() string {
	return c.privateKey.PrivateKey()
}
