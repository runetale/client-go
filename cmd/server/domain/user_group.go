package domain

import (
	"time"
)

type UserGroup struct {
	ID             uint      `db:"id"`
	AdminNetworkID uint      `db:"admin_network_id"`
	Name           string    `db:"name"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}

func NewUserGroup(
	adminNetworkID uint, name string,
) *UserGroup {
	return &UserGroup{
		AdminNetworkID: adminNetworkID,
		Name:           name,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
}
