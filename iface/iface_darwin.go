package iface

import (
	"github.com/Notch-Technologies/wizy/wireguard"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

func CreateIface(ifaceName, privateKey, address string) error {
	err := CreateWithUserSpace(ifaceName, address)
	if err != nil {
		return err
	}

	key, err := wgtypes.ParseKey(privateKey)
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

	return configureDevice(ifaceName, config)
}
