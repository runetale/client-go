package wonderwall

import (
	"context"
	"sync"

	"github.com/Notch-Technologies/dotshake/dotengine/proxy"
	"github.com/Notch-Technologies/dotshake/dotlog"
	"github.com/pion/ice/v2"
)

type Conn struct {
	agent *ice.Agent

	remoteConn *ice.Conn

	wireproxy *proxy.WireProxy

	remoteWgPubKey string
	wgPubKey       string // local wg pubkey

	uname string
	pwd   string

	disconnectFunc func() error
	connectFunc    func() error

	ctx    context.Context
	cancel context.CancelFunc

	dotlog *dotlog.DotLog
}

func newConn(
	agent *ice.Agent,
	remoteWgPubKey string,
	wireproxy *proxy.WireProxy,
	wgPubKey string,
	disconnectFunc func() error,
	connectFunc func() error,
	dotlog *dotlog.DotLog,
) *Conn {
	ctx, cancel := context.WithCancel(context.Background())

	return &Conn{
		agent:     agent,
		wireproxy: wireproxy,

		disconnectFunc: disconnectFunc,
		connectFunc:    connectFunc,

		ctx:    ctx,
		cancel: cancel,

		dotlog: dotlog,
	}
}

type WonderWall struct {
	agent     *ice.Agent
	wireproxy *proxy.WireProxy

	conn *Conn

	mu *sync.Mutex

	dotlog *dotlog.DotLog
}

func NewWonderWall(
	dotlog *dotlog.DotLog,
) *WonderWall {
	return &WonderWall{
		mu:     &sync.Mutex{},
		dotlog: dotlog,
	}
}

func (w *WonderWall) setup(
	agent *ice.Agent,
	wireproxy *proxy.WireProxy,

	remoteWgPubKey string,
	wgPubKey string,

	disconnectFunc func() error,
	connectFunc func() error,
) {
	w.agent = agent
	w.wireproxy = wireproxy
	w.conn = newConn(
		agent,
		remoteWgPubKey,
		wireproxy,
		wgPubKey,
		disconnectFunc,
		connectFunc,
		w.dotlog,
	)
	return
}

// call after setup is called
//
func (w *WonderWall) Start(uname, pwd string) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	err := w.conn.connectFunc()
	if err != nil {
		return err
	}

	err = w.conn.agent.GatherCandidates()
	if err != nil {
		w.dotlog.Logger.Errorf("failed to collect candidates")
		return err
	}

	if w.conn.wgPubKey > w.conn.remoteWgPubKey {
		w.conn.remoteConn, err = w.conn.agent.Dial(w.conn.ctx, uname, pwd)
		if err != nil {
			w.dotlog.Logger.Errorf("failed to dial agent")
			return err
		}
	} else {
		w.conn.remoteConn, err = w.conn.agent.Accept(w.conn.ctx, uname, pwd)
		if err != nil {
			w.dotlog.Logger.Errorf("failed to accept agent")
			return err
		}
	}

	err = w.conn.wireproxy.StartProxy(w.conn.remoteConn)
	if err != nil {
		w.conn.dotlog.Logger.Errorf("failed to start proxy")
		return err
	}

	return nil
}

func (w *WonderWall) Close() error {
	w.mu.Lock()
	defer w.mu.Unlock()

	err := w.conn.disconnectFunc()
	if err != nil {
		w.dotlog.Logger.Errorf("failed to update connection status, this is not possible")
		return err
	}

	err = w.conn.wireproxy.Stop()
	if err != nil {
		w.dotlog.Logger.Errorf("failed to update connection status, this is not possible")
		return err
	}

	// close the ice agent connection
	w.conn.cancel()

	return nil
}
