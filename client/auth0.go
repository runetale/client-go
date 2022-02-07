package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type Auth0Client struct {
	Domain string
	M2MClientID string
	M2MClientSecret string
	Audience string
}

func NewAuth0Client() *Auth0Client {
	domain := os.Getenv("AUTH0_DOMAIN")
	clientid := os.Getenv("AUTH0_M2M_CLIENT_ID")
	clientsecret := os.Getenv("AUTH0_M2M_CLIENT_SECRET")
	audience := os.Getenv("AUTH0_AUDIENCE")

	return &Auth0Client{
		Domain: domain,
		M2MClientID: clientid,
		M2MClientSecret: clientsecret,
		Audience: audience,
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
	ID 			  string `json:"id"`
	Name          string `json:"name"`
	DisplayName   string `json:"display_name"`
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

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var r OrganizationResponse

	if err := json.Unmarshal(body, &r); err != nil {
		return nil, err
	}

	return &r, nil
}
