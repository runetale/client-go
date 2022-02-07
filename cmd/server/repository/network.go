package repository

import (
	"database/sql"

	"github.com/Notch-Technologies/wizy/cmd/server/database"
	"github.com/Notch-Technologies/wizy/cmd/server/domain"
)

type NetworkRepositoryManager interface {
	CreateNetwork(network *domain.Network) (*domain.Network, error)
	FindByNetworkID(id uint) (*domain.Network, error)
}

type NetworkRepository struct {
	db database.SQLExecuter
}

func NewNetworkRepository(db database.SQLExecuter) *NetworkRepository {
	return &NetworkRepository{
		db: db,
	}
}

func (n *NetworkRepository) CreateNetwork(network *domain.Network) error {
	lastID, err := n.db.Exec(`
	INSERT INTO networks (
		name,
		ip,
		cidr,
		dns,
		created_at,
		updated_at
	) VALUES (?, ?, ?, ?, ?, ?)
	`,
		network.Name,
		network.IP,
		network.CIDR,
		network.DNS,
		network.CreatedAt,
		network.UpdatedAt,
	)

	if err != nil {
		return err
	}

	network.ID = uint(lastID)

	return nil
}

func (n *NetworkRepository) FindByNetworkID(id uint) (*domain.Network, error) {
	var (
		network domain.Network
	)

	row := n.db.QueryRow(
		`
			SELECT *
			FROM networks
			WHERE
				id = ?
			LIMIT 1
		`, id)

	err := row.Scan(
		&network.ID, &network.Name, &network.IP, &network.CIDR,
		&network.DNS, &network.CreatedAt, &network.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrNoRows
		}
		return nil, err
	}
	return &network, nil
}
