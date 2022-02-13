package store

// be sure to read the `WritePrivateKey` function before using this structure.
//

import (
	"fmt"
	"log"
	"sync"

	"github.com/Notch-Technologies/wizy/types/key"
)

type ServerManager interface {
	WritePrivateKey() error
	GetStateKey() string
	GetPublicKey() string
}

type ServerStore struct {
	storeManager FileStoreManager
	privateKey   key.WicsServerPrivateState

	mu sync.Mutex
}

func NewServerStore(f FileStoreManager) *ServerStore {
	return &ServerStore{
		storeManager: f,
		mu:           sync.Mutex{},
	}
}

func (s *ServerStore) WritePrivateKey() error {
	// already exists
	stateKey, err := s.storeManager.ReadState(ServerPrivateKeyStateKey)
	if err == nil {
		if err := s.privateKey.UnmarshalText(stateKey); err != nil {
			return fmt.Errorf("unable to unmarshal %s. %v", ServerPrivateKeyStateKey, err)
		}

		log.Println("server private key already exists.")
		return nil
	}

	// create new server private key
	k, err := key.NewServerPrivateKey()
	if err != nil {
		log.Fatal(err)
		return err
	}

	ke, err := k.MarshalText()
	if err != nil {
		log.Fatal(err)
		return err
	}

	if err := s.storeManager.WriteState(ServerPrivateKeyStateKey, ke); err != nil {
		log.Fatalf("error writing server private key to store: %v.", err)
		return err
	}

	s.privateKey = k

	return nil
}

func (s *ServerStore) GetStateKey() string {
	key, err := s.storeManager.ReadState(ServerPrivateKeyStateKey)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	return string(key)
}

func (s *ServerStore) GetPublicKey() string {
	return s.privateKey.PublicKey()
}
