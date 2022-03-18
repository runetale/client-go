package repository

import (
	"github.com/Notch-Technologies/wizy/cmd/server/database"
	"github.com/Notch-Technologies/wizy/cmd/server/domain"
)

type JobRepositoryManager interface {
	CreateJob(job *domain.Job) error
}

type JobRepository struct {
	db database.SQLExecuter
}

func NewJobRepository(db database.SQLExecuter) *JobRepository {
	return &JobRepository{
		db: db,
	}
}

func (o *JobRepository) CreateJob(job *domain.Job) error {
	lastID, err := o.db.Exec(`
	INSERT INTO jobs (
		admin_network_id,
		name,
		created_at,
		updated_at
	) VALUES (?, ?, ?, ?)
	`,
		job.AdminNetworkID,
		job.Name,
		job.CreatedAt,
		job.UpdatedAt,
	)

	if err != nil {
		return err
	}

	job.ID = uint(lastID)

	return nil
}
