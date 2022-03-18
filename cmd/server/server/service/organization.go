package service

import (
	"context"

	client "github.com/Notch-Technologies/wizy/auth0"
	"github.com/Notch-Technologies/wizy/cmd/server/database"
	"github.com/Notch-Technologies/wizy/cmd/server/pb/organization"
	"github.com/Notch-Technologies/wizy/cmd/server/usecase"
)

type OrganizationServerServiceCaller interface {
	Create(ctx context.Context, req *organization.OrganizationCreateRequest) (*organization.OrganizationCreateResponse, error)
	CreateAdminUser(ctx context.Context, req *organization.OrganizationCreateAdminUserRequest) (*organization.OrganizationCreateAdminUserResponse, error)
	CreateNetwork(ctx context.Context, req *organization.OrganizationCreateNetworkRequest) (*organization.OrganizationCreateNetworkResponse, error)
}

type OrganizationServerService struct {
	db          *database.Sqlite
	auth0Client *client.Auth0Client
}

func NewOrganizationServerService(
	db *database.Sqlite, client *client.Auth0Client,
) OrganizationServerServiceCaller {
	return &OrganizationServerService{
		auth0Client: client,
		db:          db,
	}
}

func (oss *OrganizationServerService) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	return ctx, nil
}

// TOOD: remove
func (oss *OrganizationServerService) Create(ctx context.Context, req *organization.OrganizationCreateRequest) (*organization.OrganizationCreateResponse, error) {
	organizationUsecase := usecase.NewOrganizationUsecase(oss.db, oss.auth0Client)

	org, err := organizationUsecase.CreateOrganizationWithAuth0(req.GetName())
	if err != nil {
		return nil, err
	}

	organizationGroup, err := organizationUsecase.CreateOrganization(req.GetName(), org.ID)
	if err != nil {
		return nil, err
	}

	err = organizationUsecase.EnableOrganizationConnection(org.ID, true)
	if err != nil {
		return nil, err
	}

	return &organization.OrganizationCreateResponse{
		OrganizationID: organizationGroup.OrgID,
	}, nil
}

// TOOD: remove
func (oss *OrganizationServerService) CreateAdminUser(ctx context.Context, req *organization.OrganizationCreateAdminUserRequest) (*organization.OrganizationCreateAdminUserResponse, error) {
	organizationUsecase := usecase.NewOrganizationUsecase(oss.db, oss.auth0Client)
	user, err := organizationUsecase.CreateUser(req.GetEmail(), req.GetPassword(), oss.auth0Client.DatabaseConnectionName)
	if err != nil {
		return nil, err
	}

	err = organizationUsecase.AddMemberOnOrganization(user.UserID, req.GetOrganizationID())
	if err != nil {
		return nil, err
	}

	err = organizationUsecase.AssignAdminRole(user.UserID)
	if err != nil {
		return nil, err
	}

	return &organization.OrganizationCreateAdminUserResponse{
		OrganizationID: user.UserID,
	}, nil
}

// 1. create organization account with auth0
// 2. register organization information in database
// 3. allow organization login methods
// 4. register a user as a member of organization
// 5. assign admin role
// 6. create company network (default)
// 7. create user group (default)
//
func (oss *OrganizationServerService) CreateNetwork(
	ctx context.Context, req *organization.OrganizationCreateNetworkRequest,
) (*organization.OrganizationCreateNetworkResponse, error) {
	companyName := req.GetCompanyName()
	userID := req.GetUserID()

	organizationUsecase := usecase.NewOrganizationUsecase(oss.db, oss.auth0Client)
	
	// 1
	//
	auth0Org, err := organizationUsecase.CreateOrganizationWithAuth0(companyName)
	if err != nil {
		return nil, err
	}

	// 2
	//
	org, err := organizationUsecase.CreateOrganization(companyName, auth0Org.ID)
	if err != nil {
		return nil, err
	}

	// 3
	//
	err = organizationUsecase.EnableOrganizationConnection(org.OrgID, true)
	if err != nil {
		return nil, err
	}

	// 4
	//
	err = organizationUsecase.AddMemberOnOrganization(userID, org.OrgID)
	if err != nil {
		return nil, err
	}

	// 5
	//
	err = organizationUsecase.AssignAdminRole(userID)
	if err != nil {
		return nil, err
	}

	return &organization.OrganizationCreateNetworkResponse{
		OrganizationID: userID,
	}, nil
}
