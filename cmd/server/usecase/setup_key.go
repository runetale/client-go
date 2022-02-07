package usecase

import (
	"errors"
	"strings"

	"github.com/Notch-Technologies/wizy/cmd/server/database"
	"github.com/Notch-Technologies/wizy/cmd/server/domain"
	"github.com/Notch-Technologies/wizy/cmd/server/repository"
	"github.com/Notch-Technologies/wizy/types/key"
)

type SetupKeyUsecaseManager interface {
	CreateSetupKey(networkName, name string, permission key.PermissionType, providerID string) error
}

type SetupKeyUsecase struct {
	setupKeyRepository  *repository.SetupKeyRepository
	userRepository      *repository.UserRepository
	orgRepository       *repository.OrgRepository
	networkRepository   *repository.NetworkRepository
	userGroupRepository *repository.UserGroupRepository
	jobRepository       *repository.JobRepository
}

func NewSetupKeyUsecase(
	db *database.Sqlite,
) *SetupKeyUsecase {
	return &SetupKeyUsecase{
		setupKeyRepository:  repository.NewSetupKeyRepository(db),
		userRepository:      repository.NewUserRepository(db),
		orgRepository:       repository.NewOrgRepository(db),
		networkRepository:   repository.NewNetworkRepository(db),
		userGroupRepository: repository.NewUserGroupRepository(db),
		jobRepository:       repository.NewJobRepository(db),
	}
}

func (s *SetupKeyUsecase) CreateSetupKey(networkName, userGroupName uint64, jobName, orgID string,
	permission key.PermissionType, sub string) (*key.SetupKey, error) {
	orgGroup, err := s.orgRepository.FindByOrganizationID(orgID)
	if err != nil {
		return nil, err
	}

	// TODO: FindByNetworkID
	// when if not found, create default network

	network := domain.NewNetwork(networkName, "", "", "")
	err = s.networkRepository.CreateNetwork(network)
	if err != nil {
		return nil, err
	}

	// TODO: FindByUserGroupID
	// when if not found, create default user group
	userGroup := domain.NewUserGroup(userGroupName, permission)
	err = s.userGroupRepository.CreateUserGroup(userGroup)
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

	setupKey, err := key.NewSetupKey(user.ID, sub, job.Name, userGroup.ID, orgGroup.ID, permission)
	if err != nil {
		return nil, err
	}

	return setupKey, nil
}
