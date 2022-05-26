package conn

import "sync"

type ConnStatus string

const Connect ConnStatus = "CONNECT"
const DisConnect ConnStatus = "DISCONNECT"

func (s ConnStatus) String() string {
	switch s {
	case Connect:
		return "connect"
	case DisConnect:
		return "disconnet"
	default:
		return "unreachable"
	}
}

type ConnectState struct {
	State ConnStatus
	conn  chan struct{}
	mu    sync.Mutex
}

func NewConnectedState() *ConnectState {
	return &ConnectState{
		State: DisConnect,
		mu:    sync.Mutex{},
	}
}

func (c *ConnectState) UpdateState(cs ConnStatus) ConnStatus {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.State = cs
	if c.State == cs {
		return DisConnect
	}

	return c.State
}

func (c *ConnectState) IsConnected() bool {
	if c.State.String() == Connect.String() {
		return true
	}
	return false
}

func (c *ConnectState) GetConnStatus() <-chan struct{} {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.conn == nil {
		c.conn = make(chan struct{})
	}

	return c.conn
}

func (c *ConnectState) Connected() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.State = Connect

	if c.conn != nil {
		close(c.conn)
		c.conn = nil
	}
}

func (c *ConnectState) DisConnected() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.State = DisConnect
}
