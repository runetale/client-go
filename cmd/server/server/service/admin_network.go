package service

import (
	"context"

	client "github.com/Notch-Technologies/wizy/auth0"
	"github.com/Notch-Technologies/wizy/cmd/server/database"
	"github.com/Notch-Technologies/wizy/cmd/server/pb/admin_network"
	"github.com/Notch-Technologies/wizy/cmd/server/usecase"
)

type AdminNetworkServerServiceCaller interface {
	CreateDefaultNetwork(ctx context.Context, req *admin_network.CreateDefaultAdminNetworkRequest) (*admin_network.CreateDefaultAdminNetworkResponse, error)
}

type AdminNetworkServerService struct {
	db          *database.Sqlite
	auth0Client *client.Auth0Client
}

func NewAdminNetworkServerService(
	db *database.Sqlite, client *client.Auth0Client,
) AdminNetworkServerServiceCaller {
	return &AdminNetworkServerService{
		db:          db,
		auth0Client: client,
	}
}

func (s *AdminNetworkServerService) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	return ctx, nil
}

func (s *AdminNetworkServerService) CreateDefaultNetwork(ctx context.Context, req *admin_network.CreateDefaultAdminNetworkRequest) (*admin_network.CreateDefaultAdminNetworkResponse, error) {

	companyName := req.GetCompanyName()
	userID := req.GetUserID()
	email := req.GetEmail()

	adminNetworkUsecase := usecase.NewAdminNetworkUsecase(s.db, s.auth0Client)
	auth0Usecase := usecase.NewAuth0Usecase(s.auth0Client)

	// 1
	//
	auth0Org, err := auth0Usecase.CreateOrganizationWithAuth0(companyName)
	if err != nil {
		return nil, err
	}

	// 2
	//
	adminNetwork, err := adminNetworkUsecase.CreateAdminNetworkWithDefault(userID, companyName, email, auth0Org.ID)
	if err != nil {
		return nil, err
	}

	// 3
	//
	err = auth0Usecase.EnableOrganizationConnection(adminNetwork.OrgID, true)
	if err != nil {
		return nil, err
	}

	// 4
	//
	err = auth0Usecase.AddMemberOnOrganization(userID, adminNetwork.OrgID)
	if err != nil {
		return nil, err
	}

	// 5
	//
	err = auth0Usecase.AssignAdminRole(userID)
	if err != nil {
		return nil, err
	}

	return &admin_network.CreateDefaultAdminNetworkResponse{
		OrganizationID: adminNetwork.OrgID,
	}, nil
}
