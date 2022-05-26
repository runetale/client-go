package polymer

import (
	"context"

	"github.com/Notch-Technologies/dotshake/client/grpc"
	"github.com/Notch-Technologies/dotshake/dotlog"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

type Polymer struct {
	signalClient grpc.SignalClientImpl
	serverClient grpc.ServerClientImpl
	dotlog       *dotlog.DotLog

	mk        string
	ip        string
	cidr      string
	wgPrivKey wgtypes.Key

	ctx    context.Context
	cancel context.CancelFunc
}

func NewPolymer(
	signalClient grpc.SignalClientImpl,
	serverClient grpc.ServerClientImpl,
	dotlog *dotlog.DotLog,
	ctx context.Context,
	cancel context.CancelFunc,
	mk string,
	ip string,
	cidr string,
	wgPrivKey string,
) (*Polymer, error) {
	pkey, err := wgtypes.ParseKey(wgPrivKey)
	if err != nil {
		return nil, err
	}

	return &Polymer{
		signalClient: signalClient,
		serverClient: serverClient,
		dotlog:       dotlog,

		mk:        mk,
		ip:        ip,
		cidr:      cidr,
		wgPrivKey: pkey,

		ctx:    ctx,
		cancel: cancel,
	}, nil
}

func (p *Polymer) Start() error {
	return nil
}
