package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/Notch-Technologies/wizy/cmd/server/database"
	"github.com/Notch-Technologies/wizy/cmd/server/domain"
	"github.com/Notch-Technologies/wizy/types/key"
)

type SetupKeyRepositoryManager interface {
	CreateSetupKey(setupKey *domain.SetupKey) error
	FindBySetupKey(setupKey string) (*domain.SetupKey, error)
}

type SetupKeyRepository struct {
	db database.SQLExecuter
}

func NewSetupKeyRepository(
	db database.SQLExecuter,
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
  			updated_at
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

func (r *SetupKeyRepository) FindBySetupKey(setupKey string) (*domain.SetupKey, error) {
	fmt.Println(setupKey)
	var (
		sk domain.SetupKey
	)

	key := key.SetupKeyPrefix + setupKey

	row := r.db.QueryRow(
	`
		SELECT *
		FROM setup_keys
		WHERE
			key = ?
		LIMIT 1
	`, key)

	err := row.Scan(&sk.ID, &sk.UserID, &sk.Key, &sk.KeyType, &sk.Revoked, &sk.CreatedAt, &sk.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrNoRows
		}
		return nil, err
	}

	return &sk, nil
}
