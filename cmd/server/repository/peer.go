package repository

import (
	"time"

	"github.com/Notch-Technologies/wizy/cmd/server/database"
	"github.com/Notch-Technologies/wizy/cmd/server/domain"
)

type PeerRepositoryManager interface {
	CreatePeer(peer *domain.Peer) error
}

type PeerRepository struct {
	db database.SQLExecuter
}

func NewPeerRepository(
	db database.SQLExecuter,
) *PeerRepository {
	return &PeerRepository{
		db: db,
	}
}

func (p *PeerRepository) CreatePeer(peer *domain.Peer) error {
	lastID, err := p.db.Exec(
		`INSERT INTO peers (
  			user_id,
  			setup_key_id,
  			organization_id,
  			user_group_id,
  			network_id,
  			ip,
  			created_at,
  			updated_at
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?)
		`,
		peer.UserID,
		peer.SetupKeyID,
		peer.OrganizationID,
		peer.UserGroupID,
		peer.NetworkID,
		peer.IP,
		peer.CreatedAt.In(time.UTC),
		peer.UpdatedAt.In(time.UTC),
	)

	if err != nil {
		return err
	}

	peer.ID = uint(lastID)

	return nil
}
