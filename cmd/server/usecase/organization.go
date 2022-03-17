package usecase

import (
	"strings"

	client "github.com/Notch-Technologies/wizy/auth0"
	"github.com/Notch-Technologies/wizy/cmd/server/database"
	"github.com/Notch-Technologies/wizy/cmd/server/domain"
	"github.com/Notch-Technologies/wizy/cmd/server/repository"
)

type OrganizationUsecaseCaller interface {
	CreateOrganizationWithAuth0(name string) (*client.OrganizationResponse, error)
	CreateOrganization(name, organizationID string) (*domain.Organization, error)
	CreateUser(email, password, connection string) (*client.CreateAuth0User, error)
	AddMemberOnOrganization(userID, organizationID string) error
	EnableOrganizationConnection(organizationID string, isAssignMembershipOnLogin bool) error
	AssignAdminRole(userID string) error
}

type OrganizationUsecase struct {
	orgRepository *repository.OrgRepository
	auth0Client   *client.Auth0Client
}

func NewOrganizationUsecase(
	db database.SQLExecuter,
	client *client.Auth0Client,
) OrganizationUsecaseCaller {
	return &OrganizationUsecase{
		orgRepository: repository.NewOrgRepository(db),
		auth0Client:   client,
	}
}

func (o *OrganizationUsecase) CreateOrganizationWithAuth0(name string) (*client.OrganizationResponse, error) {
	token, err := o.auth0Client.GetAuth0ManagementAccessToken()
	if err != nil {
		return nil, err
	}

	lname := strings.ToLower(name)
	organization, err := o.auth0Client.CreateOrganization(lname, token)
	if err != nil {
		return nil, err
	}

	return organization, nil
}

func (o *OrganizationUsecase) CreateOrganization(name, organizationID string) (*domain.Organization, error) {
	org := domain.NewOrganization(name, organizationID)
	err := o.orgRepository.CreateOrganization(org)
	if err != nil {
		return nil, err
	}

	return org, nil
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
