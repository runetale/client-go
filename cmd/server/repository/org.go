package repository

import (
	"github.com/Notch-Technologies/wizy/cmd/server/database"
	"github.com/Notch-Technologies/wizy/cmd/server/domain"
)

type OrgRepositoryManager interface {
	CreateOrganization(org *domain.OrgGroup) error
	FindByOrganizationID(orgID string) (*domain.OrgGroup, error)
}

type OrgRepository struct {
	db *database.Sqlite
}

func NewOrgRepository(db *database.Sqlite) *OrgRepository {
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

	err := o.db.QueryRow(
		&orgGroup,
		`
			SELECT *
			FROM orgs
			WHERE
				org_id = ?
			LIMIT 1
		`, orgID)
	if err != nil {
		return nil, err
	}

	return &orgGroup, nil
}
