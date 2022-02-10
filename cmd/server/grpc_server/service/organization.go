package service

import (
	"context"

	"github.com/Notch-Technologies/wizy/client"
	"github.com/Notch-Technologies/wizy/cmd/server/database"
	"github.com/Notch-Technologies/wizy/cmd/server/pb/organization"
	"github.com/Notch-Technologies/wizy/cmd/server/usecase"
)

type OrganizationServiceServer struct {
	db          *database.Sqlite
	auth0Client *client.Auth0Client

	organization.UnimplementedOrganizationServiceServer
}

func NewOrganizationServiceServer(db *database.Sqlite, client *client.Auth0Client) *OrganizationServiceServer {
	return &OrganizationServiceServer{
		auth0Client: client,
		db:          db,
	}
}

func (oss *OrganizationServiceServer) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	return ctx, nil
}

func (oss *OrganizationServiceServer) Create(ctx context.Context, req *organization.OrganizationCreateRequest) (*organization.OrganizationCreateResponse, error) {
	organizationUsecase := usecase.NewOrganizationUsecase(oss.db, oss.auth0Client)

	org, err := organizationUsecase.CreateOrganizationWithAuth0(req.GetName(), req.GetDisplayName())
	if err != nil {
		return nil, err
	}

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

func (oss *OrganizationServiceServer) CreateAdminUser(ctx context.Context, req *organization.OrganizationCreateAdminUserRequest) (*organization.OrganizationCreateAdminUserResponse, error) {
	organizationUsecase := usecase.NewOrganizationUsecase(oss.db, oss.auth0Client)
	user, err := organizationUsecase.CreateUser(req.GetEmail(), req.GetPassword(), oss.auth0Client.DatabaseConnectionName)
	if err != nil {
		return nil, err
	}

	err = organizationUsecase.AddMemberOnOrganization(user.UserID, req.GetOrganizationID())
	if err != nil {
		return nil, err
	}

	return &organization.OrganizationCreateAdminUserResponse{
		OrganizationID: user.UserID,
	}, nil
}
