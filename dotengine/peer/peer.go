package peer

import (
	"sync"

	"github.com/Notch-Technologies/client-go/notch/dotshake/v1/machine"
	"github.com/Notch-Technologies/dotshake/client/grpc"
	"github.com/Notch-Technologies/dotshake/dotlog"
	"github.com/Notch-Technologies/dotshake/rcn/unixsock"
)

type Peer struct {
	serverClient grpc.ServerClientImpl

	mk string

	// conn *conn.Conn

	mu *sync.Mutex
	ch chan struct{}

	dotlog *dotlog.DotLog

	sock *unixsock.PolymerSock
}

func NewPeer(
	serverClient grpc.ServerClientImpl,
	mk string,
	dotlog *dotlog.DotLog,
	sock *unixsock.PolymerSock,
) *Peer {
	ch := make(chan struct{})

	return &Peer{
		serverClient: serverClient,

		mk: mk,

		// conn: conn.NewConn(dotlog, sock),

		mu: &sync.Mutex{},
		ch: ch,

		dotlog: dotlog,

		sock: sock,
	}
}

func (p *Peer) Up() {
	go func() {
		err := p.SyncMachine(func(res *machine.SyncMachinesResponse) error {
			p.mu.Lock()
			defer p.mu.Unlock()

			p.dotlog.Logger.Debugf("connected sync machine")

			err := p.conn.Start(res.GetRemotePeers(), res.GetIp(), res.GetCidr())
			if err != nil {
				return err
			}

			p.dotlog.Logger.Debugf("connected peer connection")

			return nil
		})
		if err != nil {
			close(p.ch)
			return
		}
	}()
}

func (p *Peer) SyncMachine(handler func(res *machine.SyncMachinesResponse) error) error {
	err := p.serverClient.SyncMachines(p.mk, handler)
	if err != nil {
		return err
	}

	return nil
}
