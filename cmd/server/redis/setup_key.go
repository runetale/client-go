package redis

import (
	"encoding/json"

	"github.com/Notch-Technologies/wizy/cmd/server/model"
	"github.com/Notch-Technologies/wizy/types/key"
	"github.com/google/uuid"
)

type SetupKeyStoreManager interface {
	CreateSetupKey(key, userID string,
		keyType key.SetupKeyType, revoked bool) (*model.SetupKey, error)
}

type SetupKeyStore struct {
	redis *RedisClient
}

func NewSetupKeyStore(r *RedisClient) *SetupKeyStore {
	return &SetupKeyStore{
		redis: r,
	}
}

func (ss *SetupKeyStore) CreateSetupKey(key, userID string,
	keyType key.SetupKeyType, revoked bool) (*model.SetupKey, error) {
	uid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	setupKey := model.NewSetupKey(uid.String(), userID, key, keyType, revoked)

	sm, err := ss.getSetupKeys()
	if err != nil {
		return nil, err
	}

	sm[key] = setupKey

	if err := ss.setSetupKey(sm); err != nil {
		return nil, err
	}

	return setupKey, nil
}

func (ss *SetupKeyStore) setSetupKey(gm map[string]*model.SetupKey) error {
	bytes, err := json.Marshal(gm)
	if err != nil {
		return err
	}

	err = ss.redis.Set(string(SetupKeyStoreKey), bytes, 0)
	if err != nil {
		return err
	}

	return nil
}

func (ss *SetupKeyStore) getSetupKeys() (map[string]*model.SetupKey, error) {
	sm := make(map[string]*model.SetupKey)

	exists, err := ss.redis.Exists(string(SetupKeyStoreKey))
	if err != nil {
		return nil, err
	}
	if exists == 0 {
		return sm, nil
	}

	bytes, err := ss.redis.Get(string(SetupKeyStoreKey))
	if err != nil {
		return nil, err
	}

	if bytes == nil {
		return sm, err
	}

	if err := json.Unmarshal(bytes, &sm); err != nil {
		return nil, err
	}

	return sm, nil
}
