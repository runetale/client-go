package engine

import (
	"github.com/Notch-Technologies/dotshake/core"
	"github.com/Notch-Technologies/dotshake/wireguard"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

type EngineConfig struct {
	WgPort         int
	WgIface        string
	WgAddr         string
	WgPrivateKey   wgtypes.Key
	IFaceBlackList map[string]struct{}

	PreSharedKey *wgtypes.Key
}

func NewEngineConfig(key wgtypes.Key, config *core.ClientCore, ip, cidr string) *EngineConfig {
	iFaceBlackList := make(map[string]struct{})
	for i := 0; i < len(config.IfaceBlackList); i += 2 {
		iFaceBlackList[config.IfaceBlackList[i]] = struct{}{}
	}

	return &EngineConfig{
		WgIface:        config.TunName,
		WgAddr:         ip + "/" + cidr,
		IFaceBlackList: iFaceBlackList,
		WgPrivateKey:   key,
		WgPort:         wireguard.WgPort,
	}
}
