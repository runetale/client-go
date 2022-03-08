package iface

import (
	"github.com/Notch-Technologies/wizy/wireguard"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

func CreateIface(i *Iface, address string) error {
	err := i.CreateWithUserSpace(address)
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
