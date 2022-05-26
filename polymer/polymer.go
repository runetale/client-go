package polymer

import (
	"context"
	"errors"

	"github.com/Notch-Technologies/dotshake/client/grpc"
	"github.com/Notch-Technologies/dotshake/dotlog"
	"github.com/Notch-Technologies/dotshake/iface"
	"github.com/Notch-Technologies/dotshake/polymer/dotsignal"
	"github.com/Notch-Technologies/dotshake/wireguard"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

type Polymer struct {
	serverClient grpc.ServerClientImpl
	dotlog       *dotlog.DotLog

	mk        string
	tunName   string
	ip        string
	cidr      string
	wgPrivKey string
	wgPort    int
	blackList []string

	ds *dotsignal.DotSignal

	ctx    context.Context
	cancel context.CancelFunc

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

	return &Polymer{
		serverClient: serverClient,
		dotlog:       dotlog,

		mk:        mk,
		tunName:   tunName,
		ip:        ip,
		cidr:      cidr,
		wgPrivKey: wgPrivKey,
		wgPort:    wireguard.WgPort,
		blackList: blackList,

		ds: dotsignal.NewDotSignal(signalClient, mk, rootch),

		ctx:    ctx,
		cancel: cancel,

		rootch: rootch,
	}, nil
}

func (p *Polymer) createIface() error {
	i := iface.NewIface(p.tunName, p.wgPrivKey, p.ip, p.cidr, p.dotlog)
	return iface.CreateIface(i, p.ip, p.cidr, p.dotlog)
}

func (p *Polymer) Start() error {
	err := p.createIface()
	if err != nil {
		return err
	}

	go func() {
		// do somethings
		// system resouce check?
	}()
	<-p.rootch

	return errors.New("stop the polymer")
}
