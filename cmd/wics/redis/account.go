package redis

import (
	"sync"

	"github.com/Notch-Technologies/wizy/cmd/wics/model"
)

// TODO:(shintard) refactor store
type AccountManager interface {
	CreateSetupKey()
	GetPeer(machineKey string) (model.Peer, error)
}

type AccountStore struct {
	redis *RedisClient
	mu    sync.Mutex
}

func NewAccountStore(r *RedisClient) *AccountStore {
	return &AccountStore{
		redis: r,
		mu: sync.Mutex{},
	}
}

func (as *AccountStore) CreateSetupKey() {
	panic("not implement CreateSetupKey")
}

func (as *AccountStore) GetPeer(machineKey string) (*model.Peer, error) {
	var peer model.Peer
	err := as.redis.HGetAll(machineKey, peer)
	if err != nil {
		return nil, err
	}

	return &peer, nil
}
