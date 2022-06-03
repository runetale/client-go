package peer

import (
	"sync"

	"github.com/Notch-Technologies/client-go/notch/dotshake/v1/machine"
	"github.com/Notch-Technologies/dotshake/client/grpc"
	"github.com/Notch-Technologies/dotshake/dotlog"
	"github.com/Notch-Technologies/dotshake/rcn/rcnsock"
)

type Peer struct {
	serverClient grpc.ServerClientImpl

	mk string

	// conn *conn.Conn

	mu *sync.Mutex
	ch chan struct{}

	dotlog *dotlog.DotLog

	sock *rcnsock.RcnSock
}

func NewPeer(
	serverClient grpc.ServerClientImpl,
	mk string,
	dotlog *dotlog.DotLog,
	sock *rcnsock.RcnSock,
) *Peer {
	ch := make(chan struct{})

	return &Peer{
		serverClient: serverClient,

		mk: mk,

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

			if res.GetRemotePeers() != nil {
				err := p.sock.DialPeerSock(&rcnsock.PeerSock{
					Commands:    rcnsock.SyncRemotePeerConnecting,
					RemotePeers: res.GetRemotePeers(),
				})
				if err != nil {
					p.dotlog.Logger.Debugf("failed to sync remote peer connectiong")
					return err
				}
			}

			err := p.sock.DialPeerSock(&rcnsock.PeerSock{
				Commands:    rcnsock.SetupRemotePeersConn,
				RemotePeers: res.GetRemotePeers(),
				Ip:          res.GetIp(),
				Cidr:        res.GetCidr(),
			})
			if err != nil {
				p.dotlog.Logger.Debugf("failed to setup remote peer conn")
				return err
			}

			p.dotlog.Logger.Debugf("connected sync machine")

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
