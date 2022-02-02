package redis

import (
	"encoding/json"
	"errors"

	"github.com/Notch-Technologies/wizy/cmd/wics/model"
	"github.com/google/uuid"
)

type NetworkStoreManager interface {
	CreateNetwork() *NetworkStore
}

type NetworkStore struct {
	redis *RedisClient
}

func NewNetworkStore(r *RedisClient) *NetworkStore {
	return &NetworkStore{
		redis: r,
	}
}

func (ns *NetworkStore) CreateNetwork(name string) (*model.Network, error) {
	u, err := ns.FindByName(name)
	if err != nil {
		if errors.Is(err, model.ErrNetworkAlredyExists) {
			return u, nil
		}
		return nil, err
	}

	uid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	network := model.NewNetwork(uid.String(), name, "", "", "")

	nm, err := ns.getNetworks()
	if err != nil {
		return nil, err
	}

	nm[uid.String()] = network

	if err := ns.setNetwork(nm); err != nil {
		return nil, err
	}

	return network, nil
}

func (ns *NetworkStore) FindByName(name string) (*model.Network, error) {
	nm, err := ns.getNetworks()
	if err != nil {
		return nil, err
	}

	if nm[name] != nil {
		return nm[name], model.ErrNetworkAlredyExists
	}

	return nil, nil
}

func (us *NetworkStore) setNetwork(nm map[string]*model.Network) error {
	bytes, err := json.Marshal(nm)
	if err != nil {
		return err
	}

	err = us.redis.Set(string(networkStoreKey), bytes, 0)
	if err != nil {
		return err
	}

	return nil
}

func (ns *NetworkStore) getNetworks() (map[string]*model.Network, error) {
	nm := make(map[string]*model.Network)

	exists, err := ns.redis.Exists(string(networkStoreKey))
	if err != nil {
		return nil, err
	}
	if exists == 0 {
		return nm, nil
	}

	bytes, err := ns.redis.Get(string(networkStoreKey))
	if err != nil {
		return nil, err
	}

	if bytes == nil {
		return nm, err
	}

	if err := json.Unmarshal(bytes, &nm); err != nil {
		return nil, err
	}

	return nm, nil
}
