package repository

import (
	"database/sql"

	"github.com/Notch-Technologies/wizy/cmd/server/database"
	"github.com/Notch-Technologies/wizy/cmd/server/domain"
)

type RoleRepositoryCaller interface {
	CreateRole(role *domain.Role) error
	FindByAdminNetworkID(adminNetworkID uint) (*domain.Role, error)
}

type RoleRepository struct {
	db database.SQLExecuter
}

func NewRoleRepository(db database.SQLExecuter) *RoleRepository {
	return &RoleRepository{
		db: db,
	}
}

func (r *RoleRepository) CreateRole(role *domain.Role) error {
	lastID, err := r.db.Exec(`
	INSERT INTO roles (
		admin_network_id,
		name,
		permission,
		created_at,
		updated_at
	) VALUES (?, ?, ?, ?, ?)
	`,
		role.AdminNetworkID,
		role.Name,
		role.Permission,
		role.CreatedAt,
		role.UpdatedAt,
	)

	if err != nil {
		return err
	}

	role.ID = uint(lastID)

	return nil
}

func (r *RoleRepository) FindByAdminNetworkID(adminNetworkID uint) (*domain.Role, error) {
	var (
		role domain.Role
	)

	row := r.db.QueryRow(
		`
			SELECT *
			FROM roles
			WHERE
				admin_network_id = ?
			LIMIT 1
		`, adminNetworkID)

	err := row.Scan(
		&role.ID,
		&role.AdminNetworkID,
		&role.Name,
		&role.Permission,
		&role.CreatedAt,
		&role.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrNoRows
		}
		return nil, err
	}
	return &role, nil
}
