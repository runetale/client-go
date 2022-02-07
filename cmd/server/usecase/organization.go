package usecase

import (
	"github.com/Notch-Technologies/wizy/cmd/server/database"
	"github.com/Notch-Technologies/wizy/cmd/server/domain"
	"github.com/Notch-Technologies/wizy/cmd/server/repository"
)

type OrganizationUscaseManager interface {
	CreateOrganization(name, displayName, organizationID string) (*domain.OrgGroup, error)
}

type OrganizationUsecase struct {
	orgRepository       *repository.OrgRepository
}

func NewOrganizationUsecase(
	db database.SQLExecuter,
) *OrganizationUsecase{
	return &OrganizationUsecase{
		orgRepository: repository.NewOrgRepository(db),
	}
}

func (o *OrganizationUsecase) CreateOrganization(name, displayName, organizationID string) (*domain.OrgGroup, error) {
	orgGroup := domain.NewOrgGroup(name, displayName, organizationID)
	err := o.orgRepository.CreateOrganization(orgGroup)
	if err != nil {
		return nil, err
	}

	return orgGroup, nil
}
