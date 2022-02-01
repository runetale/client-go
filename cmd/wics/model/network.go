package model

import "time"

type Network struct {
	ID        string
	Name      string
	IP        string
	CIDR      string
	DNS       string
	CreatedAt time.Time
	UpdatedAt time.Time
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
