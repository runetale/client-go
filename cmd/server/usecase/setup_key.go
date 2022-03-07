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
	CreateSetupKey(networkID, userGroupID uint, jobName, orgID string,
		permission key.PermissionType, sub string) (*key.SetupKey, error)
}

type SetupKeyUsecase struct {
	setupKeyRepository  *repository.SetupKeyRepository
	userRepository      *repository.UserRepository
	orgRepository       *repository.OrgRepository
	networkRepository   *repository.NetworkRepository
	userGroupRepository *repository.UserGroupRepository
	jobRepository       *repository.JobRepository
	config 				*config.ServerConfig
}

func NewSetupKeyUsecase(
	db database.SQLExecuter, config *config.ServerConfig,
) SetupKeyUsecaseCaller {
	return &SetupKeyUsecase{
		setupKeyRepository:  repository.NewSetupKeyRepository(db),
		userRepository:      repository.NewUserRepository(db),
		orgRepository:       repository.NewOrgRepository(db),
		networkRepository:   repository.NewNetworkRepository(db),
		userGroupRepository: repository.NewUserGroupRepository(db),
		jobRepository:       repository.NewJobRepository(db),
		config: 			 config,
	}
}

// TOOD: (shintard) allows network use cases to dynamically specify CIDR and IP address
// ranges and create networks
//
func (s *SetupKeyUsecase) CreateSetupKey(networkID, userGroupID uint, jobName, orgID string,
	permission key.PermissionType, sub string) (*key.SetupKey, error) {
	orgGroup, err := s.orgRepository.FindByOrganizationID(orgID)
	if err != nil {
		return nil, err
	}

	network, err := s.networkRepository.FindByNetworkID(networkID)
	if errors.Is(err, domain.ErrNoRows) {
		// create a network in the range of default network 100.64.0.0/10 if the network does not exist
		//
		network = domain.NewNetwork("default", "100.64.0.0", 10, "")
		err = s.networkRepository.CreateNetwork(network)
		if err != nil {
			return nil, err
		}
	}
	if err != nil {
		return nil, err
	}

	userGroup, err := s.userGroupRepository.FindByUserGroupID(userGroupID)
	if errors.Is(err, domain.ErrNoRows) {
		userGroup = domain.NewUserGroup("default", permission)
		err = s.userGroupRepository.CreateUserGroup(userGroup)
		if err != nil {
			return nil, err
		}
	}
	if err != nil {
		return nil, err
	}

	i := strings.Index(sub, "|")
	provider := sub[:i]
	user := domain.NewUser(sub, provider, network.ID, userGroup.ID, orgGroup.ID, permission)

	err = s.userRepository.CreateUser(user)
	if err != nil && !errors.Is(err, domain.ErrUserAlredyExists) {
		return nil, err
	}

	job := domain.NewJob(jobName, user.ID, orgGroup.ID)
	err = s.jobRepository.CreateJob(job)
	if err != nil {
		return nil, err
	}

	sk, err := key.NewSetupKey(
		user.ID, sub, job.Name, userGroup.ID, orgGroup.ID, 
		permission, s.config.JwtConfig.Iss, s.config.JwtConfig.Aud,
		s.config.JwtConfig.Secret,
	)
	if err != nil {
		return nil, err
	}

	keyType, err := sk.KeyType()
	if err != nil {
		return nil, err
	}

	revoked, err := sk.IsRevoked()
	if err != nil {
		return nil, err
	}

	setupKey := domain.NewSetupKey(user.ID, sk.Key, keyType, revoked)

	err = s.setupKeyRepository.CreateSetupKey(setupKey)
	if err != nil {
		return nil, err
	}

	return sk, nil
}
