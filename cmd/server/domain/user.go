package domain

import (
	"encoding/json"
	"time"
)

type User struct {
	ID             uint      `db:"id"`
	ProviderID     string    `db:"provider_id"`
	AdminNetworkID uint      `db:"admin_network_id"`
	NetworkID      uint      `db:"network_id"`
	UserGroupID    uint      `db:"user_group_id"`
	RoleID         uint      `db:"role_id"`
	Provider       string    `db:"provider"`
	Email          string    `db:"email"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}

func NewUser(
	providerID string, provider string, email string,
	networkID, userGroupID, adminNetworkID, roleID uint,
) *User {
	return &User{
		ProviderID:     providerID,
		AdminNetworkID: adminNetworkID,
		NetworkID:      networkID,
		UserGroupID:    userGroupID,
		RoleID:         roleID,
		Provider:       provider,
		Email:          email,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
}

func (u User) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}
