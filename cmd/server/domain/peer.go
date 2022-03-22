package domain

import (
	"time"
)

type Peer struct {
	ID                  uint      `db:"id"`
	UserID              uint      `db:"user_id"`
	SetupKeyID          uint      `db:"setup_key_id"`
	AdminNetworkID      uint      `db:"admin_network_id"`
	UserGroupID         uint      `db:"user_group_id"`
	NetworkID           uint      `db:"network_id"`
	ClientMachinePubKey string    `db:"client_machine_pub_key"`
	WgPubKey            string    `db:"wg_pub_key"`
	IP                  string    `db:"ip"`
	CIDR                uint      `db:"cidr"`
	CreatedAt           time.Time `db:"created_at"`
	UpdatedAt           time.Time `db:"updated_at"`
}

func NewPeer(
	setupKeyID, networkID, userGroupID, userID, adminNetworkID uint,
	ip string, cidr uint, clientMachinePubKey, wgPubKey string,
) *Peer {
	return &Peer{
		UserID:              userID,
		SetupKeyID:          setupKeyID,
		AdminNetworkID:      adminNetworkID,
		UserGroupID:         userGroupID,
		ClientMachinePubKey: clientMachinePubKey,
		WgPubKey:            wgPubKey,
		NetworkID:           networkID,
		IP:                  ip,
		CIDR:                cidr,
		CreatedAt:           time.Now(),
		UpdatedAt:           time.Now(),
	}
}
