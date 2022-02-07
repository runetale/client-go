package repository

import (
	"database/sql"

	"github.com/Notch-Technologies/wizy/cmd/server/database"
	"github.com/Notch-Technologies/wizy/cmd/server/domain"
)

type UserGroupRepositoryManager interface {
	CreateUserGroup(group *domain.UserGroup) error
	FindByUserGroupID(id uint) (*domain.UserGroup, error)
}

type UserGroupRepository struct {
	db database.SQLExecuter
}

func NewUserGroupRepository(db database.SQLExecuter) *UserGroupRepository {
	return &UserGroupRepository{
		db: db,
	}
}

func (u *UserGroupRepository) CreateUserGroup(group *domain.UserGroup) error {
	lastID, err := u.db.Exec(`
	INSERT INTO user_groups (
		name,
		permission,
		created_at,
		updated_at
	) VALUES (?, ?, ?, ?)
	`,
		group.Name,
		group.Permission,
		group.CreatedAt,
		group.UpdatedAt,
	)

	if err != nil {
		return err
	}

	group.ID = uint(lastID)

	return nil
}

func (u *UserGroupRepository) FindByUserGroupID(id uint) (*domain.UserGroup, error) {
	var (
		userGroup domain.UserGroup
	)

	row := u.db.QueryRow(
		`
			SELECT *
			FROM user_groups
			WHERE
				id = ?
			LIMIT 1
		`, id)

	err := row.Scan(
		&userGroup.ID,
		&userGroup.Name,
		&userGroup.Permission,
		&userGroup.CreatedAt,
		&userGroup.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrNoRows
		}
		return nil, err
	}

	return &userGroup, nil
}
