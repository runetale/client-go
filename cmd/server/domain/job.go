package domain

import (
	"time"
)

type Job struct {
	ID             uint      `db:"id"`
	AdminNetworkID uint      `db:"admin_network_id"`
	Name           string    `db:"name"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}

func NewJob(adminNetworkID uint, name string) *Job {
	return &Job{
		Name:           name,
		AdminNetworkID: adminNetworkID,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
}
