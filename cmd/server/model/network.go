package model

import (
	"encoding/json"
	"time"
)

type Network struct {
	// create unique id by wissy
	ID string `json:"id"`
	// network name
	Name string `json:"name"`
	// your ip
	IP string `json:"ip"`
	// network cidr
	CIDR string `json:"cidr"`
	// dns name
	DNS       string    `json:"dns"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewNetwork(id string, name string, ip string, cidr string,
	dns string) *Network {
	return &Network{
		ID:        id,
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
