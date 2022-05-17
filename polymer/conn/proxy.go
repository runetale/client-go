package conn

import (
	"context"
	"fmt"
	"io"
	"net"
	"time"

	"github.com/Notch-Technologies/dotshake/iface"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

const DefaultWgKeepAlive = 25 * time.Second

type ProxyConfig struct {
	WgListenAddr     string
	RemotePeerPubKey string
	WgInterface      string
	AllowedIPs       string
	PreSharedKey     *wgtypes.Key
}

func NewProxyConfig(
	wgPort int,
	remotePeerPubKey string,
	wgIface string,
	allowedIPs string,
	preSharedKey *wgtypes.Key,

) *ProxyConfig {
	return &ProxyConfig{
		WgListenAddr:     fmt.Sprintf("127.0.0.1:%d", wgPort),
		RemotePeerPubKey: remotePeerPubKey,
		WgInterface:      wgIface,
		AllowedIPs:       allowedIPs,
		PreSharedKey:     preSharedKey,
	}
}

type Proxyer interface {
	io.Closer
	// Start creates a local remoteConn and starts proxying data from/to remoteConn
	Start(remoteConn net.Conn) error
}

type Proxy struct {
	ctx    context.Context
	cancel context.CancelFunc

	config *ProxyConfig

	iface *iface.Iface

	remoteConn net.Conn
	localConn  net.Conn
}

func NewProxy(
	config *ProxyConfig,
	iface *iface.Iface,
) *Proxy {
	p := &Proxy{
		config: config,
		iface:  iface,
	}
	p.ctx, p.cancel = context.WithCancel(context.Background())
	return p
}

func (p *Proxy) updateEndpoint() error {
	err := p.iface.UpdatePeer(p.config.RemotePeerPubKey,
		p.config.AllowedIPs, DefaultWgKeepAlive,
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
		fmt.Printf("failed dialing to local Wireguard port %s", err)
		return err
	}

	err = p.updateEndpoint()
	if err != nil {
		fmt.Printf("error while updating Wireguard peer endpoint [%s] %v\n", p.config.RemotePeerPubKey, err)
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
	err := p.iface.RemovePeer(p.config.WgInterface, p.config.RemotePeerPubKey)
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
			fmt.Printf("stopped proxying to remote peer %s due to closed connection\n", p.config.RemotePeerPubKey)
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
			fmt.Printf("stopped proxying from remote peer %s due to closed connection\n", p.config.RemotePeerPubKey)
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
