package repository

import (
	"database/sql"

	"github.com/Notch-Technologies/wizy/cmd/server/database"
	"github.com/Notch-Technologies/wizy/cmd/server/domain"
)

type OrgRepositoryManager interface {
	CreateOrganization(org *domain.OrgGroup) error
	FindByOrganizationID(orgID string) (*domain.OrgGroup, error)
}

type OrgRepository struct {
	db database.SQLExecuter
}

func NewOrgRepository(db database.SQLExecuter) *OrgRepository {
	return &OrgRepository{
		db: db,
	}
}

func (o *OrgRepository) CreateOrganization(org *domain.OrgGroup) error {
	lastID, err := o.db.Exec(`
	INSERT INTO orgs (
		name,
		display_name,
		org_id,
		created_at,
		updated_at
	) VALUES (?, ?, ?, ?, ?)
	`,
		org.Name,
		org.DisplayName,
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

func (o *OrgRepository) FindByOrganizationID(orgID string) (*domain.OrgGroup, error) {
	var (
		orgGroup domain.OrgGroup
	)

	row := o.db.QueryRow(
		`
			SELECT *
			FROM orgs
			WHERE
				org_id = ?
			LIMIT 1
		`, orgID)

	err := row.Scan(&orgGroup.ID, &orgGroup.Name, &orgGroup.DisplayName, &orgGroup.OrgID, &orgGroup.CreatedAt, &orgGroup.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrNoRows
		}
		return nil, err
	}

	return &orgGroup, nil
}
