package usecase

import (
	"strings"

	client "github.com/Notch-Technologies/wizy/auth0"
	"github.com/Notch-Technologies/wizy/cmd/server/database"
	"github.com/Notch-Technologies/wizy/cmd/server/domain"
	"github.com/Notch-Technologies/wizy/cmd/server/repository"
	"github.com/Notch-Technologies/wizy/types/key"
)

type AdminNetworkUsecaseCaller interface {
	CreateAdminNetworkWithDefault(userID, name, email, organizationID string) (*domain.AdminNetwork, error)
}

type AdminNetworkUsecase struct {
	adminNetworkRepository repository.AdminNetworkRepositoryCaller
	networkRepository      repository.NetworkRepositoryCaller
	userRepository 	   	   repository.UserRepositoryCaller
	userGroupRepository    repository.UserGroupRepositoryCaller
	jobRepository          repository.JobRepositoryCaller
	roleRepository         repository.RoleRepositoryCaller
	auth0Client            *client.Auth0Client
}

func NewAdminNetworkUsecase(
	db database.SQLExecuter,
	client *client.Auth0Client,
) AdminNetworkUsecaseCaller {
	return &AdminNetworkUsecase{
		adminNetworkRepository: repository.NewAdminNetworkRepository(db),
		networkRepository:      repository.NewNetworkRepository(db),
		userGroupRepository:    repository.NewUserGroupRepository(db),
		userRepository: 	    repository.NewUserRepository(db),
		jobRepository:          repository.NewJobRepository(db),
		roleRepository:         repository.NewRoleRepository(db),
		auth0Client:            client,
	}
}

// 1. create default admin network
// 2. create default network
// 3. create user group (default)
// 4. create job group (default)
// 5. role group (default)
// 6. create user
//
func (u *AdminNetworkUsecase) CreateAdminNetworkWithDefault(
	userID, name, email, organizationID string,
) (*domain.AdminNetwork, error) {
	const (
		dName = "default"
		ip    = "100.64.0.0"
		cidr  = 10
	)

	// google-oauth2|hogehoge
	// we'll make it look like this.
	// google-oauth2 and hogehoge
	i := strings.Index(userID, "|")
	provider := userID[:i]
	providerID := userID[i+1:]

	// 1
	//
	adminNetwork := domain.NewAdminNetwork(name, organizationID)
	err := u.adminNetworkRepository.CreateAdminNetwork(adminNetwork)
	if err != nil {
		return nil, err
	}

	// 2
	//
	network := domain.NewNetwork(adminNetwork.ID, dName, ip, cidr)
	err = u.networkRepository.CreateNetwork(network)
	if err != nil {
		return nil, err
	}

	// 3
	//
	userGroup := domain.NewUserGroup(adminNetwork.ID, dName)
	err = u.userGroupRepository.CreateUserGroup(userGroup)
	if err != nil {
		return nil, err
	}

	// 4
	//
	job := domain.NewJob(adminNetwork.ID, dName)
	err = u.jobRepository.CreateJob(job)
	if err != nil {
		return nil, err
	}

	// 5
	//
	role := domain.NewRole(adminNetwork.ID, dName, key.RWXKey)
	err = u.roleRepository.CreateRole(role)
	if err != nil {
		return nil, err
	}


	// 6
	//
	user := domain.NewUser(
		providerID,
		provider,
		email,
		network.ID,
		userGroup.ID,
		adminNetwork.ID,
		role.ID,
	)
	err = u.userRepository.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return adminNetwork, nil
}
