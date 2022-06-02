package conn

import (
	"context"
	"sync"

	"github.com/Notch-Technologies/dotshake/dotlog"
	"github.com/pion/ice/v2"
)

type Conn struct {
	agent *ice.Agent

	remoteConn *ice.Conn

	wireproxy *WireProxy

	remoteWgPubKey string
	wgPubKey       string // local wg pubkey

	uname string
	pwd   string

	closeCh chan struct{}

	disconnectFunc func() error
	connectFunc    func() error

	ctx    context.Context
	cancel context.CancelFunc

	mu *sync.Mutex

	dotlog *dotlog.DotLog
}

func NewConn(
	agent *ice.Agent,
	remoteWgPubKey string,
	wireproxy *WireProxy,
	wgPubKey string,
	closeCh chan struct{},
	disconnectFunc func() error,
	connectFunc func() error,
	dotlog *dotlog.DotLog,
) *Conn {
	ctx, cancel := context.WithCancel(context.Background())

	return &Conn{
		agent:     agent,
		wireproxy: wireproxy,

		closeCh: closeCh,

		disconnectFunc: disconnectFunc,
		connectFunc:    connectFunc,

		ctx:    ctx,
		cancel: cancel,

		mu: &sync.Mutex{},

		dotlog: dotlog,
	}
}

func (c *Conn) Start(uname, pwd string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	err := c.connectFunc()
	if err != nil {
		c.dotlog.Logger.Errorf("failed to update connection status, this is not possible")
		return err
	}

	err = c.agent.GatherCandidates()
	if err != nil {
		c.dotlog.Logger.Errorf("failed to collect candidates")
		return err
	}

	if c.wgPubKey > c.remoteWgPubKey {
		c.remoteConn, err = c.agent.Dial(c.ctx, uname, pwd)
		if err != nil {
			c.dotlog.Logger.Errorf("failed to dial agent")
			return err
		}
	} else {
		c.remoteConn, err = c.agent.Accept(c.ctx, uname, pwd)
		if err != nil {
			c.dotlog.Logger.Errorf("failed to accept agent")
			return err
		}
	}

	err = c.wireproxy.StartProxy(c.remoteConn)
	if err != nil {
		c.dotlog.Logger.Errorf("failed to start proxy")
		return err
	}

	return nil
}

func (c *Conn) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	err := c.disconnectFunc()
	if err != nil {
		c.dotlog.Logger.Errorf("failed to update connection status, this is not possible")
		return err
	}

	err = c.wireproxy.Stop()
	if err != nil {
		c.dotlog.Logger.Errorf("failed to update connection status, this is not possible")
		return err
	}

	// close the ice agent connection
	c.cancel()

	return nil
}
