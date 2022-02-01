package model

import (
	"time"
)

type UserGroup struct {
	ID        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewGroup(id, name string) *UserGroup {
	return &UserGroup{
		ID:        id,
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
