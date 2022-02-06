package repository

import (
	"github.com/Notch-Technologies/wizy/cmd/server/database"
	"github.com/Notch-Technologies/wizy/cmd/server/domain"
)

type NetworkRepositoryManager interface {
	CreateNetwork(network *domain.Network) (*domain.Network, error)
}

type NetworkRepository struct {
	db *database.Sqlite
}

func NewNetworkRepository(db *database.Sqlite) *NetworkRepository {
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
		created_at,
		updated_at
	) VALUES (?, ?, ?, ?, ?)
	`,
		network.Name,
		network.IP,
		network.CIDR,
		network.CreatedAt,
		network.UpdatedAt,
	)

	if err != nil {
		return err
	}

	network.ID = uint(lastID)

	return nil
}
