package repository

import (
	"time"

	"github.com/Notch-Technologies/wizy/cmd/server/database"
	"github.com/Notch-Technologies/wizy/cmd/server/domain"
)

type SetupKeyRepositoryManager interface {
	CreateSetupKey(setupKey *domain.SetupKey) error
}

type SetupKeyRepository struct {
	db *database.Sqlite
}

func NewSetupKeyRepository(
	db *database.Sqlite,
) *SetupKeyRepository {
	return &SetupKeyRepository{
		db: db,
	}
}

func (r *SetupKeyRepository) CreateSetupKey(setupKey *domain.SetupKey) error {
	lastID, err := r.db.Exec(
		`INSERT INTO setup_keys (
  			user_id,
  			key,
  			key_type,
  			revoked,
  			created_at,
  			updated_at,
		) VALUES (?, ?, ?, ?, ?, ?)
		`,
		setupKey.UserID,
		setupKey.Key,
		setupKey.KeyType,
		setupKey.Revoked,
		setupKey.CreatedAt.In(time.UTC),
		setupKey.UpdatedAt.In(time.UTC),
	)

	if err != nil {
		return err
	}

	setupKey.ID = uint(lastID)

	return nil
}
