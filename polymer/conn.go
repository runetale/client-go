package polymer

import (
	"context"
	"fmt"
	"io"
	"net"
	"sync"
	"time"

	"github.com/pion/ice/v2"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

type ConnStatus int

func (s ConnStatus) String() string {
	switch s {
	case StatusConnecting:
		return "StatusConnecting"
	case StatusConnected:
		return "StatusConnected"
	case StatusDisconnected:
		return "StatusDisconnected"
	default:
		fmt.Printf("unknown status: %d\n", s)
		return "INVALID_PEER_CONNECTION_STATUS"
	}
}

const (
	StatusConnected = iota
	StatusConnecting
	StatusDisconnected
)

const DefaultWgKeepAlive = 25 * time.Second

type ProxyConfig struct {
	WgListenAddr string
	RemoteKey    string
	WgInterface  string
	AllowedIps   string
	PreSharedKey *wgtypes.Key
}

type Proxy interface {
	io.Closer
	// Start creates a local remoteConn and starts proxying data from/to remoteConn
	Start(remoteConn net.Conn) error
}

// ConnConfig is a peer Connection configuration
type ConnConfig struct {

	// Key is a public key of a remote peer
	Key string
	// LocalKey is a public key of a local peer
	LocalKey string

	// StunTurn is a list of STUN and TURN URLs
	StunTurn []*ice.URL

	// InterfaceBlackList is a list of machine interfaces that should be filtered out by ICE Candidate gathering
	// (e.g. if eth0 is in the list, host candidate of this interface won't be used)
	InterfaceBlackList []string

	Timeout time.Duration

	ProxyConfig ProxyConfig
}

// IceCredentials ICE protocol credentials struct
type IceCredentials struct {
	UFrag string
	Pwd   string
}

type Conn struct {
	config ConnConfig
	mu     sync.Mutex

	// signalCandidate is a handler function to signal remote peer about local connection candidate
	signalCandidate func(candidate ice.Candidate) error
	// signalOffer is a handler function to signal remote peer our connection offer (credentials)
	signalOffer  func(uFrag string, pwd string) error
	signalAnswer func(uFrag string, pwd string) error

	// remoteOffersCh is a channel used to wait for remote credentials to proceed with the connection
	remoteOffersCh chan IceCredentials
	// remoteAnswerCh is a channel used to wait for remote credentials answer (confirmation of our offer) to proceed with the connection
	remoteAnswerCh     chan IceCredentials
	closeCh            chan struct{}
	ctx                context.Context
	notifyDisconnected context.CancelFunc

	agent  *ice.Agent
	status ConnStatus

	proxy Proxy
}

func NewConn(config ConnConfig) (*Conn, error) {
	return &Conn{
		config:         config,
		mu:             sync.Mutex{},
		status:         StatusDisconnected,
		closeCh:        make(chan struct{}),
		remoteOffersCh: make(chan IceCredentials),
		remoteAnswerCh: make(chan IceCredentials),
	}, nil
}

func (conn *Conn) Close() error {
	conn.mu.Lock()
	defer conn.mu.Unlock()
	select {
	case conn.closeCh <- struct{}{}:
	default:
		fmt.Printf("closing not started coonection %s\n", conn.config.Key)
	}
	return nil
}

func (conn *Conn) SetSignalOffer(handler func(uFrag string, pwd string) error) {
	conn.signalOffer = handler
}

func (conn *Conn) cleanup() error {
	fmt.Printf("trying to cleanup %s", conn.config.Key)
	conn.mu.Lock()
	defer conn.mu.Unlock()

	if conn.agent != nil {
		err := conn.agent.Close()
		if err != nil {
			return err
		}
		conn.agent = nil
	}

	if conn.proxy != nil {
		err := conn.proxy.Close()
		if err != nil {
			return err
		}
		conn.proxy = nil
	}

	if conn.notifyDisconnected != nil {
		conn.notifyDisconnected()
		conn.notifyDisconnected = nil
	}

	conn.status = StatusDisconnected

	fmt.Printf("cleaned up connection to peer %s", conn.config.Key)

	return nil
}

func (conn *Conn) Open() error {
	defer func() {
		err := conn.cleanup()
		if err != nil {
			fmt.Printf("error while cleaning up peer connection %s: %v\n", conn.config.Key, err)
			return
		}
	}()


	err := conn.reCreateAgent()
	if err != nil {
		return err
	}

	err = conn.sendOffer()
	if err != nil {
		return err
	}

	fmt.Printf("connection offer sent to peer %s, waiting for the confirmation\n", conn.config.Key)

	return nil
}

func interfaceFilter(blackList []string) func(string) bool {
	var blackListMap map[string]struct{}
	if blackList != nil {
		blackListMap = make(map[string]struct{})
		for _, s := range blackList {
			blackListMap[s] = struct{}{}
		}
	}
	return func(iFace string) bool {
		if len(blackListMap) == 0 {
			return true
		}
		_, ok := blackListMap[iFace]
		return !ok
	}
}

func (conn *Conn) reCreateAgent() error {
	conn.mu.Lock()
	defer conn.mu.Unlock()

	failedTimeout := 6 * time.Second
	var err error
	conn.agent, err = ice.NewAgent(&ice.AgentConfig{
		MulticastDNSMode: ice.MulticastDNSModeDisabled,
		NetworkTypes:     []ice.NetworkType{ice.NetworkTypeUDP4},
		Urls:             conn.config.StunTurn,
		CandidateTypes:   []ice.CandidateType{ice.CandidateTypeHost, ice.CandidateTypeServerReflexive, ice.CandidateTypeRelay},
		FailedTimeout:    &failedTimeout,
		InterfaceFilter:  interfaceFilter(conn.config.InterfaceBlackList),
	})
	if err != nil {
		return err
	}

	err = conn.agent.OnCandidate(conn.onICECandidate)
	if err != nil {
		return err
	}

	err = conn.agent.OnConnectionStateChange(conn.onICEConnectionStateChange)
	if err != nil {
		return err
	}

	err = conn.agent.OnSelectedCandidatePairChange(conn.onICESelectedCandidatePair)
	if err != nil {
		return err
	}

	return nil
}

func (conn *Conn) sendOffer() error {
	conn.mu.Lock()
	defer conn.mu.Unlock()

	localUFrag, localPwd, err := conn.agent.GetLocalUserCredentials()
	if err != nil {
		return err
	}
	err = conn.signalOffer(localUFrag, localPwd)
	if err != nil {
		return err
	}
	return nil
}

// onICECandidate is a callback attached to an ICE Agent to receive new local connection candidates
// and then signals them to the remote peer
func (conn *Conn) onICECandidate(candidate ice.Candidate) {
	if candidate != nil {
		//log.Debugf("discovered local candidate %s", candidate.String())
		go func() {
			err := conn.signalCandidate(candidate)
			if err != nil {
				fmt.Errorf("failed signaling candidate to the remote peer %s %s", conn.config.Key, err)
			}
		}()
	}
}

func (conn *Conn) onICESelectedCandidatePair(c1 ice.Candidate, c2 ice.Candidate) {
	fmt.Printf("selected candidate pair [local <-> remote] -> [%s <-> %s], peer %s", conn.config.Key,
		c1.String(), c2.String())
}

// onICEConnectionStateChange registers callback of an ICE Agent to track connection state
func (conn *Conn) onICEConnectionStateChange(state ice.ConnectionState) {
	fmt.Printf("peer %s ICE ConnectionState has changed to %s", conn.config.Key, state.String())
	if state == ice.ConnectionStateFailed || state == ice.ConnectionStateDisconnected {
		conn.notifyDisconnected()
	}
}
