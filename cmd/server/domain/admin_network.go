package domain

import "time"

type AdminNetwork struct {
	ID        uint      `db:"id"`
	Name      string    `db:"name"`
	OrgID     string    `db:"org_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func NewAdminNetwork(name, orgID string) *AdminNetwork {
	return &AdminNetwork{
		Name:      name,
		OrgID:     orgID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
