package service

import (
	"context"
	"fmt"

	client "github.com/Notch-Technologies/wizy/auth0"
	"github.com/Notch-Technologies/wizy/cmd/server/database"
	"github.com/Notch-Technologies/wizy/cmd/server/pb/organization"
	"github.com/Notch-Technologies/wizy/cmd/server/usecase"
)

type OrganizationServerServiceCaller interface {
	Create(ctx context.Context, req *organization.OrganizationCreateRequest) (*organization.OrganizationCreateResponse, error)
	CreateAdminUser(ctx context.Context, req *organization.OrganizationCreateAdminUserRequest) (*organization.OrganizationCreateAdminUserResponse, error)
}

type OrganizationServerService struct {
	db          *database.Sqlite
	auth0Client *client.Auth0Client
}

func NewOrganizationServerService(db *database.Sqlite, client *client.Auth0Client) *OrganizationServerService {
	return &OrganizationServerService{
		auth0Client: client,
		db:          db,
	}
}

func (oss *OrganizationServerService) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	return ctx, nil
}

func (oss *OrganizationServerService) Create(ctx context.Context, req *organization.OrganizationCreateRequest) (*organization.OrganizationCreateResponse, error) {
	organizationUsecase := usecase.NewOrganizationUsecase(oss.db, oss.auth0Client)

	org, err := organizationUsecase.CreateOrganizationWithAuth0(req.GetName(), req.GetDisplayName())
	if err != nil {
		return nil, err
	}

	fmt.Println(org.ID)
	organizationGroup, err := organizationUsecase.CreateOrganization(req.GetName(), req.GetDisplayName(), org.ID)
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
