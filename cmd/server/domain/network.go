package domain

import (
	"encoding/json"
	"time"
)

type Network struct {
	ID             uint      `db:"id"`
	AdminNetworkID uint      `db:"admin_network_id"`
	Name           string    `db:"name"`
	IP             string    `db:"ip"`
	CIDR           uint      `db:"cidr"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}

func NewNetwork(
	adminNetworkID uint, name string,
	ip string, cidr uint,
) *Network {
	return &Network{
		AdminNetworkID: adminNetworkID,
		Name:           name,
		IP:             ip,
		CIDR:           cidr,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
}

func (n Network) MarshalBinary() ([]byte, error) {
	return json.Marshal(n)
}
