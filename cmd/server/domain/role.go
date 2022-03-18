package domain

import (
	"time"

	"github.com/Notch-Technologies/wizy/types/key"
)

type Role struct {
	ID             uint               `db:"id"`
	AdminNetworkID uint               `db:"admin_network_id"`
	Name           string             `db:"name"`
	Permission     key.PermissionType `db:"permission"`
	CreatedAt      time.Time          `db:"created_at"`
	UpdatedAt      time.Time          `db:"updated_at"`
}

func NewRole(
	adminNetworkID uint, name string,
	permission key.PermissionType,
) *Job {
	return &Job{
		Name:           name,
		AdminNetworkID: adminNetworkID,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
}
