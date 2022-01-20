package store

// be sure to read the `WritePrivateKey` function before using this structure!
//

import (
	"fmt"
	"log"
	"sync"

	"github.com/Notch-Technologies/wizy/types/key"
)

type ServerManager interface {
	WritePrivateKey() error
	// Client側からSetupKeyとClientMachineKeyを使用してサーバーからServerPrivateKeyのPublicKeyをもらう時に使う
	GetPublicKey() string
	GetBase64Key() string
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
			return fmt.Errorf("invalid key in %s key of %v: %w", ServerPrivateKeyStateKey, s.storeManager, err)
		}

		if s.privateKey.IsZero() {
			return fmt.Errorf("invalid zero key stored in %v key of %v", ServerPrivateKeyStateKey, s.storeManager)
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

func (s *ServerStore) GetPublicKey() string {
	key, err := s.storeManager.ReadState(ServerPrivateKeyStateKey)
	if err != nil {
		log.Fatal(err)
		return ""
	}
	return string(key)
}

func (s *ServerStore) GetBase64Key() string {
	return s.privateKey.String()
}
