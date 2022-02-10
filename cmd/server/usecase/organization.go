package usecase

import (
	"github.com/Notch-Technologies/wizy/client"
	"github.com/Notch-Technologies/wizy/cmd/server/database"
	"github.com/Notch-Technologies/wizy/cmd/server/domain"
	"github.com/Notch-Technologies/wizy/cmd/server/repository"
)

type OrganizationUscaseManager interface {
	CreateOrganizationWithAuth0(name, displayName string) (*client.OrganizationResponse, error)
	CreateOrganization(name, displayName, organizationID, logoURL string) (*domain.Organization, error)
	CreateUser(email, password, connection string) (*client.CreateAuth0User, error)
	AddMemberOnOrganization(userID, organizationID string) error
	EnableOrganizationConnection(organizationID string, isAssignMembershipOnLogin bool) error
	AssignUserAdminRole(userID string) error
}

type OrganizationUsecase struct {
	orgRepository       *repository.OrgRepository
	auth0Client *client.Auth0Client
}

func NewOrganizationUsecase(
	db database.SQLExecuter,
	client *client.Auth0Client,
) *OrganizationUsecase{
	return &OrganizationUsecase{
		orgRepository: repository.NewOrgRepository(db),
		auth0Client: client,
	}
}

func (o *OrganizationUsecase) CreateOrganizationWithAuth0(name, displayName string) (*client.OrganizationResponse, error) {
	token, err := o.auth0Client.GetAuth0ManagementAccessToken()
	if err != nil {
		return nil, err
	}

	organization, err := o.auth0Client.CreateOrganization(name, displayName, token)
	if err != nil {
		return nil, err
	}

	return organization, nil
}

func (o *OrganizationUsecase) CreateOrganization(name, displayName, organizationID string) (*domain.Organization, error) {
	orgGroup := domain.NewOrganization(name, displayName, organizationID)
	err := o.orgRepository.CreateOrganization(orgGroup)
	if err != nil {
		return nil, err
	}

	return orgGroup, nil
}

func (o *OrganizationUsecase) CreateUser(email, password, connection string) (*client.CreateAuth0User, error) {
	token, err := o.auth0Client.GetAuth0ManagementAccessToken()
	if err != nil {
		return nil, err
	}

	user, err := o.auth0Client.CreateUser(email, password, connection, token)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (o *OrganizationUsecase) AddMemberOnOrganization(userID, organizationID string) error {
	token, err := o.auth0Client.GetAuth0ManagementAccessToken()
	if err != nil {
		return err
	}

	err = o.auth0Client.AddMemberOnOrganization(userID, organizationID, token)
	return err
}

func (o *OrganizationUsecase) EnableOrganizationConnection(organizationID string, isAssignMembershipOnLogin bool) error {
	token, err := o.auth0Client.GetAuth0ManagementAccessToken()
	if err != nil {
		return err
	}

	err = o.auth0Client.EnableOraganizationConnection(token, organizationID, isAssignMembershipOnLogin)
	return err
}

func (o *OrganizationUsecase) AssignAdminRole(userID string) error {
	token, err := o.auth0Client.GetAuth0ManagementAccessToken()
	if err != nil {
		return err
	}

	err = o.auth0Client.AssignUserAdminRole(token, userID)
	return err
}
