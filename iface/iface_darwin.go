package iface

import (
	"github.com/Notch-Technologies/wizy/wireguard"
	"github.com/Notch-Technologies/wizy/wislog"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

func CreateIface(
	i *Iface, ip, cidr string,
	wislog *wislog.WisLog,
) error {
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
