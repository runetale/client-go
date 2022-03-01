package store

import (
	"fmt"
	"log"
	"sync"

	"github.com/Notch-Technologies/wizy/types/key"
)

type ClientManager interface {
	WritePrivateKey() error
	GetState64Key() string
	GetPublicKey() string
}

type ClientStore struct {
	storeManager FileStoreManager
	privateKey   key.WicsClientPrivateState

	mu sync.Mutex
}

func NewClientStore(f FileStoreManager) *ClientStore {
	return &ClientStore{
		storeManager: f,
		mu:           sync.Mutex{},
	}
}

func (c *ClientStore) WritePrivateKey() error {
	// already exists
	stateKey, err := c.storeManager.ReadState(ClientPrivateKeyStateKey)
	if err == nil {
		if err := c.privateKey.UnmarshalText(stateKey); err != nil {
			return fmt.Errorf("unable to unmarshal %s. %v", ClientPrivateKeyStateKey, err)
		}

		log.Println("client private key already exists.")
		return nil
	}

	// create new server private key
	k, err := key.NewClientPrivateKey()
	if err != nil {
		log.Fatal(err)
		return err
	}

	ke, err := k.MarshalText()
	if err != nil {
		log.Fatal(err)
		return err
	}

	if err := c.storeManager.WriteState(ClientPrivateKeyStateKey, ke); err != nil {
		log.Fatalf("error writing client private key to store: %v.", err)
		return err
	}

	c.privateKey = k

	return nil
}

func (c *ClientStore) GetStateKey() string {
	key, err := c.storeManager.ReadState(ClientPrivateKeyStateKey)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	return string(key)
}

func (c *ClientStore) GetPublicKey() string {
	return c.privateKey.PublicKey()
}

func (c *ClientStore) GetPrivateKey() string {
	return c.privateKey.PrivateKey()
}
