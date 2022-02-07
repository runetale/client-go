package domain

import "time"

type Peer struct {
	ID          uint      `db:"id"`
	UserID      uint      `db:"user_id"`
	SetupKeyID  uint      `db:"setup_key_id"`
	OrgID       uint      `db:"org_id"`
	UserGroupID uint      `db:"user_group_id"`
	NetworkID   uint      `db:"network_id"`
	IP          string    `db:"ip"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

func NewPeer(setupKey, setupKeyID, networkID, userGroupID,
	userID, orgID uint, ip string) *Peer {
	return &Peer{
		UserID:      userID,
		SetupKeyID:  setupKeyID,
		OrgID:       orgID,
		UserGroupID: userGroupID,
		NetworkID:   networkID,
		CreatedAt:   time.Now(),
		IP:          ip,
		UpdatedAt:   time.Now(),
	}
}
