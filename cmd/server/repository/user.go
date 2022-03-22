package repository

import (
	"database/sql"

	"github.com/Notch-Technologies/wizy/cmd/server/database"
	"github.com/Notch-Technologies/wizy/cmd/server/domain"
)

type UserRepositoryCaller interface {
	CreateUser(user *domain.User) error
	FindByProviderID(providerID string) (*domain.User, error)
	FindByUserID(userID uint) (*domain.User, error)
}

type UserRepository struct {
	db database.SQLExecuter
}

func NewUserRepository(db database.SQLExecuter) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) CreateUser(user *domain.User) error {
	lastID, err := u.db.Exec(`
	INSERT INTO users (
		provider_id,
		admin_network_id,
		network_id,
		user_group_id,
		role_id,
		provider,
		email,
		created_at,
		updated_at
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`,
		user.ProviderID,
		user.AdminNetworkID,
		user.NetworkID,
		user.UserGroupID,
		user.RoleID,
		user.Provider,
		user.Email,
		user.CreatedAt,
		user.UpdatedAt,
	)

	if err != nil {
		return err
	}

	user.ID = uint(lastID)

	return nil
}

func (u *UserRepository) FindByProviderID(providerID string) (*domain.User, error) {
	var (
		user domain.User
	)

	row := u.db.QueryRow(
		`
			SELECT *
			FROM users
			WHERE
				provider_id = ?
			LIMIT 1
		`, providerID)

	err := row.Scan(
		&user.ID,
		&user.ProviderID,
		&user.AdminNetworkID,
		&user.NetworkID,
		&user.UserGroupID,
		&user.RoleID,
		&user.Provider,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrNoRows
		}
		return nil, err
	}

	return &user, nil
}

func (u *UserRepository) FindByUserID(userID uint) (*domain.User, error) {
	var (
		user domain.User
	)

	row := u.db.QueryRow(
		`
			SELECT *
			FROM users
			WHERE
				id = ?
			LIMIT 1
		`, userID)

	err := row.Scan(
		&user.ID,
		&user.ProviderID,
		&user.AdminNetworkID,
		&user.NetworkID,
		&user.UserGroupID,
		&user.RoleID,
		&user.Provider,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrNoRows
		}
		return nil, err
	}

	return &user, nil
}
