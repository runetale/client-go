package repository

import (
	"github.com/Notch-Technologies/wizy/cmd/server/database"
	"github.com/Notch-Technologies/wizy/cmd/server/domain"
)

type UserRepositoryManager interface {
	CreateUser(user *domain.User) error
}

type UserRepository struct {
	db *database.Sqlite
}

func NewUserRepository(db *database.Sqlite) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) CreateUser(user *domain.User) error {
	lastID, err := u.db.Exec(`
	INSERT INTO users (
		provider_id,
		provider,
		org_group_id,
		network_id,
		user_group_id,
		permission,
		created_at,
		updated_at
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`,
		user.ProviderID,
		user.Provider,
		user.OrgGroupID,
		user.NetworkID,
		user.UserGroupID,
		user.Permission,
		user.CreatedAt,
		user.UpdatedAt,
	)

	if err != nil {
		return err
	}

	user.ID = uint(lastID)

	return nil
}
