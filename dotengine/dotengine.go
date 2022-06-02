package dotengine

import (
	"context"
	"errors"
	"sync"

	"github.com/Notch-Technologies/dotshake/client/grpc"
	"github.com/Notch-Technologies/dotshake/dotengine/peer"
	"github.com/Notch-Technologies/dotshake/dotlog"
	"github.com/Notch-Technologies/dotshake/iface"
	"github.com/Notch-Technologies/dotshake/rcn/unixsock"
	"github.com/Notch-Technologies/dotshake/wireguard"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

type DotEngine struct {
	dotlog *dotlog.DotLog

	mk        string
	tunName   string
	ip        string
	cidr      string
	wgPrivKey string
	wgPort    int
	blackList []string

	peer *peer.Peer

	ctx    context.Context
	cancel context.CancelFunc

	mu *sync.Mutex

	rootch chan struct{}

	sock *unixsock.PolymerSock
}

func NewDotEngine(
	serverClient grpc.ServerClientImpl,
	dotlog *dotlog.DotLog,
	tunName string,
	mk string,
	ip string,
	cidr string,
	wgPrivKey string,
	blackList []string,
	ctx context.Context,
	cancel context.CancelFunc,
) (*DotEngine, error) {
	_, err := wgtypes.ParseKey(wgPrivKey)
	if err != nil {
		return nil, err
	}

	ch := make(chan struct{})
	mu := &sync.Mutex{}

	sock := unixsock.NewPolymerSock(dotlog, ch, nil)

	return &DotEngine{
		dotlog: dotlog,

		mk:        mk,
		tunName:   tunName,
		ip:        ip,
		cidr:      cidr,
		wgPrivKey: wgPrivKey,
		wgPort:    wireguard.WgPort,
		blackList: blackList,

		peer: peer.NewPeer(serverClient, mk, dotlog, sock),

		ctx:    ctx,
		cancel: cancel,

		mu: mu,

		rootch: ch,

		sock: sock,
	}, nil
}

func (p *DotEngine) createIface() error {
	i := iface.NewIface(p.tunName, p.wgPrivKey, p.ip, p.cidr, p.dotlog)
	return iface.CreateIface(i, p.ip, p.cidr, p.dotlog)
}

func (p *DotEngine) test() error {
	return p.sock.DialPuncherSignal(&unixsock.SignalSock{Commands: 1})
}

func (p *DotEngine) Start() error {
	p.mu.Lock()
	defer p.mu.Unlock()

	err := p.createIface()
	if err != nil {
		return err
	}

	p.peer.Up()

	err = p.test()
	if err != nil {
		return err
	}

	go func() {
		// do somethings
		// system resouce check?
	}()
	<-p.rootch

	return errors.New("stop the dotengine")
}
