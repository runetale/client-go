package service

import (
	"context"

	"github.com/Notch-Technologies/wizy/cmd/server/database"
	"github.com/Notch-Technologies/wizy/cmd/server/pb/organization"
	//"github.com/Notch-Technologies/wizy/cmd/server/usecase"
)

type OrganizationServiceServer struct {
	db *database.Sqlite

	organization.UnimplementedOrganizationServiceServer
}

func NewOrganizationServiceServer(db *database.Sqlite) *OrganizationServiceServer {
	return &OrganizationServiceServer{
		db: db,
	}
}

func (oss *OrganizationServiceServer) Create(ctx context.Context, req *organization.OrganizationCreateRequest) (*organization.OrganizationCreateResponse, error) {
	//organizationUsecase := usecase.NewOrganizationUsecase(oss.db)


	return &organization.OrganizationCreateResponse{}, nil
}
