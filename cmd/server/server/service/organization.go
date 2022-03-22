package service

import (
	"context"

	"github.com/Notch-Technologies/wizy/cmd/server/database"
	"github.com/Notch-Technologies/wizy/cmd/server/pb/organization"
	"github.com/Notch-Technologies/wizy/cmd/server/usecase"
)

type OrganizationServerServiceCaller interface {
	GetNetwork(ctx context.Context, req *organization.GetNetworkRequest) (*organization.GetNetworkResponse, error)
}

type OrganizationServerService struct {
	db          *database.Sqlite
}

func NewOrganizationServerService(
	db *database.Sqlite,
) OrganizationServerServiceCaller {
	return &OrganizationServerService{
		db:          db,
	}
}

func (s *OrganizationServerService) GetNetwork(ctx context.Context, req *organization.GetNetworkRequest) (*organization.GetNetworkResponse, error) {
	sub := getSub(ctx)
	orgID := req.GetOrgID()
	networkUsecase := usecase.NewNetworkUsecase(s.db)
	return networkUsecase.GetNetwork(sub, orgID)
}
