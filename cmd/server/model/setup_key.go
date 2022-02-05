package model

import (
	"time"

	"github.com/Notch-Technologies/wizy/types/key"
)

type SetupKey struct {
	ID         string           `db:"id"`
	UserID     string           `db:"user_id"`
	Key    	   string	 		`db:"key"`
	KeyType    key.SetupKeyType `db:"key_type"`
	Revoked    bool             `db:"revoked"`
	CreatedAt  time.Time        `db:"created_at"`
	UpdatedAt  time.Time        `db:"updated_at"`
}

func NewSetupKey(id, userID, key string, keytype key.SetupKeyType, revoked bool) *SetupKey {
	return &SetupKey{
		ID:         id,
		Key:        key,
		UserID:     userID,
		KeyType:    keytype,
		Revoked:    revoked,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}
