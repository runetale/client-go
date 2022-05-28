package polymer

import (
	"context"
	"errors"
	"sync"

	"github.com/Notch-Technologies/dotshake/client/grpc"
	"github.com/Notch-Technologies/dotshake/dotlog"
	"github.com/Notch-Technologies/dotshake/iface"
	"github.com/Notch-Technologies/dotshake/polymer/dotmachine"
	"github.com/Notch-Technologies/dotshake/polymer/dotsignal"
	"github.com/Notch-Technologies/dotshake/wireguard"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

type Polymer struct {
	dotlog *dotlog.DotLog

	mk        string
	tunName   string
	ip        string
	cidr      string
	wgPrivKey string
	wgPort    int
	blackList []string

	ds *dotsignal.DotSignal
	dm *dotmachine.DotMachine

	ctx    context.Context
	cancel context.CancelFunc

	mu *sync.Mutex

	rootch chan struct{}
}

func NewPolymer(
	signalClient grpc.SignalClientImpl,
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
) (*Polymer, error) {
	_, err := wgtypes.ParseKey(wgPrivKey)
	if err != nil {
		return nil, err
	}

	rootch := make(chan struct{})
	mu := &sync.Mutex{}

	return &Polymer{
		dotlog: dotlog,

		mk:        mk,
		tunName:   tunName,
		ip:        ip,
		cidr:      cidr,
		wgPrivKey: wgPrivKey,
		wgPort:    wireguard.WgPort,
		blackList: blackList,

		ds: dotsignal.NewDotSignal(signalClient, mk, rootch, mu, dotlog),
		dm: dotmachine.NewDotMachine(serverClient, mk, rootch, mu, dotlog),

		ctx:    ctx,
		cancel: cancel,

		mu: mu,

		rootch: rootch,
	}, nil
}

func (p *Polymer) createIface() error {
	i := iface.NewIface(p.tunName, p.wgPrivKey, p.ip, p.cidr, p.dotlog)
	return iface.CreateIface(i, p.ip, p.cidr, p.dotlog)
}

func (p *Polymer) Start() error {
	p.mu.Lock()
	defer p.mu.Unlock()

	err := p.createIface()
	if err != nil {
		return err
	}

	p.ds.ConnectDotSignal()
	p.dm.Up()

	go func() {
		// do somethings
		// system resouce check?
	}()
	<-p.rootch

	return errors.New("stop the polymer")
}
