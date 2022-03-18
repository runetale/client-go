package repository

import (
	"database/sql"

	"github.com/Notch-Technologies/wizy/cmd/server/database"
	"github.com/Notch-Technologies/wizy/cmd/server/domain"
)

type AdminNetworkRepositoryCaller interface {
	CreateAdminNetwork(org *domain.AdminNetwork) error
	FindByOrganizationID(orgID string) (*domain.AdminNetwork, error)
}

type AdminNetworkRepository struct {
	db database.SQLExecuter
}

func NewAdminNetworkRepository(db database.SQLExecuter) *AdminNetworkRepository {
	return &AdminNetworkRepository{
		db: db,
	}
}

func (a *AdminNetworkRepository) CreateAdminNetwork(admin *domain.AdminNetwork) error {
	lastID, err := a.db.Exec(`
	INSERT INTO admin_networks (
		name,
		org_id,
		created_at,
		updated_at
	) VALUES (?, ?, ?, ?)
	`,
		admin.Name,
		admin.OrgID,
		admin.CreatedAt,
		admin.UpdatedAt,
	)

	if err != nil {
		return err
	}

	admin.ID = uint(lastID)

	return nil
}

func (a *AdminNetworkRepository) FindByOrganizationID(orgID string) (*domain.AdminNetwork, error) {
	var (
		admin domain.AdminNetwork
	)

	row := a.db.QueryRow(
		`
			SELECT *
			FROM admin_networks
			WHERE
				org_id = ?
			LIMIT 1
		`, orgID)

	err := row.Scan(
		&admin.ID,
		&admin.Name,
		&admin.OrgID,
		&admin.CreatedAt,
		&admin.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrNoRows
		}
		return nil, err
	}

	return &admin, nil
}
