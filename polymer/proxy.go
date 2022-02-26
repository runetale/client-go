package polymer

import (
	"context"
	"fmt"
	"io"
	"net"
	"time"

	"github.com/Notch-Technologies/wizy/iface"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

const DefaultWgKeepAlive = 25 * time.Second

type ProxyConfig struct {
	WgListenAddr string
	RemoteKey    string
	WgInterface  string
	AllowedIps   string
	PreSharedKey *wgtypes.Key
}

type Proxyer interface {
	io.Closer
	// Start creates a local remoteConn and starts proxying data from/to remoteConn
	Start(remoteConn net.Conn) error
}

type Proxy struct {
	ctx    context.Context
	cancel context.CancelFunc

	config ProxyConfig

	remoteConn net.Conn
	localConn  net.Conn
}

func NewProxy(config ProxyConfig) *Proxy {
	p := &Proxy{config: config}
	p.ctx, p.cancel = context.WithCancel(context.Background())
	return p
}

func (p *Proxy) updateEndpoint() error {
	err := iface.UpdatePeer(p.config.WgInterface, p.config.RemoteKey,
		p.config.AllowedIps, DefaultWgKeepAlive,
		p.localConn.LocalAddr().String(), p.config.PreSharedKey)
	if err != nil {
		return err
	}
	return nil
}

func (p *Proxy) Start(remoteConn net.Conn) error {
	p.remoteConn = remoteConn

	var err error
	p.localConn, err = net.Dial("udp", p.config.WgListenAddr)
	if err != nil {
		fmt.Errorf("failed dialing to local Wireguard port %s", err)
		return err
	}

	err = p.updateEndpoint()
	if err != nil {
		fmt.Errorf("error while updating Wireguard peer endpoint [%s] %v", p.config.RemoteKey, err)
		return err
	}

	go p.proxyToRemote()
	go p.proxyToLocal()

	return nil
}

func (p *Proxy) Close() error {
	p.cancel()
	if c := p.localConn; c != nil {
		err := p.localConn.Close()
		if err != nil {
			return err
		}
	}
	err := iface.RemovePeer(p.config.WgInterface, p.config.RemoteKey)
	if err != nil {
		return err
	}
	return nil
}

func (p *Proxy) proxyToRemote() {

	buf := make([]byte, 1500)
	for {
		select {
		case <-p.ctx.Done():
			fmt.Printf("stopped proxying to remote peer %s due to closed connection\n", p.config.RemoteKey)
			return
		default:
			n, err := p.localConn.Read(buf)
			if err != nil {
				continue
			}

			_, err = p.remoteConn.Write(buf[:n])
			if err != nil {
				continue
			}
		}
	}
}

func (p *Proxy) proxyToLocal() {

	buf := make([]byte, 1500)
	for {
		select {
		case <-p.ctx.Done():
			fmt.Printf("stopped proxying from remote peer %s due to closed connection\n", p.config.RemoteKey)
			return
		default:
			n, err := p.remoteConn.Read(buf)
			if err != nil {
				continue
			}

			_, err = p.localConn.Write(buf[:n])
			if err != nil {
				continue
			}
		}
	}
}
