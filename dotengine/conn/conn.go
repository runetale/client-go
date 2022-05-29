package conn

// import (
// 	"context"
// 	"fmt"
// 	"net"
// 	"sync"
// 	"time"

// 	"github.com/Notch-Technologies/dotshake/iface"
// 	"github.com/pion/ice/v2"
// )

// // ConnConfig is a peer Connection configuration
// type ConnConfig struct {
// 	// Key is a public key of a remote peer
// 	RemotePeerPubKey string
// 	// LocalKey is a public key of a local peer
// 	LocalKey string

// 	// StunTurn is a list of STUN and TURN URLs
// 	StunTurn []*ice.URL

// 	// InterfaceBlackList is a list of machine interfaces that should be filtered out by ICE Candidate gathering
// 	// (e.g. if eth0 is in the list, host candidate of this interface won't be used)
// 	InterfaceBlackList []string

// 	Timeout time.Duration

// 	ProxyConfig *ProxyConfig
// }

// func NewConnConfig(
// 	remotePeerPubKey string,
// 	localKey string,
// 	stunTurn []*ice.URL,
// 	interfaceBlacklist []string,
// 	timeout time.Duration,
// 	proxyConfig *ProxyConfig,
// ) *ConnConfig {
// 	return &ConnConfig{
// 		RemotePeerPubKey:   remotePeerPubKey,
// 		LocalKey:           localKey,
// 		StunTurn:           stunTurn,
// 		InterfaceBlackList: interfaceBlacklist,
// 		Timeout:            timeout,
// 		ProxyConfig:        proxyConfig,
// 	}

// }

// // IceCredentials ICE protocol credentials struct
// type IceCredentials struct {
// 	UFrag string
// 	Pwd   string
// }

// type Conn struct {
// 	config *ConnConfig
// 	iface  *iface.Iface

// 	mu sync.Mutex

// 	// signalCandidate is a handler function to signal remote peer about local connection candidate
// 	signalCandidate func(candidate ice.Candidate) error
// 	// signalOffer is a handler function to signal remote peer our connection offer (credentials)
// 	signalOffer  func(uFrag string, pwd string) error
// 	signalAnswer func(uFrag string, pwd string) error

// 	// remoteOffersCh is a channel used to wait for remote credentials to proceed with the connection
// 	remoteOffersCh chan IceCredentials
// 	// remoteAnswerCh is a channel used to wait for remote credentials answer (confirmation of our offer) to proceed with the connection
// 	remoteAnswerCh     chan IceCredentials
// 	closeCh            chan struct{}
// 	ctx                context.Context
// 	notifyDisconnected context.CancelFunc

// 	agent  *ice.Agent
// 	status ConnStatus

// 	proxy Proxyer
// }

// func NewConn(config *ConnConfig, iface *iface.Iface) (*Conn, error) {
// 	return &Conn{
// 		config: config,
// 		iface:  iface,

// 		mu: sync.Mutex{},

// 		remoteOffersCh: make(chan IceCredentials),
// 		remoteAnswerCh: make(chan IceCredentials),

// 		closeCh: make(chan struct{}),

// 		status: StatusDisconnected,
// 	}, nil
// }

// func (conn *Conn) Close() error {
// 	conn.mu.Lock()
// 	defer conn.mu.Unlock()
// 	select {
// 	case conn.closeCh <- struct{}{}:
// 	default:
// 		fmt.Printf("closing not started coonection %s\n", conn.config.RemotePeerPubKey)
// 	}
// 	return nil
// }

// func (conn *Conn) SetSignalOffer(handler func(uFrag string, pwd string) error) {
// 	conn.signalOffer = handler
// }

// func (conn *Conn) SetSignalAnswer(handler func(uFrag string, pwd string) error) {
// 	conn.signalAnswer = handler
// }

// func (conn *Conn) SetSignalCandidate(handler func(candidate ice.Candidate) error) {
// 	conn.signalCandidate = handler
// }

// func (conn *Conn) cleanup() error {
// 	fmt.Printf("trying to cleanup %s\n", conn.config.RemotePeerPubKey)
// 	conn.mu.Lock()
// 	defer conn.mu.Unlock()

// 	if conn.agent != nil {
// 		err := conn.agent.Close()
// 		if err != nil {
// 			return err
// 		}
// 		conn.agent = nil
// 	}

// 	if conn.proxy != nil {
// 		err := conn.proxy.Close()
// 		if err != nil {
// 			return err
// 		}
// 		conn.proxy = nil
// 	}

// 	if conn.notifyDisconnected != nil {
// 		conn.notifyDisconnected()
// 		conn.notifyDisconnected = nil
// 	}

// 	conn.status = StatusDisconnected

// 	fmt.Printf("cleaned up connection to peer %s\n", conn.config.RemotePeerPubKey)

// 	return nil
// }

// func (conn *Conn) Open() error {
// 	defer func() {
// 		err := conn.cleanup()
// 		if err != nil {
// 			fmt.Printf("error while cleaning up peer connection %s: %v\n", conn.config.RemotePeerPubKey, err)
// 			return
// 		}
// 	}()

// 	err := conn.reCreateAgent()
// 	if err != nil {
// 		return err
// 	}

// 	err = conn.sendOffer()
// 	if err != nil {
// 		return err
// 	}

// 	fmt.Printf("connection offer sent to peer %s, waiting for the confirmation\n", conn.config.RemotePeerPubKey)

// 	var remoteCredentials IceCredentials

// 	select {
// 	case remoteCredentials = <-conn.remoteOffersCh:
// 		fmt.Println("** send answer **")
// 		err = conn.sendAnswer()
// 		if err != nil {
// 			return err
// 		}
// 	case remoteCredentials = <-conn.remoteAnswerCh:
// 	case <-time.After(conn.config.Timeout):
// 		fmt.Println("** timeout from Open() **")
// 		return NewConnectionTimeoutError(conn.config.RemotePeerPubKey, conn.config.Timeout)
// 	case <-conn.closeCh:
// 		fmt.Println("** closeCh from Open() **")
// 		return NewConnectionClosedError(conn.config.RemotePeerPubKey)
// 	}

// 	conn.mu.Lock()
// 	conn.status = StatusConnected
// 	conn.ctx, conn.notifyDisconnected = context.WithCancel(context.Background())
// 	defer conn.notifyDisconnected()
// 	conn.mu.Unlock()

// 	err = conn.agent.GatherCandidates()
// 	if err != nil {
// 		fmt.Println("[ERR] gather candidates error")
// 		return err
// 	}

// 	var remoteConn *ice.Conn
// 	isControlling := conn.config.LocalKey > conn.config.RemotePeerPubKey
// 	if isControlling {
// 		remoteConn, err = conn.agent.Dial(conn.ctx, remoteCredentials.UFrag, remoteCredentials.Pwd)
// 	} else {
// 		remoteConn, err = conn.agent.Accept(conn.ctx, remoteCredentials.UFrag, remoteCredentials.Pwd)
// 	}

// 	if err != nil {
// 		fmt.Println("[ERR] Dial or Accept Error")
// 		return err
// 	}

// 	fmt.Println("** Start proxy **")
// 	err = conn.startProxy(remoteConn)
// 	if err != nil {
// 		return err
// 	}

// 	fmt.Println("** Connected **")
// 	fmt.Printf("connected to peer %s [laddr <-> raddr] [%s <-> %s]\n", conn.config.RemotePeerPubKey, remoteConn.LocalAddr().String(), remoteConn.RemoteAddr().String())

// 	// wait until connection disconnected or has been closed externally (upper layer, e.g. engine)
// 	select {
// 	case <-conn.closeCh:
// 		// closed externally
// 		return NewConnectionClosedError(conn.config.RemotePeerPubKey)
// 	case <-conn.ctx.Done():
// 		// disconnected from the remote peer
// 		return NewConnectionDisconnectedError(conn.config.RemotePeerPubKey)
// 	}
// }

// func (conn *Conn) startProxy(remoteConn net.Conn) error {
// 	conn.mu.Lock()
// 	defer conn.mu.Unlock()

// 	conn.proxy = NewProxy(conn.config.ProxyConfig, conn.iface)

// 	err := conn.proxy.Start(remoteConn)
// 	if err != nil {
// 		return err
// 	}

// 	conn.status = StatusConnected

// 	return nil
// }

// func (conn *Conn) sendAnswer() error {
// 	conn.mu.Lock()
// 	defer conn.mu.Unlock()

// 	localUFrag, localPwd, err := conn.agent.GetLocalUserCredentials()
// 	if err != nil {
// 		return err
// 	}

// 	fmt.Printf("sending asnwer to %s\n", conn.config.RemotePeerPubKey)
// 	err = conn.signalAnswer(localUFrag, localPwd)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func interfaceFilter(blackList []string) func(string) bool {
// 	var blackListMap map[string]struct{}
// 	if blackList != nil {
// 		blackListMap = make(map[string]struct{})
// 		for _, s := range blackList {
// 			blackListMap[s] = struct{}{}
// 		}
// 	}
// 	return func(iFace string) bool {
// 		if len(blackListMap) == 0 {
// 			return true
// 		}
// 		_, ok := blackListMap[iFace]
// 		return !ok
// 	}
// }

// func (conn *Conn) reCreateAgent() error {
// 	conn.mu.Lock()
// 	defer conn.mu.Unlock()

// 	failedTimeout := 6 * time.Second
// 	var err error
// 	conn.agent, err = ice.NewAgent(&ice.AgentConfig{
// 		MulticastDNSMode: ice.MulticastDNSModeDisabled,
// 		NetworkTypes:     []ice.NetworkType{ice.NetworkTypeUDP4},
// 		Urls:             conn.config.StunTurn,
// 		CandidateTypes:   []ice.CandidateType{ice.CandidateTypeHost, ice.CandidateTypeServerReflexive, ice.CandidateTypeRelay},
// 		FailedTimeout:    &failedTimeout,
// 		InterfaceFilter:  interfaceFilter(conn.config.InterfaceBlackList),
// 	})
// 	if err != nil {
// 		return err
// 	}

// 	err = conn.agent.OnCandidate(conn.onICECandidate)
// 	if err != nil {
// 		return err
// 	}

// 	err = conn.agent.OnConnectionStateChange(conn.onICEConnectionStateChange)
// 	if err != nil {
// 		return err
// 	}

// 	err = conn.agent.OnSelectedCandidatePairChange(conn.onICESelectedCandidatePair)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (conn *Conn) sendOffer() error {
// 	conn.mu.Lock()
// 	defer conn.mu.Unlock()

// 	localUFrag, localPwd, err := conn.agent.GetLocalUserCredentials()
// 	if err != nil {
// 		fmt.Println("error get local user credentials")
// 		return err
// 	}
// 	err = conn.signalOffer(localUFrag, localPwd)
// 	if err != nil {
// 		fmt.Println("can not send signal offer")
// 		return err
// 	}
// 	return nil
// }

// func (conn *Conn) onICECandidate(candidate ice.Candidate) {
// 	if candidate != nil {
// 		go func() {
// 			err := conn.signalCandidate(candidate)
// 			if err != nil {
// 				fmt.Printf("failed signaling candidate to the remote peer %s %s\n", conn.config.RemotePeerPubKey, err)
// 			}
// 		}()
// 	}
// }

// func (conn *Conn) onICESelectedCandidatePair(c1 ice.Candidate, c2 ice.Candidate) {
// 	fmt.Printf("selected candidate pair [local <-> remote] -> [%s <-> %s], peer %s\n", conn.config.RemotePeerPubKey,
// 		c1.String(), c2.String())
// }

// // onICEConnectionStateChange registers callback of an ICE Agent to track connection state
// func (conn *Conn) onICEConnectionStateChange(state ice.ConnectionState) {
// 	fmt.Printf("** peer %s ICE ConnectionState has changed to %s **\n", conn.config.RemotePeerPubKey, state.String())
// 	if state == ice.ConnectionStateFailed || state == ice.ConnectionStateDisconnected {
// 		fmt.Println("** Failed or ConnectionStateDisconnected onICEConnectionStateChange **")
// 		conn.notifyDisconnected()
// 	}
// }

// func (conn *Conn) RemoteOffer(offer IceCredentials) {
// 	fmt.Printf("OnRemoteOffer from peer %s on status %s\n", conn.config.RemotePeerPubKey, conn.status.String())
// 	select {
// 	case conn.remoteOffersCh <- offer:
// 	default:
// 		fmt.Printf("OnRemoteOffer skipping message from peer %s on status %s because is not ready\n", conn.config.RemotePeerPubKey, conn.status.String())
// 	}
// }

// func (conn *Conn) RemoteAnswer(answer IceCredentials) {
// 	select {
// 	case conn.remoteAnswerCh <- answer:
// 	default:
// 		fmt.Printf("OnRemoteAnswer skipping message from peer %s on status %s because is not ready\n", conn.config.RemotePeerPubKey, conn.status.String())
// 	}
// }

// // onICECandidate ga yobareru
// func (conn *Conn) OnRemoteCandidate(candidate ice.Candidate) {
// 	fmt.Printf("OnRemoteCandidate from peer %s -> %s\n", conn.config.RemotePeerPubKey, candidate.String())
// 	go func() {
// 		conn.mu.Lock()
// 		defer conn.mu.Unlock()

// 		if conn.agent == nil {
// 			return
// 		}

// 		err := conn.agent.AddRemoteCandidate(candidate)
// 		if err != nil {
// 			fmt.Printf("error while handling remote candidate from peer %s\n", conn.config.RemotePeerPubKey)
// 			return
// 		}
// 	}()
// }
