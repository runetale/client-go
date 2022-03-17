package repository

import (
	"database/sql"

	"github.com/Notch-Technologies/wizy/cmd/server/database"
	"github.com/Notch-Technologies/wizy/cmd/server/domain"
)

type OrgRepositoryManager interface {
	CreateOrganization(org *domain.Organization) error
	FindByOrganizationID(orgID string) (*domain.Organization, error)
}

type OrgRepository struct {
	db database.SQLExecuter
}

func NewOrgRepository(db database.SQLExecuter) *OrgRepository {
	return &OrgRepository{
		db: db,
	}
}

func (o *OrgRepository) CreateOrganization(org *domain.Organization) error {
	lastID, err := o.db.Exec(`
	INSERT INTO organizations (
		name,
		org_id,
		created_at,
		updated_at
	) VALUES (?, ?, ?, ?)
	`,
		org.Name,
		org.OrgID,
		org.CreatedAt,
		org.UpdatedAt,
	)

	if err != nil {
		return err
	}

	org.ID = uint(lastID)

	return nil
}

func (o *OrgRepository) FindByOrganizationID(orgID string) (*domain.Organization, error) {
	var (
		org domain.Organization
	)

	row := o.db.QueryRow(
		`
			SELECT *
			FROM organizations
			WHERE
				org_id = ?
			LIMIT 1
		`, orgID)

	err := row.Scan(&org.ID, &org.Name, &org.OrgID, &org.CreatedAt, &org.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrNoRows
		}
		return nil, err
	}

	return &org, nil
}
