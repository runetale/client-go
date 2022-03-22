package repository

import (
	"database/sql"

	"github.com/Notch-Technologies/wizy/cmd/server/database"
	"github.com/Notch-Technologies/wizy/cmd/server/domain"
)

type JobRepositoryCaller interface {
	CreateJob(job *domain.Job) error
	FindByID(id uint) (*domain.Job, error)
	FindByAdminNetworkID(id uint) (*domain.Job, error)
}

type JobRepository struct {
	db database.SQLExecuter
}

func NewJobRepository(db database.SQLExecuter) *JobRepository {
	return &JobRepository{
		db: db,
	}
}

func (r *JobRepository) CreateJob(job *domain.Job) error {
	lastID, err := r.db.Exec(`
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

func (r *JobRepository) FindByID(id uint) (*domain.Job, error) {
	var (
		job domain.Job
	)

	row := r.db.QueryRow(
		`
			SELECT *
			FROM jobs
			WHERE
				id = ?
			LIMIT 1
		`, id)

	err := row.Scan(
		&job.ID,
		&job.AdminNetworkID,
		&job.Name,
		&job.CreatedAt,
		&job.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrNoRows
		}
		return nil, err
	}
	return &job, nil
}

func (r *JobRepository) FindByAdminNetworkID(id uint) (*domain.Job, error) {
	var (
		job domain.Job
	)

	row := r.db.QueryRow(
		`
			SELECT *
			FROM jobs
			WHERE
				admin_network_id = ?
			LIMIT 1
		`, id)

	err := row.Scan(
		&job.ID,
		&job.AdminNetworkID,
		&job.Name,
		&job.CreatedAt,
		&job.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrNoRows
		}
		return nil, err
	}
	return &job, nil
}
