package peer

import (
	"sync"

	"github.com/Notch-Technologies/dotshake/client/grpc"
)

type Peer struct {
	serverClient grpc.ServerClientImpl

	mk string

	mu *sync.Mutex
	ch chan struct{}
}

func NewPeer(
	serverClient grpc.ServerClientImpl,
	mk string,
	ch chan struct{},
) *Peer {
	return &Peer{
		serverClient: serverClient,

		mk: mk,

		mu: &sync.Mutex{},
		ch: ch,
	}
}
