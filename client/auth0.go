package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func GetAuth0ManagementAccessToken() (string, error) {
	domain := os.Getenv("AUTH0_DOMAIN")
	url := fmt.Sprintf("https://%s/oauth/token", domain)

	clientid := os.Getenv("AUTH0_M2M_CLIENT_ID")
	clientsecret := os.Getenv("AUTH0_M2M_CLIENT_SECRET")
	audience := os.Getenv("AUTH0_AUDIENCE")

	payload := strings.NewReader(fmt.Sprintf("grant_type=client_credentials&client_id=%s&client_secret=%s&audience=%s", clientid, clientsecret, audience))

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

func IsAdmin(sub, token string) (bool, error) {
	roles, err := GetUserRoles(sub, token)
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

func GetUserRoles(sub, token string) (*[]roleResponse, error) {
	audience := os.Getenv("AUTH0_AUDIENCE")
	url := fmt.Sprintf("%susers/%s/roles", audience, sub)

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
