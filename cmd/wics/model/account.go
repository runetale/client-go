package model

import "time"

type Account struct {
	ID         string
	Networks   []*Network
	OrgGroups []*OrgGroup
	Peers      []*Peer
	Users      []*User
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func NewAccount(id string) *Account {
	return &Account{
		Networks:   []*Network{},
		OrgGroups: []*OrgGroup{},
		Peers:      []*Peer{},
		Users:      []*User{},
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}
