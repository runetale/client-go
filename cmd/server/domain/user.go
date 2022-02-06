package domain

import (
	"encoding/json"
	"time"

	"github.com/Notch-Technologies/wizy/types/key"
)

type User struct {
	ID          uint               `db:"id"`
	ProviderID  string             `db:"provider_id"`
	Provider    string             `db:"provider"`
	OrgGroupID  uint               `db:"org_id"`
	NetworkID   uint               `db:"network_id"`
	UserGroupID uint               `db:"user_group_id"`
	Permission  key.PermissionType `db:"permission"`
	CreatedAt   time.Time          `db:"created_at"`
	UpdatedAt   time.Time          `db:"updated_at"`
}

func NewUser(providerID string, provider string,
	networkID, userGroupID, orgGroupID uint, permission key.PermissionType) *User {
	return &User{
		ProviderID:  providerID,
		Provider:    provider,
		OrgGroupID:  orgGroupID,
		NetworkID:   networkID,
		UserGroupID: userGroupID,
		Permission:  permission,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func (u User) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}
