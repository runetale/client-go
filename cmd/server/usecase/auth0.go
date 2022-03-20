package usecase

import (
	"strings"

	client "github.com/Notch-Technologies/wizy/auth0"
)

type Auth0UsecaseCaller interface {
	CreateOrganizationWithAuth0(name string) (*client.OrganizationResponse, error)
	CreateUser(email, password, connection string) (*client.CreateAuth0User, error)
	AddMemberOnOrganization(userID, organizationID string) error
	EnableOrganizationConnection(organizationID string, isAssignMembershipOnLogin bool) error
	AssignAdminRole(userID string) error
}

type Auth0Usecase struct {
	auth0Client *client.Auth0Client
}

func NewAuth0Usecase(
	client *client.Auth0Client,
) Auth0UsecaseCaller {
	return &Auth0Usecase{
		auth0Client: client,
	}
}

func (u *Auth0Usecase) CreateOrganizationWithAuth0(name string) (*client.OrganizationResponse, error) {
	token, err := u.auth0Client.GetAuth0ManagementAccessToken()
	if err != nil {
		return nil, err
	}

	lname := strings.ToLower(name)
	organization, err := u.auth0Client.CreateOrganization(lname, token)
	if err != nil {
		return nil, err
	}

	return organization, nil
}

func (u *Auth0Usecase) CreateUser(email, password, connection string) (*client.CreateAuth0User, error) {
	token, err := u.auth0Client.GetAuth0ManagementAccessToken()
	if err != nil {
		return nil, err
	}

	user, err := u.auth0Client.CreateUser(email, password, connection, token)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *Auth0Usecase) AddMemberOnOrganization(userID, organizationID string) error {
	token, err := u.auth0Client.GetAuth0ManagementAccessToken()
	if err != nil {
		return err
	}

	err = u.auth0Client.AddMemberOnOrganization(userID, organizationID, token)
	return err
}

func (u *Auth0Usecase) EnableOrganizationConnection(organizationID string, isAssignMembershipOnLogin bool) error {
	token, err := u.auth0Client.GetAuth0ManagementAccessToken()
	if err != nil {
		return err
	}

	// TODO: Google, Microsoft, GithubのConnectionを作成し、有効化する
	err = u.auth0Client.EnableOraganizationConnection(token, organizationID, isAssignMembershipOnLogin)
	return err
}

func (u *Auth0Usecase) AssignAdminRole(userID string) error {
	token, err := u.auth0Client.GetAuth0ManagementAccessToken()
	if err != nil {
		return err
	}

	err = u.auth0Client.AssignUserAdminRole(token, userID)
	return err
}
