package domain

import (
	"time"

	"github.com/Notch-Technologies/wizy/types/key"
)

type UserGroup struct {
	ID         uint               `db:"id"`
	Name       string             `db:"name"`
	Permission key.PermissionType `db:"permission"`
	CreatedAt  time.Time          `db:"created_at"`
	UpdatedAt  time.Time          `db:"updated_at"`
}

func NewUserGroup(name string, permission key.PermissionType) *UserGroup {
	return &UserGroup{
		Name:       name,
		Permission: permission,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
}
