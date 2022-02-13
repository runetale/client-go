package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type Auth0Client struct {
	Domain                 string
	M2MClientID            string
	M2MClientSecret        string
	Audience               string
	DatabaseConnectionID   string
	DatabaseConnectionName string
	AdminRoleID            string
}

func NewAuth0Client() *Auth0Client {
	domain := os.Getenv("AUTH0_DOMAIN")
	clientid := os.Getenv("AUTH0_M2M_CLIENT_ID")
	clientsecret := os.Getenv("AUTH0_M2M_CLIENT_SECRET")
	audience := os.Getenv("AUTH0_AUDIENCE")
	databaseConnectionID := os.Getenv("AUTH0_DATABASE_CONNCTION_ID")
	databaseConnectionName := os.Getenv("AUTH0_DATABASE_CONNCTION_NAME")
	adminRoleID := os.Getenv("AUTH0_ADMIN_ROLE_ID")

	return &Auth0Client{
		Domain:                 domain,
		M2MClientID:            clientid,
		M2MClientSecret:        clientsecret,
		Audience:               audience,
		DatabaseConnectionID:   databaseConnectionID,
		DatabaseConnectionName: databaseConnectionName,
		AdminRoleID:            adminRoleID,
	}
}

func (a *Auth0Client) GetAuth0ManagementAccessToken() (string, error) {
	url := fmt.Sprintf("https://%s/oauth/token", a.Domain)

	payload := strings.NewReader(
		fmt.Sprintf("grant_type=client_credentials&client_id=%s&client_secret=%s&audience=%s",
			a.M2MClientID, a.M2MClientSecret, a.Audience,
		))

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	res, _ := http.DefaultClient.Do(req)

	body, _ := ioutil.ReadAll(res.Body)

	type managementResponse struct {
		AccessToken string `json:"access_token"`
	}
	var m managementResponse

	if err := json.Unmarshal(body, &m); err != nil {
		return "", err
	}

	defer res.Body.Close()

	return m.AccessToken, nil
}

func (a *Auth0Client) IsAdmin(sub, token string) (bool, error) {
	roles, err := a.GetUserRoles(sub, token)
	if err != nil {
		return false, err
	}

	for _, r := range *roles {
		if r.Name == "admin" {
			return true, nil
		}
	}
	return false, nil
}

type roleResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (a *Auth0Client) GetUserRoles(sub, token string) (*[]roleResponse, error) {
	url := fmt.Sprintf("%susers/%s/roles", a.Audience, sub)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("authorization", fmt.Sprintf("Bearer %s", token))

	res, _ := http.DefaultClient.Do(req)

	body, _ := ioutil.ReadAll(res.Body)

	var r []roleResponse

	if err := json.Unmarshal(body, &r); err != nil {
		return nil, err
	}

	defer res.Body.Close()

	return &r, nil
}

type OrganizationResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
}

func (a *Auth0Client) CreateOrganization(name, displayName string, token string) (*OrganizationResponse, error) {
	url := fmt.Sprintf("https://%s/api/v2/organizations", a.Domain)

	b := fmt.Sprintf(`{"name": "%s", "display_name": "%s"}`,
		name, displayName)

	payload := strings.NewReader(b)

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return nil, err
	}

	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Add("cache-control", "no-cache")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var r OrganizationResponse

	if err := json.Unmarshal(body, &r); err != nil {
		return nil, err
	}

	defer res.Body.Close()

	return &r, nil
}

type CreateAuth0User struct {
	UserID        string        `json:"user_id"`
	Email         string        `json:"email"`
	EmailVerified bool          `json:"email_verified"`
	Name          string        `json:"name"`
	NickName      string        `json:"nick_name"`
	Picture       string        `json:"picture"`
	Identities    []*Identities `json:"identities"`
	CreatedAt     string        `json:"created_at"`
	UpdatedAt     string        `json:"updated_at"`
}

type Identities struct {
	Connection string `json:"connection"`
	UserID     string `json:"user_id"`
	Provider   string `json:"provider"`
	IsSocial   bool   `json:"isSocial"`
}

func (a *Auth0Client) CreateUser(email, password, connection, token string) (*CreateAuth0User, error) {
	url := fmt.Sprintf("https://%s/api/v2/users", a.Domain)

	b := fmt.Sprintf(`{"email": "%s", "password": "%s", "connection": "%s"}`,
		email, password, connection)

	payload := strings.NewReader(b)

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return nil, err
	}

	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Add("cache-control", "no-cache")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var r CreateAuth0User

	if err := json.Unmarshal(body, &r); err != nil {
		return nil, err
	}

	defer res.Body.Close()

	return &r, nil
}

func (a *Auth0Client) AddMemberOnOrganization(userID, organizationID, token string) error {
	url := fmt.Sprintf("https://%s/api/v2/organizations/%s/members", a.Domain, organizationID)

	type param struct {
		Members []string `json:"members"`
	}

	p := param{Members: []string{userID}}
	payload, err := json.Marshal(p)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}

	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Add("cache-control", "no-cache")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	return nil
}

func (a *Auth0Client) EnableOraganizationConnection(token, organizationID string, isAssignMembershipOnLogin bool) error {
	url := fmt.Sprintf("https://%s/api/v2/organizations/%s/enabled_connections", a.Domain, organizationID)

	b := fmt.Sprintf(`{"connection_id": "%s", "assign_membership_on_login": %t}`,
		a.DatabaseConnectionID, isAssignMembershipOnLogin)

	payload := strings.NewReader(b)

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return err
	}

	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Add("cache-control", "no-cache")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	return nil
}

func (a *Auth0Client) AssignUserAdminRole(token, userID string) error {
	url := fmt.Sprintf("https://%s/api/v2/roles/%s/users", a.Domain, a.AdminRoleID)

	type param struct {
		Users []string `json:"users"`
	}

	p := param{Users: []string{userID}}
	payload, err := json.Marshal(p)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}

	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Add("cache-control", "no-cache")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	return nil

}
