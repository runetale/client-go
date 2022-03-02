package engine

import (
	"github.com/Notch-Technologies/wizy/core"
	"github.com/Notch-Technologies/wizy/wireguard"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

type EngineConfig struct {
	WgPort  int
	WgIface string
	// WgAddr is a Wireguard local address (Wiretrustee Network IP)
	WgAddr string
	// WgPrivateKey is a Wireguard private key of our peer (it MUST never leave the machine)
	WgPrivateKey wgtypes.Key
	// IFaceBlackList is a list of network interfaces to ignore when discovering connection candidates (ICE related)
	IFaceBlackList map[string]struct{}

	PreSharedKey *wgtypes.Key
}

func NewEngineConfig(key wgtypes.Key, config *core.ClientCore, wgAddr string) *EngineConfig {
	iFaceBlackList := make(map[string]struct{})
	for i := 0; i < len(config.IfaceBlackList); i += 2 {
		iFaceBlackList[config.IfaceBlackList[i]] = struct{}{}
	}

	return &EngineConfig{
		WgIface:        config.TUNName,
		WgAddr:         wgAddr,
		IFaceBlackList: iFaceBlackList,
		WgPrivateKey:   key,
		WgPort:         wireguard.WgPort,
	}
}
