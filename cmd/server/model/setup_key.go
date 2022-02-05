package model

import (
	"time"

	"github.com/Notch-Technologies/wizy/types/key"
)

type SetupKey struct {
	ID         string           `json:"id"`
	Key        string           `json:"key"`
	UserID     string           `json:"user_id"`
	KeyType    key.SetupKeyType `json:"key_type"`
	Revoked    bool             `json:"revoked"`
	CreatedAt  time.Time        `json:"created_at"`
	LastusedAt time.Time        `json:"lastused_at"`
	UpdatedAt  time.Time        `json:"updated_at"`
}

func NewSetupKey(id, userID, key string, keytype key.SetupKeyType, revoked bool) *SetupKey {
	return &SetupKey{
		ID:         id,
		Key:        key,
		UserID:     userID,
		KeyType:    keytype,
		Revoked:    revoked,
		CreatedAt:  time.Now(),
		LastusedAt: time.Now(),
		UpdatedAt:  time.Now(),
	}
}
