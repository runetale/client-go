package iface

import (
	"github.com/Notch-Technologies/wizy/wireguard"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

func CreateIface(i *Iface, ip, cidr string) error {
	addr := ip + "/" + cidr

	err := i.CreateWithUserSpace(addr)
	if err != nil {
		return err
	}

	key, err := wgtypes.ParseKey(i.WgPrivateKey)
	if err != nil {
		return err
	}

	fwmark := 0
	port := wireguard.WgPort

	config := wgtypes.Config{
		PrivateKey:   &key,
		ReplacePeers: false,
		FirewallMark: &fwmark,
		ListenPort:   &port,
	}

	return i.configureDevice(config)
}
