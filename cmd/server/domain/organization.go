package domain

import (
	"time"
)

type Organization struct {
	ID          uint      `db:"id"`
	Name        string    `db:"name"`
	OrgID       string    `db:"org_id"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

func NewOrganization(name, orgID string) *Organization {
	return &Organization{
		Name:        name,
		OrgID:       orgID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}
