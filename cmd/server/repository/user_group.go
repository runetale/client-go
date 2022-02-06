package repository

import (
	"github.com/Notch-Technologies/wizy/cmd/server/database"
	"github.com/Notch-Technologies/wizy/cmd/server/domain"
)

type UserGroupRepositoryManager interface {
	CreateOrganization(org *domain.OrgGroup) error
}

type UserGroupRepository struct {
	db *database.Sqlite
}

func NewUserGroupRepository(db *database.Sqlite) *UserGroupRepository {
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
