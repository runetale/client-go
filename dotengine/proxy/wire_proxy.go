package proxy

import (
	"context"
	"errors"
	"net"

	"github.com/Notch-Technologies/dotshake/dotlog"
	"github.com/Notch-Technologies/dotshake/iface"
	"github.com/Notch-Technologies/dotshake/wireguard"
	"github.com/pion/ice/v2"
)

type WireProxy struct {
	iface *iface.Iface

	// proxy config
	remoteWgPubKey string // remote peer wg pub key
	remoteIp       string // remote peer ip
	wgIface        string // your wg iface
	addr           string // proxy addr
	preSharedKey   string // your preshared key

	remoteConn net.Conn
	localConn  net.Conn

	localProxyBuffer  []byte
	remoteProxyBuffer []byte

	ctx        context.Context
	cancelFunc context.CancelFunc

	dotlog *dotlog.DotLog
}

// TODO: (shinta) rewrite to proxy using sock5?
func NewWireProxy(
	iface *iface.Iface,
	remoteWgPubKey string,
	remoteip string,

	wgiface string,
	addr string,
	presharedkey string,
	dotlog *dotlog.DotLog,
) *WireProxy {
	ctx, cancel := context.WithCancel(context.Background())

	return &WireProxy{
		iface: iface,

		remoteWgPubKey: remoteWgPubKey,
		remoteIp:       remoteip,

		wgIface:      wgiface,
		addr:         addr,
		preSharedKey: presharedkey,

		localProxyBuffer:  make([]byte, 1500),
		remoteProxyBuffer: make([]byte, 1500),

		ctx:        ctx,
		cancelFunc: cancel,

		dotlog: dotlog,
	}
}

func (w *WireProxy) setup(remote *ice.Conn) error {
	w.remoteConn = remote
	udpConn, err := net.Dial("udp", w.addr)
	if err != nil {
		return err
	}
	w.localConn = udpConn

	return nil
}

func (w *WireProxy) configureToRemotePeer() error {
	err := w.iface.ConfigureToRemotePeer(
		w.remoteWgPubKey,
		w.remoteIp,
		w.localConn.LocalAddr().String(),
		wireguard.DefaultWgKeepAlive,
		w.preSharedKey,
	)
	if err != nil {
		w.dotlog.Logger.Errorf("")
		return err
	}

	return nil
}

func (w *WireProxy) Stop() error {
	w.cancelFunc()

	if w.localConn == nil {
		w.dotlog.Logger.Errorf("error is unexpected, you are most likely referring to locallConn without calling the setup function")
		return errors.New("error is unexpected")
	}

	err := w.iface.RemoveRemotePeer(w.wgIface, w.remoteIp, w.remoteWgPubKey)
	if err != nil {
		return err
	}

	return nil
}

func (w *WireProxy) StartProxy(remote *ice.Conn) error {
	err := w.setup(remote)
	if err != nil {
		return err
	}

	err = w.configureToRemotePeer()
	if err != nil {
		return err
	}

	w.startMon()

	return nil
}

func (w *WireProxy) startMon() {
	go w.monLocalToRemoteProxy()
	go w.monRemoteToLocalProxy()
}

func (w *WireProxy) monLocalToRemoteProxy() {
	for {
		select {
		case <-w.ctx.Done():
			return
		default:
			n, err := w.localConn.Read(w.remoteProxyBuffer)
			if err != nil {
				w.dotlog.Logger.Errorf("localConn cannot read remoteProxyBuffer [%s], size is %d", string(w.remoteProxyBuffer), n)
				continue
			}

			_, err = w.remoteConn.Write(w.remoteProxyBuffer[:n])
			if err != nil {
				w.dotlog.Logger.Errorf("localConn cannot write remoteProxyBuffer [%s], size is %d", string(w.remoteProxyBuffer), n)
				continue
			}

			// TODO: gathering buffer with dotmon
			w.dotlog.Logger.Debugf("remoteConn read remoteProxyBuffer [%s]", string(w.remoteProxyBuffer[:n]))
		}
	}
}

func (w *WireProxy) monRemoteToLocalProxy() {
	for {
		select {
		case <-w.ctx.Done():
			w.dotlog.Logger.Errorf("close the local proxy. close the remote ip here [%s], ", w.remoteIp)
			return
		default:
			n, err := w.remoteConn.Read(w.localProxyBuffer)
			if err != nil {
				w.dotlog.Logger.Errorf("remoteConn cannot read localProxyBuffer [%s], size is %d", string(w.localProxyBuffer), n)
				continue
			}

			_, err = w.localConn.Write(w.localProxyBuffer[:n])
			if err != nil {
				w.dotlog.Logger.Errorf("localConn cannot write localProxyBuffer [%s], size is %d", string(w.localProxyBuffer), n)
				continue
			}

			// TODO: gathering buffer with dotmon
			w.dotlog.Logger.Debugf("localConn read localProxyBuffer [%s]", string(w.localProxyBuffer[:n]))
		}
	}
}
