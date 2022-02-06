package domain

import (
	"time"
)

type Job struct {
	ID        uint      `db:"id"`
	Name      string    `db:"name"`
	UserID    uint      `db:"user_id"`
	OrgID     uint      `db:"org_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func NewJob(name string, userID, orgID uint) *Job {
	return &Job{
		Name:      name,
		UserID:    userID,
		OrgID:     orgID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
