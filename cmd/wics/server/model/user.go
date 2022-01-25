package model

import "time"

type User struct {
	ID string
	ProviderID string
	Provider string
	NetworkID string
	UserGroupID string
	UserID  string
	Permission int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(id string, providerID string, provider string,
	networkID string, userGroupID string, userID string, 
	permission int64) *User {
	return &User{
		ID: id,
		ProviderID: providerID,
		Provider: provider,
		NetworkID: networkID,
		UserGroupID: userGroupID,
		UserID: userID,
		Permission: permission,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
