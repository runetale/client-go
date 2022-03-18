package repository

import (
	"database/sql"
	"time"

	"github.com/Notch-Technologies/wizy/cmd/server/database"
	"github.com/Notch-Technologies/wizy/cmd/server/domain"
)

type PeerRepositoryManager interface {
	CreatePeer(peer *domain.Peer) error
	FindBySetupKeyID(id uint) (*domain.Peer, error)
	FindByClientPubKey(clientPubKey string) (*domain.Peer, error)
	FindPeersByClientPubKey(clientPubKey string) ([]*domain.Peer, error)
	FindPeersByOrganizationID(organizationID string) ([]*domain.Peer, error)
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
  			admin_network_id,
  			user_group_id,
  			network_id,
			client_pub_key,
  			wg_pub_key,
  			ip,
			cidr,
  			created_at,
  			updated_at
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
		`,
		peer.UserID,
		peer.SetupKeyID,
		peer.AdminNetworkID,
		peer.UserGroupID,
		peer.NetworkID,
		peer.ClientPubKey,
		peer.WgPubKey,
		peer.IP,
		peer.CIDR,
		peer.CreatedAt.In(time.UTC),
		peer.UpdatedAt.In(time.UTC),
	)

	if err != nil {
		return err
	}

	peer.ID = uint(lastID)

	return nil
}

func (p *PeerRepository) FindBySetupKeyID(id uint, clientPubKey string) (*domain.Peer, error) {
	var (
		peer domain.Peer
	)

	row := p.db.QueryRow(
		`
			SELECT *
			FROM peers
			WHERE
  				setup_key_id = ? AND
				client_pub_key = ?
			LIMIT 1
		`, id, clientPubKey)
	err := row.Scan(
		&peer.ID,
		&peer.UserID,
		&peer.SetupKeyID,
		&peer.AdminNetworkID,
		&peer.UserGroupID,
		&peer.NetworkID,
		&peer.ClientPubKey,
		&peer.WgPubKey,
		&peer.IP,
		&peer.CIDR,
		&peer.CreatedAt,
		&peer.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrNoRows
		}
		return nil, err
	}

	return &peer, nil
}

func (p *PeerRepository) FindByClientPubKey(clientPubKey string) (*domain.Peer, error) {
	var (
		peer domain.Peer
	)

	row := p.db.QueryRow(
		`
			SELECT *
			FROM peers
			WHERE
				client_pub_key = ?
			LIMIT 1
		`, clientPubKey)
	err := row.Scan(
		&peer.ID,
		&peer.UserID,
		&peer.SetupKeyID,
		&peer.AdminNetworkID,
		&peer.UserGroupID,
		&peer.NetworkID,
		&peer.ClientPubKey,
		&peer.WgPubKey,
		&peer.IP,
		&peer.CIDR,
		&peer.CreatedAt,
		&peer.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrNoRows
		}
		return nil, err
	}

	return &peer, nil
}

func (p *PeerRepository) FindPeersByClientPubKey(clientPubKey string) ([]*domain.Peer, error) {
	peers := make([]*domain.Peer, 0)

	rows, err := p.db.Query(
		`
			SELECT *
			FROM peers
			WHERE
				client_pub_key = ?
		`, clientPubKey)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrNoRows
		}
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		peer := new(domain.Peer)
		if err := rows.Scan(
			&peer.ID,
			&peer.UserID,
			&peer.SetupKeyID,
			&peer.AdminNetworkID,
			&peer.UserGroupID,
			&peer.NetworkID,
			&peer.ClientPubKey,
			&peer.WgPubKey,
			&peer.IP,
			&peer.CIDR,
			&peer.CreatedAt,
			&peer.UpdatedAt,
		); err != nil {
			return nil, err
		}
		peers = append(peers, peer)
	}
	if err := rows.Err(); err != nil {
		panic(err)
	}

	return peers, nil
}

func (p *PeerRepository) FindPeersByOrganizationID(organizationID uint) ([]*domain.Peer, error) {
	peers := make([]*domain.Peer, 0)

	rows, err := p.db.Query(
		`
			SELECT *
			FROM peers
			WHERE organization_id = ?
		`, organizationID)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrNoRows
		}
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		peer := new(domain.Peer)
		if err := rows.Scan(
			&peer.ID,
			&peer.UserID,
			&peer.SetupKeyID,
			&peer.AdminNetworkID,
			&peer.UserGroupID,
			&peer.NetworkID,
			&peer.ClientPubKey,
			&peer.WgPubKey,
			&peer.IP,
			&peer.CIDR,
			&peer.CreatedAt,
			&peer.UpdatedAt,
		); err != nil {
			return nil, err
		}
		peers = append(peers, peer)
	}
	if err := rows.Err(); err != nil {
		panic(err)
	}

	return peers, nil
}
