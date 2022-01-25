package model

import "time"

type SetupKey struct {
	ID         string
	KeyType    string
	Revoked    bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
	LastUsedAt time.Time
}

func NewSetupKey(id, keytype string, revoked bool) *SetupKey {
	return &SetupKey{
		ID:         id,
		KeyType:    keytype,
		Revoked:    revoked,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		LastUsedAt: time.Now(),
	}
}
