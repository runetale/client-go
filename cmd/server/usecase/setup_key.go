package usecase

import (
	"errors"
	"strings"

	"github.com/Notch-Technologies/wizy/cmd/server/config"
	"github.com/Notch-Technologies/wizy/cmd/server/database"
	"github.com/Notch-Technologies/wizy/cmd/server/domain"
	"github.com/Notch-Technologies/wizy/cmd/server/repository"
	"github.com/Notch-Technologies/wizy/types/key"
)

type SetupKeyUsecaseCaller interface {
	CreateSetupKey(networkID, userGroupID, jobID, roleID uint,
		orgID string, sub, email string) (*key.SetupKey, error)
}

type SetupKeyUsecase struct {
	setupKeyRepository     repository.SetupKeyRepositoryCaller
	userRepository         repository.UserRepositoryCaller
	networkRepository      repository.NetworkRepositoryCaller
	userGroupRepository    repository.UserGroupRepositoryCaller
	adminNetworkRepository repository.AdminNetworkRepositoryCaller
	jobRepository          repository.JobRepositoryCaller
	roleRepository         repository.RoleRepositoryCaller
	config                 *config.ServerConfig
}

func NewSetupKeyUsecase(
	db database.SQLExecuter, config *config.ServerConfig,
) SetupKeyUsecaseCaller {
	return &SetupKeyUsecase{
		setupKeyRepository:     repository.NewSetupKeyRepository(db),
		userRepository:         repository.NewUserRepository(db),
		networkRepository:      repository.NewNetworkRepository(db),
		userGroupRepository:    repository.NewUserGroupRepository(db),
		adminNetworkRepository: repository.NewAdminNetworkRepository(db),
		jobRepository:          repository.NewJobRepository(db),
		roleRepository:         repository.NewRoleRepository(db),
		config:                 config,
	}
}

// TOOD: (shintard) allows network use cases to dynamically specify CIDR and IP address
// ranges and create networks
//
func (s *SetupKeyUsecase) CreateSetupKey(
	networkID, userGroupID, jobID, roleID uint,
	orgID string, sub, email string,
) (*key.SetupKey, error) {
	adminNetwork, err := s.adminNetworkRepository.FindByOrganizationID(orgID)
	if err != nil {
		return nil, err
	}

	network, err := s.networkRepository.FindByNetworkID(networkID)
	if err != nil {
		return nil, err
	}

	userGroup, err := s.userGroupRepository.FindByUserGroupID(userGroupID)
	if err != nil {
		return nil, err
	}

	role, err := s.roleRepository.FindByID(roleID)
	if err != nil {
		return nil, err
	}

	i := strings.Index(sub, "|")
	provider := sub[:i]
	user := domain.NewUser(sub, provider, email, network.ID, userGroup.ID, adminNetwork.ID, role.ID)

	err = s.userRepository.CreateUser(user)
	if err != nil && !errors.Is(err, domain.ErrUserAlredyExists) {
		return nil, err
	}

	job, err := s.jobRepository.FindByID(jobID)
	if err != nil {
		return nil, err
	}

	sk, err := key.NewSetupKey(
		user.ID, sub, job.Name, userGroup.ID, adminNetwork.ID,
		role.Permission, s.config.JwtConfig.Iss, s.config.JwtConfig.Aud,
		s.config.JwtConfig.Secret,
	)
	if err != nil {
		return nil, err
	}

	keyType, err := sk.KeyType()
	if err != nil {
		return nil, err
	}

	_, err = sk.IsRevoked()
	if err != nil {
		return nil, err
	}

	setupKey := domain.NewSetupKey(adminNetwork.ID, user.ID, sk.Key, keyType)

	err = s.setupKeyRepository.CreateSetupKey(setupKey)
	if err != nil {
		return nil, err
	}

	return sk, nil
}
