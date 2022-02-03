package repository

import (
	"errors"

	"github.com/Notch-Technologies/wizy/cmd/wics/config"
	"github.com/Notch-Technologies/wizy/cmd/wics/model"
	"github.com/Notch-Technologies/wizy/cmd/wics/redis"
	"github.com/Notch-Technologies/wizy/store"
	"github.com/Notch-Technologies/wizy/types/key"
)

type SetupKeyRepositoryManager interface {
	CreateSetupKey(sub, group, job, network string, 
	permissionType key.PermissionType) (*model.SetupKey, error)
}

type SetupKeyRepository struct {
	redis         *redis.RedisClient
	config        *config.Config
	accountStore  *redis.AccountStore
	serverStore   *store.ServerStore
	userStore     *redis.UserStore
	networkStore  *redis.NetworkStore
	orgGroupStore *redis.OrgGroupStore
	setupKeyStore *redis.SetupKeyStore
}

func NewSetupKeyRepository(
	r *redis.RedisClient, config *config.Config, account *redis.AccountStore,
	server *store.ServerStore, user *redis.UserStore, network *redis.NetworkStore,
	group *redis.OrgGroupStore, setupKey *redis.SetupKeyStore,
) *SetupKeyRepository {
	return &SetupKeyRepository{
		redis:        r,
		config:       config,
		accountStore: account,
		serverStore:  server,
		userStore:    user,
		networkStore: network,
		orgGroupStore: group,
		setupKeyStore: setupKey,
	}
}

func (r *SetupKeyRepository) CreateSetupKey(sub, group, job, network string, 
	permissionType key.PermissionType) (*model.SetupKey, error) {
	var (
		user *model.User
	)
	setupKey, err := key.NewSetupKey(sub, group, job, permissionType)
	if err != nil {
		return nil, err
	}

	err = r.redis.Tx(
		func() error {
			n, err := r.networkStore.CreateNetwork(network)
			if err != nil {
				return err
			}
    		
			g, err := r.orgGroupStore.CreateOrgGroup(group)
			if err != nil {
				return err
			}

			user, err = r.userStore.CreateUser(sub, n.ID, g.ID, permissionType)
			if err != nil {
				return err
			}
			return nil
		},
	)
 
	if err != nil {
		if errors.Is(err, model.ErrUserAlredyExists) {
			t, err := setupKey.KeyType()
			if err != nil {
				return nil, err
			}

			revoked, err := setupKey.IsRevoked()
			if err != nil {
				return nil, err
			}

			setupKey, err := r.setupKeyStore.CreateSetupKey(setupKey.Key, user.ID, t, revoked)
			if err != nil {
				return nil, err
			}

			return setupKey, nil
		}
		return nil, err
	}

	t, err := setupKey.KeyType()
	if err != nil {
		return nil, err
	}

	revoked, err := setupKey.IsRevoked()
	if err != nil {
		return nil, err
	}

	sk, err := r.setupKeyStore.CreateSetupKey(setupKey.Key, user.ID, t, revoked)
	if err != nil {
		return nil, err
	}

	return sk, nil
}
