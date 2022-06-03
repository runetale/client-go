package rcn

// rcn package is realtime communication nucleus
// provides communication status and P2P communication aids
//

import (
	"sync"

	"github.com/Notch-Technologies/dotshake/client/grpc"
	"github.com/Notch-Technologies/dotshake/conf"
	"github.com/Notch-Technologies/dotshake/dotlog"
	"github.com/Notch-Technologies/dotshake/rcn/controlplane"
	"github.com/Notch-Technologies/dotshake/rcn/rcnsock"
)

type Rcn struct {
	cp   *controlplane.ControlPlane
	sock *rcnsock.RcnSock

	mk string
	mu *sync.Mutex

	dotlog *dotlog.DotLog
}

func NewRcn(
	signalClient grpc.SignalClientImpl,
	clientConf *conf.ClientConf,
	mk string,
	ch chan struct{},
	mu *sync.Mutex,
	dotlog *dotlog.DotLog,
) *Rcn {
	cp := controlplane.NewControlPlane(signalClient, mk, clientConf, ch, dotlog)
	return &Rcn{
		cp:   cp,
		sock: rcnsock.NewRcnSock(dotlog, ch, cp),

		mk: mk,

		mu: mu,

		dotlog: dotlog,
	}
}

// listen to rcn sock
//
func (p *Rcn) connectRcnSock() {
	go func() {
		err := p.sock.Connect()
		if err != nil {
			p.dotlog.Logger.Errorf("failed to connect rcn sock. %s", err.Error())
		}
		p.dotlog.Logger.Debugf("connection with rcn sock connect has been disconnected")
	}()
}

func (p *Rcn) Start() {
	err := p.cp.ConfigureStunTurnConf()
	if err != nil {
		p.dotlog.Logger.Errorf("failed to set up puncher. %s", err.Error())
	}

	go p.cp.WaitForRemoteConn()

	p.connectRcnSock()

	p.cp.ConnectSignalServer()
}
