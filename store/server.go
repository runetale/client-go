package store

import (
	"log"
	"sync"

	"github.com/Notch-Technologies/wizy/types/key"
)

type ServerManager interface {
	WritePrivateKey() error
}

type Server struct {
	storeManager FileStoreManager
	mu               sync.Mutex
}

func NewServer(f FileStoreManager) *Server {
	return &Server{
		storeManager: f,
		mu:               sync.Mutex{},
	}
}

func (s *Server) WritePrivateKey() error {
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
		log.Fatalf("error writing server private key to store: %v", err)
		return err
	}

	return nil
}
