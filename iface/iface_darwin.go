package iface

import (
	"github.com/Notch-Technologies/dotshake/dotlog"
	"github.com/Notch-Technologies/dotshake/wireguard"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

func CreateIface(
	i *Iface, ip, cidr string,
	dotlog *dotlog.DotLog,
) error {
	addr := ip + "/" + cidr

	err := i.CreateWithUserSpace(addr)
	if err != nil {
		dotlog.Logger.Fatalf("failed to user space. because %v", err)
		return err
	}

	key, err := wgtypes.ParseKey(i.WgPrivateKey)
	if err != nil {
		dotlog.Logger.Fatalf("failed to parsing wireguard private key %v", err)
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
