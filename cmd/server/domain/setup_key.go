package domain

import (
	"time"

	"github.com/Notch-Technologies/wizy/types/key"
)

type SetupKey struct {
	ID        uint             `db:"id"`
	UserID    uint             `db:"user_id"`
	Key       string           `db:"key"`
	KeyType   key.SetupKeyType `db:"key_type"`
	Revoked   bool             `db:"revoked"`
	CreatedAt time.Time        `db:"created_at"`
	UpdatedAt time.Time        `db:"updated_at"`
}

func NewSetupKey(userID uint, key string, keytype key.SetupKeyType, revoked bool) *SetupKey {
	return &SetupKey{
		Key:       key,
		UserID:    userID,
		KeyType:   keytype,
		Revoked:   revoked,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
