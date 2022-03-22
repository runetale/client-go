package usecase

import (
	"strings"

	client "github.com/Notch-Technologies/wizy/auth0"
	"github.com/Notch-Technologies/wizy/cmd/server/database"
	"github.com/Notch-Technologies/wizy/cmd/server/pb/organization"
	"github.com/Notch-Technologies/wizy/cmd/server/repository"
)

type NetworkUsecaseCaller interface {
	GetNetwork(sub, organizationID string) (*organization.GetNetworkResponse, error)
}

type NetworkUsecase struct {
	adminNetworkRepository repository.AdminNetworkRepositoryCaller
	networkRepository      repository.NetworkRepositoryCaller
	userRepository         repository.UserRepositoryCaller
	userGroupRepository    repository.UserGroupRepositoryCaller
	jobRepository          repository.JobRepositoryCaller
	roleRepository         repository.RoleRepositoryCaller
	auth0Client            *client.Auth0Client
}

func NewNetworkUsecase(
	db database.SQLExecuter,
) NetworkUsecaseCaller {
	return &NetworkUsecase{
		adminNetworkRepository: repository.NewAdminNetworkRepository(db),
		networkRepository:      repository.NewNetworkRepository(db),
		userGroupRepository:    repository.NewUserGroupRepository(db),
		userRepository:         repository.NewUserRepository(db),
		jobRepository:          repository.NewJobRepository(db),
		roleRepository:         repository.NewRoleRepository(db),
	}
}

func (u *NetworkUsecase) GetNetwork(sub, organizationID string) (*organization.GetNetworkResponse, error) {
	i := strings.Index(sub, "|")
	providerID := sub[i+1:]
	_, err := u.userRepository.FindByProviderID(providerID)

	adminNetwork, err := u.adminNetworkRepository.FindByOrganizationID(organizationID)
	if err != nil {
		return nil, err
	}

	userGroup, err := u.userGroupRepository.FindByAdminNetworkID(adminNetwork.ID)
	if err != nil {
		return nil, err
	}

	network, err := u.networkRepository.FindByAdminNetworkID(adminNetwork.ID)
	if err != nil {
		return nil, err
	}

	job, err := u.networkRepository.FindByAdminNetworkID(adminNetwork.ID)
	if err != nil {
		return nil, err
	}

	role, err := u.roleRepository.FindByAdminNetworkID(adminNetwork.ID)
	if err != nil {
		return nil, err
	}

	return &organization.GetNetworkResponse{
		UserGroupID: uint64(userGroup.ID),
		OrgID:       organizationID,
		NetworkID:   uint64(network.ID),
		JobID:       uint64(job.ID),
		RoleID:      uint64(role.ID),
	}, nil
}
