package engine

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
	"github.com/Notch-Technologies/wizy/polymer"
	"github.com/Notch-Technologies/wizy/wislog"
	"github.com/pion/ice/v2"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

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

type Engine struct {
	client *grpc_client.GrpcClient
	stream negotiation.Negotiation_ConnectStreamClient

	peerConns map[string]*polymer.Conn

	STUNs []*ice.URL
	TURNs []*ice.URL

	config *EngineConfig

	syncMsgMux *sync.Mutex

	wislog *wislog.WisLog

	cancel context.CancelFunc

	// not used
	ctx context.Context
}

func NewEngine(
	log *wislog.WisLog,
	client *grpc_client.GrpcClient,
	stream negotiation.Negotiation_ConnectStreamClient,
	cancel context.CancelFunc,
	ctx context.Context,
) *Engine {
	return &Engine{
		client: client,
		stream: stream,

		STUNs: []*ice.URL{},
		TURNs: []*ice.URL{},

		syncMsgMux: &sync.Mutex{},

		wislog: log,
		cancel:     cancel,
		ctx: ctx,
	}
}

// TODO:
// 1. create Engine
// 2. send stream message
// 3. create management json
// 4. return to stun and sturn
// 5. send to stun and turn udp request
// 6. send to signal offer
// 7. connection peer to peer test
// 8. management peers and how save the peer connectivity state? maybe sync mutex??
func (e *Engine) Start(publicMachineKey string) {
	e.syncMsgMux.Lock()
	defer e.syncMsgMux.Unlock()

	e.receiveClient(publicMachineKey)
	e.syncClient(publicMachineKey)
}

func (e *Engine) receiveClient(machineKey string) {
	go func() {
		err := e.client.Receive(machineKey, func(msg *negotiation.Body) error {
			e.syncMsgMux.Lock()
			defer e.syncMsgMux.Unlock()

			conn := e.peerConns[msg.Key]
			if conn == nil {
				return fmt.Errorf("wrongly addressed message %s", msg.Key)
			}

			switch msg.GetType() {
			case negotiation.Body_OFFER:
				fmt.Println("Offer")
				fmt.Println(msg)
			case negotiation.Body_ANSWER:
				fmt.Println("Answer")
				fmt.Println(msg)
			case negotiation.Body_CANDIDATE:
				fmt.Println("Candidate")
				fmt.Println(msg)
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

func (e *Engine) syncClient(machineKey string) {
	go func() {
		err := e.client.Sync(machineKey, func(update *peer.SyncResponse) error {
			e.syncMsgMux.Lock()
			defer e.syncMsgMux.Unlock()

			err := e.StartConn(update.GetRemotePeers())
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
	fmt.Println("(3) start conn")
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

	fmt.Println("(4) remove peers")
	fmt.Println(remotePeers)

	// create connection remotePeers
	for _, p := range remotePeers {
		peerKey := p.GetWgPubKey()
		peerIPs := p.GetAllowedIps()
		fmt.Println("(5) get remote peers")
		fmt.Println(peerKey)
		fmt.Println(peerIPs)

		if _, ok := e.peerConns[peerKey]; !ok {
			conn, err := e.createPeerConn(peerKey, strings.Join(peerIPs, ","))
			if err != nil {
				return err
			}
			e.peerConns[peerKey] = conn

			go e.connWorker(conn, peerKey)
		}
	}

	return nil
}

func (e *Engine) createPeerConn(peerPubKey string, allowedIPs string) (*polymer.Conn, error) {
	var stunTurn []*ice.URL

	// store existing STUN, TURN
	stunTurn = append(stunTurn, e.STUNs...)
	stunTurn = append(stunTurn, e.TURNs...)

	// create blacklist
	interfaceBlacklist := make([]string, 0, len(e.config.IFaceBlackList))
	for k := range e.config.IFaceBlackList {
		interfaceBlacklist = append(interfaceBlacklist, k)
	}

	pc := polymer.ProxyConfig{
		RemoteKey:    peerPubKey,
		WgListenAddr: fmt.Sprintf("127.0.0.1:%d", e.config.WgPort),
		WgInterface:  e.config.WgIface,
		AllowedIps:   allowedIPs,
		PreSharedKey: e.config.PreSharedKey,
	}

	config := polymer.ConnConfig{
		Key:                peerPubKey,
		LocalKey:           e.config.WgPrivateKey.PublicKey().String(),
		StunTurn:           stunTurn,
		InterfaceBlackList: interfaceBlacklist,
		Timeout:            45000,
		ProxyConfig:        pc,
	}

	peerConn, err := polymer.NewConn(config)
	if err != nil {
		return nil, err
	}

	wgPubKey, err := wgtypes.ParseKey(peerPubKey)
	if err != nil {
		return nil, err
	}

	signalOffer := func(uFrag string, pwd string) error {
		return signalAuth(uFrag, pwd, e.config.WgPrivateKey, wgPubKey, e.client, false)
	}

	peerConn.SetSignalOffer(signalOffer)

	return peerConn, nil
}

func (e *Engine) connWorker(conn *polymer.Conn, peerKey string) {
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
	}
}

func (e *Engine) peerExists(peerKey string) bool {
	e.syncMsgMux.Lock()
	defer e.syncMsgMux.Unlock()
	_, ok := e.peerConns[peerKey]
	return ok
}

func signalAuth(uFrag string, pwd string, myKey wgtypes.Key, remoteKey wgtypes.Key, s *grpc_client.GrpcClient, isAnswer bool) error {

	var t negotiation.Body_Type
	if isAnswer {
		t = negotiation.Body_ANSWER
	} else {
		t = negotiation.Body_OFFER
	}

	err := s.Send(&negotiation.Body{Type: t})
	if err != nil {
		return err
	}

	return nil
}
