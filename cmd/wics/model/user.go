package model

import (
	"encoding/json"
	"time"

	"github.com/Notch-Technologies/wizy/types/key"
)

type User struct {
	// create unique id by wissy
	ID string `json:"id"`
	// openid providerid
	ProviderID string `json:"provider_id"`
	// openid provider
	Provider string `json:"provider"`
	// store network id
	NetworkID string `json:"network_id"`
	// store group id
	OrgGroupID    string             `json:"org_group_id"`
	Permission key.PermissionType `json:"permission"`
	CreatedAt  time.Time          `json:"created_at"`
	UpdatedAt  time.Time          `json:"updated_at"`
}

func NewUser(id string, providerID string, provider string,
	networkID string, orgGroupID string, permission key.PermissionType) *User {
	return &User{
		ID:         id,
		ProviderID: providerID,
		Provider:   provider,
		NetworkID:  networkID,
		OrgGroupID: orgGroupID,
		Permission: permission,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}

func (u *User) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}
