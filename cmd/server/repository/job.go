package repository

import (
	"github.com/Notch-Technologies/wizy/cmd/server/database"
	"github.com/Notch-Technologies/wizy/cmd/server/domain"
)

type JobRepositoryManager interface {
	CreateJob(job *domain.Job) error
}

type JobRepository struct {
	db *database.Sqlite
}

func NewJobRepository(db *database.Sqlite) *JobRepository {
	return &JobRepository{
		db: db,
	}
}

func (o *JobRepository) CreateJob(job *domain.Job) error {
	lastID, err := o.db.Exec(`
	INSERT INTO jobs (
		name,
		user_id,
		org_id,
		created_at,
		updated_at
	) VALUES (?, ?, ?, ?, ?)
	`,
		job.Name,
		job.UserID,
		job.OrgID,
		job.CreatedAt,
		job.UpdatedAt,
	)

	if err != nil {
		return err
	}

	job.ID = uint(lastID)

	return nil
}
