package registry

import "github.com/Notch-Technologies/wizy/cmd/signaling/pb/negotiation"

type Peer struct {
	ClientMachineKey string

	Stream negotiation.Negotiation_ConnectStreamServer
}

func NewPeer(key string, stream negotiation.Negotiation_ConnectStreamServer) *Peer {
	return &Peer{
		ClientMachineKey: key,
		Stream:           stream,
	}
}
