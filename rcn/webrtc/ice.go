package webrtc

// ice and provides webrtc functionalit
//

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/Notch-Technologies/dotshake/client/grpc"
	"github.com/Notch-Technologies/dotshake/dotengine/proxy"
	"github.com/Notch-Technologies/dotshake/dotengine/wonderwall"
	"github.com/Notch-Technologies/dotshake/dotlog"
	"github.com/Notch-Technologies/dotshake/iface"
	"github.com/Notch-Technologies/dotshake/rcn/puncher"
	"github.com/pion/ice/v2"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

type Ice struct {
	signalClient grpc.SignalClientImpl

	sigexec *SigExecuter

	iface *iface.Iface

	wireproxy *proxy.WireProxy

	wonderwallsock *wonderwall.WonderWallSock

	// channel to use when making a peer connection
	remoteOfferCh  chan Credentials
	remoteAnswerCh chan Credentials

	agent *ice.Agent

	stunTurn *puncher.StunTurnConfig

	disconnectFunc func() error
	connectFunc    func() error

	// remote
	remoteWgPubKey   string
	remoteIp         string
	remoteMachineKey string

	// local
	wgPubKey     string
	wgPrivKey    wgtypes.Key
	wgIface      string
	wgPort       int
	preSharedKey string
	tun          string
	ip           string
	cidr         string
	mk           string

	blackList []string

	mu      *sync.Mutex
	closeCh chan struct{}

	failedTimeout *time.Duration

	dotlog *dotlog.DotLog
}

func NewIce(
	signalClient grpc.SignalClientImpl,

	// remote
	remoteWgPubKey string,
	remoteip string,
	remoteMachineKey string,

	// local
	ip string,
	cidr string,
	tun string,
	wgPrivateKey wgtypes.Key,
	wgPort int,
	wgIface string,
	presharedKey string,
	mk string,

	// conn state func
	disconnect func() error,
	connect func() error,

	stunTurn *puncher.StunTurnConfig,
	blacklist []string,

	dotlog *dotlog.DotLog,
) *Ice {
	failedtimeout := time.Second * 5
	return &Ice{
		signalClient: signalClient,

		remoteOfferCh:  make(chan Credentials),
		remoteAnswerCh: make(chan Credentials),

		stunTurn: stunTurn,

		disconnectFunc: disconnect,
		connectFunc:    connect,

		remoteWgPubKey:   remoteWgPubKey,
		remoteIp:         remoteip,
		remoteMachineKey: remoteMachineKey,

		wgPubKey:     wgPrivateKey.PublicKey().String(),
		wgPrivKey:    wgPrivateKey,
		wgIface:      wgIface,
		wgPort:       wgPort,
		preSharedKey: presharedKey,
		tun:          tun,
		ip:           ip,
		cidr:         cidr,
		mk:           mk,

		blackList: blacklist,

		mu:      &sync.Mutex{},
		closeCh: make(chan struct{}),

		failedTimeout: &failedtimeout,

		dotlog: dotlog,
	}
}

// must be called before calling ConfigureGatherProcess
//
func (i *Ice) setup() (err error) {
	i.mu.Lock()
	defer i.mu.Unlock()

	// configure sigexe
	se := NewSigExecuter(i.signalClient, i.remoteMachineKey, i.mk, i.dotlog)
	i.sigexec = se

	// configure ice agent
	i.agent, err = ice.NewAgent(&ice.AgentConfig{
		MulticastDNSMode: ice.MulticastDNSModeDisabled,
		NetworkTypes:     []ice.NetworkType{ice.NetworkTypeUDP4},
		Urls:             i.stunTurn.GetStunTurnsURL(),
		CandidateTypes:   []ice.CandidateType{ice.CandidateTypeHost, ice.CandidateTypeServerReflexive, ice.CandidateTypeRelay},
		FailedTimeout:    i.failedTimeout,
		InterfaceFilter:  i.getBlackListWithInterfaceFilter(),
	})
	if err != nil {
		return err
	}

	// configure ice candidate functions
	err = i.agent.OnCandidate(i.sigexec.Candidate)
	if err != nil {
		return err
	}

	err = i.agent.OnConnectionStateChange(i.IceConnectionHasBeenChanged)
	if err != nil {
		return err
	}

	err = i.agent.OnSelectedCandidatePairChange(i.IceSelectedHasCandidatePairChanged)
	if err != nil {
		return err
	}

	// configure iface
	iface := iface.NewIface(i.tun, i.wgPrivKey.String(), i.ip, i.cidr, i.dotlog)

	// configure wire proxy
	wireproxy := proxy.NewWireProxy(
		iface,
		i.remoteWgPubKey,
		i.remoteIp,
		i.wgIface,
		fmt.Sprintf("127.0.0.1:%d", i.wgPort),
		i.preSharedKey,
		i.dotlog,
	)

	i.wireproxy = wireproxy

	i.wonderwallsock = wonderwall.NewWonderWallSock(nil, i.dotlog)

	return nil
}

// TODO: (shinta)
// more detailed handling is needed.
// by handling failures, we need to establish a connection path using DoubleNat? or
// Ether(call me エーテル) when a connection cannot be made.
func (i *Ice) IceConnectionHasBeenChanged(state ice.ConnectionState) {
	switch state {
	case ice.ConnectionStateNew: // ConnectionStateNew ICE agent is gathering addresses
		i.dotlog.Logger.Debugf("new connections collected, [%s]", state.String())
	case ice.ConnectionStateChecking: // ConnectionStateNew ICE agent is gathering addresses
		i.dotlog.Logger.Debugf("agent has been given local and remote candidates, and is attempting to find a match, [%s]", state.String())
	case ice.ConnectionStateConnected: // ConnectionStateConnected ICE agent has a pairing, but is still checking other pairs
		i.dotlog.Logger.Debugf("agent has a pairing, but is still checking other pairs, [%s]", state.String())
	case ice.ConnectionStateCompleted: // ConnectionStateConnected ICE agent has a pairing, but is still checking other pairs
		err := i.connectFunc()
		if err != nil {
			i.dotlog.Logger.Errorf("the agent connection was successful but I received an error in the function that updates the status to connect, [%s]", state.String())
		}
		i.dotlog.Logger.Debugf("successfully connected to agent, [%s]", state.String())
	case ice.ConnectionStateFailed: // ConnectionStateFailed ICE agent never could successfully connect
		err := i.disconnectFunc()
		if err != nil {
			i.dotlog.Logger.Errorf("agent connection failed, but failed to set the connection state to disconnect, [%s]", state.String())
		}
	case ice.ConnectionStateDisconnected: // ConnectionStateDisconnected ICE agent connected successfully, but has entered a failed state
		err := i.disconnectFunc()
		if err != nil {
			i.dotlog.Logger.Errorf("agent connected successfully, but has entered a failed state, [%s]", state.String())
		}
	case ice.ConnectionStateClosed: // ConnectionStateClosed ICE agent has finished and is no longer handling requests
		i.dotlog.Logger.Debugf("agent has finished and is no longer handling requests, [%s]", state.String())
	}
}

func (i *Ice) IceSelectedHasCandidatePairChanged(local ice.Candidate, remote ice.Candidate) {
	i.dotlog.Logger.Debugf("agent candidates were found, local:[%s] <-> remote:[%s]", local.Address(), remote.Address())
}

// be sure to read this function before using the Ice structures
//
func (i *Ice) ConfigureGatherProcess() error {
	err := i.setup()
	if err != nil {
		i.dotlog.Logger.Errorf("failed to configure gather process")
		return err
	}

	i.dotlog.Logger.Debugf("complete configure gather process")
	return nil
}

func (i *Ice) GetRemoteMachineKey() string {
	return i.remoteMachineKey
}

func (i *Ice) getBlackListWithInterfaceFilter() func(string) bool {
	var blackListMap map[string]struct{}
	if i.blackList != nil {
		blackListMap = make(map[string]struct{})
		for _, s := range i.blackList {
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

func (i *Ice) clean() error {
	i.dotlog.Logger.Debugf("starting clean ice agent process")
	i.mu.Lock()
	defer i.mu.Unlock()

	if i.agent == nil {
		i.dotlog.Logger.Errorf("the agent is nil and you are trying to clean it, please make sure that the agent is properly initialized")
		return errors.New("clean is called even though the agent is nil")
	}

	i.agent.Close()
	i.agent = nil

	i.disconnectFunc()

	return nil
}

func (i *Ice) getLocalUserIceAgentCredentials() (string, string, error) {
	uname, pwd, err := i.agent.GetLocalUserCredentials()
	if err != nil {
		return "", "", err
	}

	return uname, pwd, nil
}

// be sure to read ConfigureGatherProcess before calling this function
//
func (i *Ice) StartGatheringProcess() {
	defer func() {
		err := i.clean()
		if err != nil {
			i.dotlog.Logger.Errorf("failed to clean up gathering process, because %s", err.Error())
			return
		}
		i.dotlog.Logger.Debugf("clean ice agent")
	}()

	err := i.signalOffer()
	if err != nil {
		i.dotlog.Logger.Errorf("failed to clean up gathering process, because %s", err.Error())
	}

	err = i.waitForSignalingProcess()
	if err != nil {
		i.dotlog.Logger.Errorf("failed to clean up gathering process, because %s", err.Error())
	}
}

func (i *Ice) Close() {
	i.mu.Lock()
	defer i.mu.Unlock()
	i.closeCh <- struct{}{}
}

func (i *Ice) waitForSignalingProcess() error {
	for {
		var credentials Credentials
		select {
		case credentials = <-i.remoteAnswerCh:
			i.dotlog.Logger.Debugf("receive a channel for a remote offer")

			err := i.wonderwallsock.DialStartWonderWall(&wonderwall.StartWonderWallSock{
				Uname: credentials.UserName,
				Pwd:   credentials.Pwd,

				Agent:     i.agent,
				WireProxy: i.wireproxy,

				RemoteWgPubKey: i.remoteWgPubKey,
				WgPubKey:       i.wgPubKey,

				DisconnectFunc: i.disconnectFunc,
				ConnectFunc:    i.connectFunc,

				Dotlog: i.dotlog,
			})
			if err != nil {
				i.dotlog.Logger.Errorf("can not dial start wonderwall")
				return err
			}
		case credentials = <-i.remoteOfferCh:
			i.dotlog.Logger.Debugf("receive a channel for a remote answer")
			err := i.signalAnswer()
			if err != nil {
				return err
			}
		case <-i.closeCh:
			i.dotlog.Logger.Debugf("abort signal process wait, will called clean function called")

			// it is necessary to stop the wonderwall connection running on the dotengine
			// side when ice's closech is called
			//
			err := i.wonderwallsock.DialStopWonderWall(&wonderwall.CloseWonderWallSock{
				CloseCh: i.closeCh,
			})
			if err != nil {
				i.dotlog.Logger.Errorf("can not dial start wonderwall")
				return err
			}
			return errors.New("abort waitForSignalingProcess")
		}
	}
}

func (i *Ice) signalAnswer() error {
	i.mu.Lock()
	defer i.mu.Unlock()

	uname, pwd, err := i.getLocalUserIceAgentCredentials()
	if err != nil {
		return err
	}

	err = i.sigexec.Answer(uname, pwd)
	if err != nil {
		return err
	}

	i.dotlog.Logger.Debugf("answer has been sent to the signal server")

	return nil
}

func (i *Ice) signalOffer() error {
	i.mu.Lock()
	defer i.mu.Unlock()

	uname, pwd, err := i.getLocalUserIceAgentCredentials()
	if err != nil {
		return err
	}

	err = i.sigexec.Offer(uname, pwd)
	if err != nil {
		return err
	}

	i.dotlog.Logger.Debugf("offer has been sent to the signal server")

	return nil
}

func (i *Ice) SendRemoteOfferCh(uname, pwd string) {
	select {
	case i.remoteOfferCh <- *NewCredentials(uname, pwd):
	default:
	}
}

func (i *Ice) SendRemoteAnswerCh(uname, pwd string) {
	select {
	case i.remoteAnswerCh <- *NewCredentials(uname, pwd):
	default:
	}
}

func (i *Ice) SendRemoteCandidate(candidate ice.Candidate) {
	go func() {
		i.mu.Lock()
		defer i.mu.Unlock()

		if i.agent == nil {
			return
		}

		err := i.agent.AddRemoteCandidate(candidate)
		if err != nil {
			return
		}
	}()
}
