package domain

import (
	"time"

	"github.com/Notch-Technologies/wizy/types/key"
)

type SetupKey struct {
	ID             uint             `db:"id"`
	AdminNetworkID uint             `db:"admin_network_id"`
	UserID         uint             `db:"user_id"`
	Key            string           `db:"key"`
	KeyType        key.SetupKeyType `db:"key_type"`
	CreatedAt      time.Time        `db:"created_at"`
	UpdatedAt      time.Time        `db:"updated_at"`
}

func NewSetupKey(
	adminNetworkID, userID uint,
	key string, keytype key.SetupKeyType,
) *SetupKey {
	return &SetupKey{
		AdminNetworkID: adminNetworkID,
		UserID:         userID,
		Key:            key,
		KeyType:        keytype,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
}
