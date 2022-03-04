package domain

import (
	"encoding/json"
	"time"
)

type Network struct {
	ID        uint      `db:"id"`
	Name      string    `db:"name"`
	IP        string    `db:"ip"`
	CIDR      uint      `db:"cidr"`
	DNS       string    `db:"dns"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func NewNetwork(
	name string, ip string,
	cidr uint, dns string, mask string,
) *Network {
	return &Network{
		Name:      name,
		IP:        ip,
		CIDR:      cidr,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (n Network) MarshalBinary() ([]byte, error) {
	return json.Marshal(n)
}
