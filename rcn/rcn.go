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
	"github.com/Notch-Technologies/dotshake/rcn/unixsock"
)

type Rcn struct {
	scp *controlplane.ControlPlane
	mk  string
	mu  *sync.Mutex
	ch  chan struct{}

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
	return &Rcn{
		scp: controlplane.NewControlPlane(signalClient, mk, clientConf, ch, dotlog),

		mk: mk,

		mu: mu,
		ch: ch,

		dotlog: dotlog,
	}
}

func (p *Rcn) connectSock() {
	sock := unixsock.NewPolymerSock(p.dotlog, p.ch, p.scp)
	go func() {
		err := sock.Connect()
		if err != nil {
			p.dotlog.Logger.Errorf("failed to connect rcn sock. %s", err.Error())
		}
		p.dotlog.Logger.Debugf("connection with sock connect has been disconnected")
	}()
}

func (p *Rcn) Start() {
	err := p.scp.ConfigureStunTurnConf()
	if err != nil {
		p.dotlog.Logger.Errorf("failed to set up puncher. %s", err.Error())
	}

	p.scp.WaitForRemoteConn()

	p.connectSock()

	p.scp.ConnectSignalServer()
}
