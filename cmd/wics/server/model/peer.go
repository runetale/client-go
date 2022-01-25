package model

import "time"

type Peer struct {
	ID string
	SetupKey string
	ClientToken string
	NetworkID string
	UserGroupID string
	UserID  string
	Permission int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewPeer(id string, setupKey string, clientToken string,
	networkID string, userGroupID string, userID string, 
	permission int64) *Peer {
	return &Peer{
		ID: id,
		SetupKey: setupKey,
		ClientToken: clientToken,
		NetworkID: networkID,
		UserGroupID: userGroupID,
		UserID: userID,
		Permission: permission,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
