package redis

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/Notch-Technologies/wizy/cmd/wics/model"
	"github.com/Notch-Technologies/wizy/types/key"
	"github.com/google/uuid"
)

type UserStoreManager interface {
	CreateUser(providerID string, networkID, groupID string,
		permission key.PermissionType, setupKey string) (*model.User, error)
	FindByProviderID(providerID string) (*model.User, error)
}

type UserStore struct {
	redis *RedisClient
}

func NewUserStore(r *RedisClient) *UserStore {
	return &UserStore{
		redis: r,
	}
}

func (us *UserStore) CreateUser(providerID string, networkid, groupid string,
	permission key.PermissionType) (*model.User, error) {
	i := strings.Index(providerID, "|")
	provider := providerID[:i]
	id := providerID[i+1:]

	// check if the user exists.
	//
	u, err := us.FindByProviderID(id)
	if err != nil {
		if errors.Is(err, model.ErrUserAlredyExists) {
			return u, err
		}
		return nil, err
	}

	// create user
	//
	um, err := us.getUsers()
	if err != nil {
		return nil, err
	}

	uid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	user := model.NewUser(uid.String(), providerID, provider, networkid, groupid, permission)

	um[id] = user

	//if err := us.setUser(um); err != nil {
	//	return nil, err
	//}

	return user, nil
}

func (us *UserStore) FindByProviderID(providerID string) (*model.User, error) {
	um, err := us.getUsers()
	if err != nil {
		return nil, err
	}

	if um[providerID] != nil {
		return um[providerID], model.ErrUserAlredyExists
	}

	return nil, nil
}

func (us *UserStore) setUser(um map[string]*model.User) error {
	bytes, err := json.Marshal(um)
	if err != nil {
		return err
	}

	err = us.redis.Set(string(UserStoreKey), bytes, 0)
	if err != nil {
		return err
	}

	return nil
}

func (us *UserStore) getUsers() (map[string]*model.User, error) {
	um := make(map[string]*model.User)

	exists, err := us.redis.Exists(string(UserStoreKey))
	if err != nil {
		return nil, err
	}
	if exists == 0 {
		return um, nil
	}

	bytes, err := us.redis.Get(string(UserStoreKey))
	if err != nil {
		return nil, err
	}

	if bytes == nil {
		return um, err
	}

	if err := json.Unmarshal(bytes, &um); err != nil {
		return nil, err
	}

	return um, nil
}
