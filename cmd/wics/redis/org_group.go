package redis

import (
	"encoding/json"
	"errors"

	"github.com/Notch-Technologies/wizy/cmd/wics/model"
	"github.com/google/uuid"
)

type OrgGroupStoreManager interface {

}

type OrgGroupStore struct {
	redis *RedisClient
}

func NewOrgGroupStore(r *RedisClient) *OrgGroupStore {
	return &OrgGroupStore{
		redis: r,
	}
}

func (os *OrgGroupStore) CreateOrgGroup(name string) (*model.OrgGroup, error) {
	u, err := os.FindByGroupName(name)
	if err != nil {
		if errors.Is(err, model.ErrGroupAlredyExists) {
			return u, nil
		}
		return nil, err
	}

	uid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	group := model.NewOrgGroup(uid.String(), name)

	gm, err := os.getGroups()
	if err != nil {
		return nil, err
	}

	gm[uid.String()] = group

	if err := os.setGroup(gm); err != nil {
		return nil, err
	}

	return group, nil
}

func (os *OrgGroupStore) FindByGroupName(name string) (*model.OrgGroup, error) {
	gm, err := os.getGroups()
	if err != nil {
		return nil, err
	}

	if gm[name] != nil {
		return gm[name], model.ErrGroupAlredyExists
	}

	return nil, nil
}

func (os *OrgGroupStore) setGroup(gm map[string]*model.OrgGroup) error {
	bytes, err := json.Marshal(gm)
	if err != nil {
		return err
	}

	err = os.redis.Set(string(orgGroupStoreKey), bytes, 0)
	if err != nil {
		return err
	}

	return nil
}

func (os *OrgGroupStore) getGroups() (map[string]*model.OrgGroup, error) {
	gm := make(map[string]*model.OrgGroup)

	exists, err := os.redis.Exists(string(orgGroupStoreKey))
	if err != nil {
		return nil, err
	}
	if exists == 0 {
		return gm, nil
	}

	bytes, err := os.redis.Get(string(orgGroupStoreKey))
	if err != nil {
		return nil, err
	}

	if bytes == nil {
		return gm, err
	}

	if err := json.Unmarshal(bytes, &gm); err != nil {
		return nil, err
	}

	return gm, nil
}
