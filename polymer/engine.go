package polymer

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"

	grpc_client "github.com/Notch-Technologies/wizy/cmd/server/grpc_client"
	"github.com/Notch-Technologies/wizy/cmd/server/pb/negotiation"
	"github.com/Notch-Technologies/wizy/cmd/server/pb/peer"
	"github.com/Notch-Technologies/wizy/cmd/wissy/client"
	"github.com/Notch-Technologies/wizy/wizlog"
	"github.com/pion/ice/v2"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

const WgPort = 51820

type EngineConfig struct {
	WgPort  int
	WgIface string
	// WgAddr is a Wireguard local address (Wiretrustee Network IP)
	WgAddr string
	// WgPrivateKey is a Wireguard private key of our peer (it MUST never leave the machine)
	WgPrivateKey wgtypes.Key
	// IFaceBlackList is a list of network interfaces to ignore when discovering connection candidates (ICE related)
	IFaceBlackList map[string]struct{}

	PreSharedKey *wgtypes.Key
}

func NewEngineConfig(key wgtypes.Key, config *client.Config, wgAddr string) *EngineConfig {
	iFaceBlackList := make(map[string]struct{})
	for i := 0; i < len(config.IfaceBlackList); i += 2 {
		iFaceBlackList[config.IfaceBlackList[i]] = struct{}{}
	}

	return &EngineConfig{
		WgIface:        config.TUNName,
		WgAddr:         wgAddr,
		IFaceBlackList: iFaceBlackList,
		WgPrivateKey:   key,
		WgPort:         WgPort,
	}
}

type Engine struct {
	client *grpc_client.GrpcClient
	// stream negotiation.Negotiation_ConnectStreamClient

	peerConns map[string]*Conn

	STUNs []*ice.URL
	TURNs []*ice.URL

	config *EngineConfig

	syncMsgMux *sync.Mutex

	wizlog *wizlog.WizLog

	cancel context.CancelFunc

	// not used
	ctx context.Context

	machineKey string

	wgPrivateKey wgtypes.Key
}

func NewEngine(
	log *wizlog.WizLog,
	client *grpc_client.GrpcClient,
	cancel context.CancelFunc,
	ctx context.Context,
	engineConfig *EngineConfig,
	machineKey string,
	wgPrivateKey wgtypes.Key,
) *Engine {
	return &Engine{
		client: client,

		peerConns: map[string]*Conn{},

		STUNs: []*ice.URL{},
		TURNs: []*ice.URL{},

		config: engineConfig,

		syncMsgMux: &sync.Mutex{},

		wizlog: log,
		cancel: cancel,
		ctx:    ctx,

		machineKey: machineKey,

		wgPrivateKey: wgPrivateKey,
	}
}

func (e *Engine) Start() {
	e.syncMsgMux.Lock()
	defer e.syncMsgMux.Unlock()

	// signal
	e.receiveClient()

	// management
	e.syncClient()
}

func (e *Engine) receiveClient() {
	go func() {
		err := e.client.Receive(e.wgPrivateKey.PublicKey().String(), func(msg *negotiation.Body) error {
			e.syncMsgMux.Lock()
			defer e.syncMsgMux.Unlock()

			fmt.Println("** Recieve Client, peerConns list **")
			fmt.Println(e.peerConns)
			conn := e.peerConns[msg.Key]
			if conn == nil {
				return fmt.Errorf("wrongly addressed message %s", msg.Key)
			}

			switch msg.GetType() {
			case negotiation.Body_OFFER:
				fmt.Println("** Offer is Coming **")
				conn.RemoteOffer(IceCredentials{
					UFrag: msg.UFlag,
					Pwd:   msg.Pwd,
				})
			case negotiation.Body_ANSWER:
				fmt.Println("** Answer is Coming **")
				conn.RemoteAnswer(IceCredentials{
					UFrag: msg.UFlag,
					Pwd:   msg.Pwd,
				})
			case negotiation.Body_CANDIDATE:
				fmt.Println("** Candidate is Coming **")
				candidate, err := ice.UnmarshalCandidate(msg.Payload)
				if err != nil {
					fmt.Println("failed parse ice candidate")
					return err
				}
				conn.OnRemoteCandidate(candidate)
			}
			return nil
		})
		if err != nil {
			e.cancel()
			return
		}
	}()

	e.client.WaitStreamConnected()
}

func (e *Engine) updateTurns() error {
	var newTurns []*ice.URL
	url, err := ice.ParseURL("turn:www.notchturn.net:3478")
	if err != nil {
		return err
	}
	url.Username = "root"
	url.Password = "password"
	newTurns = append(newTurns, url)
	e.TURNs = newTurns
	return nil
}

func (e *Engine) updateStuns() error {
	var newStuns []*ice.URL
	url, err := ice.ParseURL("stun:www.notchturn.net:3478")
	if err != nil {
		return err
	}
	// url.Username = "root"
	// url.Password = "password"
	newStuns = append(newStuns, url)
	e.STUNs = append(newStuns, url)
	return nil
}

func (e *Engine) syncClient() {
	go func() {
		err := e.client.Sync(e.machineKey, func(update *peer.SyncResponse) error {
			e.syncMsgMux.Lock()
			defer e.syncMsgMux.Unlock()

			err := e.updateTurns()
			if err != nil {
				return err
			}

			err = e.updateStuns()
			if err != nil {
				return err
			}

			err = e.StartConn(update.GetRemotePeers())
			if err != nil {
				return err
			}

			return nil
		})
		if err != nil {
			fmt.Println("stopping recive management server")
			e.cancel()
			return
		}
		fmt.Println("stopped receiving updates from Management Service")
	}()
	fmt.Println("connecting management server")
}

func (e *Engine) removePeers(peers []string) error {
	for _, p := range peers {
		err := e.removePeer(p)
		if err != nil {
			return err
		}
		fmt.Printf("removed peer %s\n", p)
	}
	return nil
}

func (e *Engine) removePeer(peerKey string) error {
	fmt.Printf("removing peer from engine %s\n", peerKey)
	conn, exists := e.peerConns[peerKey]
	if exists {
		delete(e.peerConns, peerKey)
		return conn.Close()
	}
	return nil
}

// starting connection
func (e *Engine) StartConn(remotePeers []*peer.RemotePeer) error {
	// remove old out peers
	remotePeerMap := make(map[string]struct{})
	for _, p := range remotePeers {
		remotePeerMap[p.GetClientMachineKey()] = struct{}{}
	}

	toRemove := []string{}
	for p := range e.peerConns {
		if _, ok := remotePeerMap[p]; !ok {
			toRemove = append(toRemove, p)
		}
	}

	err := e.removePeers(toRemove)
	if err != nil {
		return err
	}

	// create connection remotePeers
	for _, p := range remotePeers {
		peerKey := p.GetWgPubKey()
		peerIPs := p.GetAllowedIps()

		if _, ok := e.peerConns[peerKey]; !ok {
			conn, err := e.createPeerConn(peerKey, strings.Join(peerIPs, ","))
			if err != nil {
				fmt.Println("create peer conn error")
				return err
			}

			e.peerConns[peerKey] = conn

			// setuzoku sarerumadeha kokoga loop
			go e.connWorker(conn, peerKey)
		}
	}

	return nil
}

func (e *Engine) createPeerConn(peerPubKey string, allowedIPs string) (*Conn, error) {
	var stunTurn []*ice.URL

	// store existing STUN, TURN
	stunTurn = append(stunTurn, e.STUNs...)
	stunTurn = append(stunTurn, e.TURNs...)

	// create blacklist
	interfaceBlacklist := make([]string, 0, len(e.config.IFaceBlackList))
	for k := range e.config.IFaceBlackList {
		interfaceBlacklist = append(interfaceBlacklist, k)
	}

	pc := ProxyConfig{
		RemoteKey:    peerPubKey,
		WgListenAddr: fmt.Sprintf("127.0.0.1:%d", e.config.WgPort),
		WgInterface:  e.config.WgIface,
		AllowedIps:   allowedIPs,
		PreSharedKey: e.config.PreSharedKey,
	}

	const PeerConnectionTimeoutMax = 45000 //ms
	const PeerConnectionTimeoutMin = 30000 //ms
	timeout := time.Duration(rand.Intn(PeerConnectionTimeoutMax-PeerConnectionTimeoutMin)+PeerConnectionTimeoutMin) * time.Millisecond
	config := ConnConfig{
		Key:                peerPubKey,
		LocalKey:           e.config.WgPrivateKey.PublicKey().String(),
		StunTurn:           stunTurn,
		InterfaceBlackList: interfaceBlacklist,
		Timeout:            timeout,
		ProxyConfig:        pc,
	}

	peerConn, err := NewConn(config)
	if err != nil {
		return nil, err
	}

	wgPubKey, err := wgtypes.ParseKey(peerPubKey)
	if err != nil {
		return nil, err
	}

	signalOffer := func(uFrag string, pwd string) error {
		return signalAuth(uFrag, pwd, e.config.WgPrivateKey, wgPubKey, e.machineKey, e.client, false)
	}

	signalAnswer := func(uFrag string, pwd string) error {
		return signalAuth(uFrag, pwd, e.config.WgPrivateKey, wgPubKey, e.machineKey, e.client, true)
	}

	signalCandidate := func(candidate ice.Candidate) error {
		return signalCandidate(candidate, e.config.WgPrivateKey, wgPubKey, e.machineKey, e.client)
	}

	peerConn.SetSignalOffer(signalOffer)
	peerConn.SetSignalAnswer(signalAnswer)
	peerConn.SetSignalCandidate(signalCandidate)

	return peerConn, nil
}

func (e *Engine) connWorker(conn *Conn, peerKey string) {
	for {
		min := 500
		max := 2000
		time.Sleep(time.Duration(rand.Intn(max-min)+min) * time.Millisecond)

		if !e.peerExists(peerKey) {
			fmt.Printf("peer %s doesn't exist anymore, won't retry connection", peerKey)
			return
		}

		if !e.client.Ready() {
			fmt.Printf("signal client isn't ready, skipping connection attempt %s", peerKey)
			continue
		}

		fmt.Println("start conn worker")

		// MEMO: リモートピアのconnに対してDialができていない
		err := conn.Open()
		if err != nil {
			fmt.Printf("connection to failed: %v\n", err)
		}
	}
}

func (e *Engine) peerExists(peerKey string) bool {
	e.syncMsgMux.Lock()
	defer e.syncMsgMux.Unlock()
	_, ok := e.peerConns[peerKey]
	return ok
}

func signalAuth(uFrag string, pwd string, myKey wgtypes.Key, remoteKey wgtypes.Key, clientMachineKey string, s *grpc_client.GrpcClient, isAnswer bool) error {
	var t negotiation.Body_Type
	if isAnswer {
		t = negotiation.Body_ANSWER
	} else {
		t = negotiation.Body_OFFER
	}

	err := s.Send(&negotiation.Body{
		UFlag:            uFrag,
		Pwd:              pwd,
		Key:              myKey.PublicKey().String(),
		Remotekey:        remoteKey.String(),
		ClientMachineKey: clientMachineKey,
		Type:             t,
	})
	if err != nil {
		fmt.Println("can not send negotiation send")
		fmt.Println(err)
		return err
	}

	return nil
}

func signalCandidate(candidate ice.Candidate, myKey wgtypes.Key, remoteKey wgtypes.Key, clientMachineKey string, s *grpc_client.GrpcClient) error {
	err := s.Send(&negotiation.Body{
		Key:              myKey.PublicKey().String(),
		Remotekey:        remoteKey.String(),
		ClientMachineKey: clientMachineKey,
		Type:             negotiation.Body_CANDIDATE,
		Payload:          candidate.Marshal(),
	})
	if err != nil {
		fmt.Errorf("failed signaling candidate to the remote peer %s %s", remoteKey.String(), err)
		//todo ??
		return err
	}

	return nil
}
