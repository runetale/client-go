package polymer

import (
	"sync"

	"github.com/Notch-Technologies/client-go/notch/dotshake/v1/negotiation"
	"github.com/Notch-Technologies/dotshake/client/grpc"
	"github.com/Notch-Technologies/dotshake/dotlog"
	"github.com/Notch-Technologies/dotshake/polymer/puncher"
	"github.com/Notch-Technologies/dotshake/unixsock"
	"github.com/pion/ice/v2"
)

type Polymer struct {
	signalClient grpc.SignalClientImpl

	mk string

	stconf *puncher.StunTurnConfig

	mu *sync.Mutex
	ch chan struct{}

	dotlog *dotlog.DotLog
}

func NewPolymer(
	signalClient grpc.SignalClientImpl,
	mk string,
	ch chan struct{},
	mu *sync.Mutex,
	dotlog *dotlog.DotLog,
) *Polymer {
	return &Polymer{
		signalClient: signalClient,

		mk: mk,

		mu: mu,
		ch: ch,

		dotlog: dotlog,
	}
}

func (p *Polymer) parseStun(url string) (*ice.URL, error) {
	stun, err := ice.ParseURL(url)
	if err != nil {
		return nil, err
	}

	return stun, err
}

func (p *Polymer) parseTurn(url, uname, pw string) (*ice.URL, error) {
	turn, err := ice.ParseURL(url)
	if err != nil {
		return nil, err
	}
	turn.Username = uname
	turn.Password = pw

	return turn, err
}

// configure webrtc
// (shinta) be sure to call this function before using the polymer structure.
//
func (p *Polymer) configurePuncher() error {
	conf, err := p.signalClient.GetStunTurnConfig()
	if err != nil {
		// TOOD: (shintard) retry
		return err
	}

	stun, err := p.parseStun(conf.RtcConfig.StunHost.Url)
	if err != nil {
		return err
	}

	turn, err := p.parseTurn(
		conf.RtcConfig.TurnHost.Url,
		conf.RtcConfig.TurnHost.Username,
		conf.RtcConfig.TurnHost.Password,
	)
	if err != nil {
		return err
	}

	stcof := puncher.NewStunTurnConfig(stun, turn)

	p.stconf = stcof

	return nil
}

// connecting signal server
// must be connected
//
func (p *Polymer) startConnDotSignal() {
	go func() {
		err := p.signalClient.StartConnect(p.mk, func(msg *negotiation.NegotiationResponse) error {
			p.dotlog.Logger.Debugf("connect to signal server")
			p.mu.Lock()
			defer p.mu.Unlock()

			return nil
		})
		if err != nil {
			close(p.ch)
			return
		}
	}()
	p.signalClient.WaitStartConnect()
}

func (p *Polymer) connectSock() {
	sock := unixsock.NewPolyerSock(p.dotlog, p.ch)
	go func() {
		err := sock.Connect()
		if err != nil {
			p.dotlog.Logger.Errorf("failed to connect polymer sock. %s", err.Error())
		}
		p.dotlog.Logger.Debugf("connection with sock connect has been disconnected")
	}()
}

func (p *Polymer) Start() {
	err := p.configurePuncher()
	if err != nil {
		p.dotlog.Logger.Errorf("failed to set up puncher. %s", err.Error())
	}

	p.connectSock()

	p.startConnDotSignal()
}
