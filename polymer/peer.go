package polymer

import (
	"net/url"
)

// Peer
type Peer struct {
	WgPrivateKey   string
	Host           *url.URL
	IFaceBlackList []string
	WgIface        string
}

func NewPeer(privateKey string, host *url.URL, blackList []string, iface string) *Peer {
	return &Peer{
		WgPrivateKey:   privateKey,
		Host:           host,
		IFaceBlackList: blackList,
		WgIface:        iface,
	}
}
