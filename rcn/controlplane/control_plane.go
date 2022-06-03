package controlplane

// this package is responsible for communication with the signal server
// it also has the structure of ice of the remote peer as a map key with the machine key of the remote peer
// when the communication with the signal server is performed and operations are performed on the peer, they will basically be performed here.
//

import (
	"errors"
	"strings"
	"sync"

	"github.com/Notch-Technologies/client-go/notch/dotshake/v1/machine"
	"github.com/Notch-Technologies/client-go/notch/dotshake/v1/negotiation"
	"github.com/Notch-Technologies/dotshake/client/grpc"
	"github.com/Notch-Technologies/dotshake/conf"
	"github.com/Notch-Technologies/dotshake/dotlog"
	"github.com/Notch-Technologies/dotshake/rcn/puncher"
	"github.com/Notch-Technologies/dotshake/rcn/webrtc"
	"github.com/Notch-Technologies/dotshake/wireguard"
	"github.com/pion/ice/v2"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

type ControlPlane struct {
	signalClient grpc.SignalClientImpl
	peerConns    map[string]*webrtc.Ice //  with ice structure per clientmachinekey
	mk           string
	clientConf   *conf.ClientConf
	stconf       *puncher.StunTurnConfig

	mu                  *sync.Mutex
	ch                  chan struct{}
	waitForRemoteConnCh chan *webrtc.Ice

	dotlog *dotlog.DotLog
}

func NewControlPlane(
	signalClient grpc.SignalClientImpl,
	mk string,
	clientConf *conf.ClientConf,
	ch chan struct{},
	dotlog *dotlog.DotLog,
) *ControlPlane {
	return &ControlPlane{
		signalClient: signalClient,
		peerConns:    make(map[string]*webrtc.Ice),
		mk:           mk,
		clientConf:   clientConf,

		mu:                  &sync.Mutex{},
		ch:                  ch,
		waitForRemoteConnCh: make(chan *webrtc.Ice),

		dotlog: dotlog,
	}
}

func (c *ControlPlane) parseStun(url string) (*ice.URL, error) {
	stun, err := ice.ParseURL(url)
	if err != nil {
		return nil, err
	}

	return stun, err
}

func (c *ControlPlane) parseTurn(url, uname, pw string) (*ice.URL, error) {
	turn, err := ice.ParseURL(url)
	if err != nil {
		return nil, err
	}
	turn.Username = uname
	turn.Password = pw

	return turn, err
}

// set stun turn url to use webrtc
// (shinta) be sure to call this function before using the ConnectSignalServer
//
func (c *ControlPlane) ConfigureStunTurnConf() error {
	conf, err := c.signalClient.GetStunTurnConfig()
	if err != nil {
		// TOOD: (shinta) retry
		return err
	}

	stun, err := c.parseStun(conf.RtcConfig.StunHost.Url)
	if err != nil {
		return err
	}

	turn, err := c.parseTurn(
		conf.RtcConfig.TurnHost.Url,
		conf.RtcConfig.TurnHost.Username,
		conf.RtcConfig.TurnHost.Password,
	)
	if err != nil {
		return err
	}

	stcof := puncher.NewStunTurnConfig(stun, turn)

	c.stconf = stcof

	return nil
}

func (c *ControlPlane) receiveSignalingProcess(
	msgType *negotiation.NegotiationType,
	peer *webrtc.Ice,
	uname string,
	pwd string,
	candidate string,
) error {
	switch msgType {
	case negotiation.NegotiationType_ANSWER.Enum():
		peer.SendRemoteAnswerCh(uname, pwd)
	case negotiation.NegotiationType_OFFER.Enum():
		peer.SendRemoteOfferCh(uname, pwd)
	case negotiation.NegotiationType_CANDIDATE.Enum():
		candidate, err := ice.UnmarshalCandidate(candidate)
		if err != nil {
			c.dotlog.Logger.Errorf("can not unmarshal candidate => [%s]", candidate)
			return err
		}
		peer.SendRemoteCandidate(candidate)
	}

	return nil
}

// through StartConnect, the results of the execution of functions such as
// candidate required for udp hole punching are received from the dotengine side
//
func (c *ControlPlane) ConnectSignalServer() {
	go func() {
		err := c.signalClient.StartConnect(c.mk, func(res *negotiation.NegotiationResponse) error {
			c.mu.Lock()
			defer c.mu.Unlock()

			dstPeer := res.GetDstPeerMachineKey()
			if dstPeer == "" {
				c.dotlog.Logger.Errorf("empty dst peer machine key")
				return errors.New("empty dst peer machine key")
			}

			peer := c.peerConns[res.GetDstPeerMachineKey()]
			if peer == nil {
				c.dotlog.Logger.Errorf("empty remote peer connection, dst remote peer machine key is [%s]", res.GetDstPeerMachineKey())
				return errors.New("empty remote peer connection")
			}

			err := c.receiveSignalingProcess(res.GetType().Enum(), peer, res.GetUFlag(), res.GetPwd(), res.GetCandidate())
			if err != nil {
				return err
			}

			return nil
		})
		if err != nil {
			close(c.ch)
			return
		}
	}()
	c.signalClient.WaitStartConnect()
}

// keep the latest state of Peers received from the server
//
func (c *ControlPlane) SyncRemotePeerConnecting(remotePeers []*machine.RemotePeer) error {
	remotePeerMap := make(map[string]struct{})
	for _, p := range remotePeers {
		remotePeerMap[p.GetRemoteClientMachineKey()] = struct{}{}
	}

	unnecessary := []string{}
	for p := range c.peerConns {
		if _, ok := remotePeerMap[p]; !ok {
			unnecessary = append(unnecessary, p)
		}
	}

	for _, p := range unnecessary {
		conn, exists := c.peerConns[p]
		if exists {
			delete(c.peerConns, p)
			conn.Close()
		}
		c.dotlog.Logger.Errorf("there are no peers, even though there should be. machine key => [%s]", p)
	}

	if len(unnecessary) == 0 {
		c.dotlog.Logger.Debugf("completed peer delete in control plane, but it was nil")
		return nil
	}

	c.dotlog.Logger.Debugf("completed peer delete in signal control plane => %v", unnecessary)
	return nil
}

func (c *ControlPlane) configureIce(peer *machine.RemotePeer, ip, cidr string) (*webrtc.Ice, error) {
	k, err := wgtypes.ParseKey(c.clientConf.WgPrivateKey)
	if err != nil {
		return nil, err
	}

	var pk string
	if c.clientConf.PreSharedKey != "" {
		k, err := wgtypes.ParseKey(c.clientConf.PreSharedKey)
		if err != nil {
			return nil, err
		}
		pk = k.String()
	}

	allowip := strings.Join(peer.GetAllowedIPs(), ",")
	c.dotlog.Logger.Debugf("get allowe ip => [%s]", allowip)
	i := webrtc.NewIce(
		c.signalClient,
		peer.RemoteWgPubKey,
		ip,
		peer.GetRemoteClientMachineKey(),
		allowip,
		cidr,
		c.clientConf.TunName,
		k,
		wireguard.WgPort,
		c.clientConf.TunName,
		pk,
		c.mk,
		c.signalClient.DisConnected,
		c.signalClient.DisConnected,
		c.stconf,
		c.clientConf.BlackList,
		c.dotlog,
	)

	return i, nil
}

func (c *ControlPlane) SetupRemotePeerConn(connPeers []*machine.RemotePeer, ip, cidr string) error {
	for _, p := range connPeers {
		c.mu.Lock()
		defer c.mu.Unlock()
		_, ok := c.peerConns[p.GetRemoteClientMachineKey()]
		if !ok {
			i, err := c.configureIce(p, ip, cidr)
			if err != nil {
				return err
			}
			c.peerConns[p.GetRemoteClientMachineKey()] = i
			c.waitForRemoteConnCh <- i
		}
	}
	return nil
}

func (c *ControlPlane) isExistPeer(remoteMachineKey string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	_, exist := c.peerConns[remoteMachineKey]
	return exist
}

// function to wait until the channel is sent from SetupRemotePeerConn to waitForRemoteConnCh
//
func (c *ControlPlane) WaitForRemoteConn() {
	for {
		select {
		case ice := <-c.waitForRemoteConnCh:
			c.dotlog.Logger.Debugf("credential to standby for remote channels => [%v]", ice)
			if !c.signalClient.IsReady() || !c.isExistPeer(ice.GetRemoteMachineKey()) {
				c.dotlog.Logger.Debugf("signal client is not available, execute loop. applicable remote peer => [%s]", ice.GetRemoteMachineKey())
				continue
			}
			c.dotlog.Logger.Debugf("starting gathering process  => [%s]", ice.GetRemoteMachineKey())

			ice.ConfigureGatherProcess()
			go ice.StartGatheringProcess()
		}
	}
}
