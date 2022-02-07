package domain

import (
	"time"
)

type OrgGroup struct {
	ID          uint      `db:"id"`
	Name        string    `db:"name"`
	DisplayName string    `db:"display_name"`
	OrgID       string    `db:"org_id"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

func NewOrgGroup(name, displayName, orgID string) *OrgGroup {
	return &OrgGroup{
		Name:        name,
		DisplayName: displayName,
		OrgID:       orgID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}
