package iface

import (
	"fmt"
	"net"
	"os/exec"
	"strings"
	"time"

	"github.com/Notch-Technologies/dotshake/dotlog"
	"github.com/Notch-Technologies/dotshake/wireguard"
	"golang.zx2c4.com/wireguard/conn"
	"golang.zx2c4.com/wireguard/device"
	"golang.zx2c4.com/wireguard/tun"
	"golang.zx2c4.com/wireguard/wgctrl"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

type Iface struct {
	// your wireguard interface name
	Tun string
	// your wireguard private key
	WgPrivateKey string
	// your ip
	IP string
	// your cidr range
	CIDR string

	dotlog *dotlog.DotLog
}

func NewIface(
	tunName, wgPrivateKey, ip string,
	cidr string, dotlog *dotlog.DotLog,
) *Iface {
	return &Iface{
		Tun:          tunName,
		WgPrivateKey: wgPrivateKey,
		IP:           ip,
		CIDR:         cidr,

		dotlog: dotlog,
	}
}

func (i *Iface) ConfigureToRemotePeer(
	remotePeerPubKey, remoteip, endpoint string,
	keepAlive time.Duration,
	preSharedKey string,
) error {
	i.dotlog.Logger.Debugf("configuring %s to remote peer [%s], your endpoint [%s]", i.Tun, remotePeerPubKey, endpoint)

	_, ipNet, err := net.ParseCIDR(remoteip)
	if err != nil {
		return err
	}

	parsedRemotePeerPubKey, err := wgtypes.ParseKey(remotePeerPubKey)
	if err != nil {
		return err
	}

	var parsedPreSharedkey wgtypes.Key
	if preSharedKey != "" {
		parsedPreSharedkey, err = wgtypes.ParseKey(preSharedKey)
		if err != nil {
			return err
		}
	}

	peer := wgtypes.PeerConfig{
		PublicKey:                   parsedRemotePeerPubKey,
		ReplaceAllowedIPs:           true,
		AllowedIPs:                  []net.IPNet{*ipNet},
		PersistentKeepaliveInterval: &keepAlive,
		PresharedKey:                &parsedPreSharedkey,
	}

	config := wgtypes.Config{
		Peers: []wgtypes.PeerConfig{peer},
	}

	err = i.configureDevice(config)
	if err != nil {
		return err
	}

	if endpoint != "" {
		return i.updatePeerEndpoint(remotePeerPubKey, remoteip, endpoint)
	}

	return nil
}

func (i *Iface) updatePeerEndpoint(
	remotePeerPubKey string,
	remoteip string,
	newEndpoint string,
) error {
	i.dotlog.Logger.Debugf("updating your ip [%s], here is remote ip and remote wg pub key => [%s:%s]", newEndpoint, remoteip, remotePeerPubKey)

	peerAddr, err := net.ResolveUDPAddr("udp4", newEndpoint)
	if err != nil {
		return err
	}

	pubKey, err := wgtypes.ParseKey(remotePeerPubKey)
	if err != nil {
		return err
	}

	peer := wgtypes.PeerConfig{
		PublicKey:         pubKey,
		ReplaceAllowedIPs: false,
		UpdateOnly:        true,
		Endpoint:          peerAddr,
	}

	config := wgtypes.Config{
		Peers: []wgtypes.PeerConfig{peer},
	}

	return i.configureDevice(config)
}

func (i *Iface) configureDevice(config wgtypes.Config) error {
	wg, err := wgctrl.New()
	if err != nil {
		return err
	}
	defer wg.Close()

	_, err = wg.Device(i.Tun)
	if err != nil {
		return err
	}

	fmt.Printf("create wg device %s\n", i.Tun)

	return wg.ConfigureDevice(i.Tun, config)
}

func (i *Iface) addRoute(ipNet *net.IPNet) error {
	cmd := exec.Command("route", "add", "-net", ipNet.String(), "-interface", i.Tun)
	if out, err := cmd.CombinedOutput(); err != nil {
		fmt.Printf("Command: %v failed with output %s and error: %v", cmd.String(), out, err)
		return err
	}

	return nil
}

func (i *Iface) assignAddr(address string) error {
	ip := strings.Split(address, "/")
	cmd := exec.Command("ifconfig", i.Tun, "inet", address, ip[0])
	if out, err := cmd.CombinedOutput(); err != nil {
		fmt.Printf("Command: %v failed with output %s and error: %v", cmd.String(), out, err)
		return err
	}

	_, resolvedNet, err := net.ParseCIDR(address)
	if err != nil {
		return err
	}

	err = i.addRoute(resolvedNet)
	if err != nil {
		fmt.Printf("Adding route failed with error: %v", err)
	}

	return nil
}

func (i *Iface) CreateWithUserSpace(address string) error {
	tunIface, err := tun.CreateTUN(i.Tun, wireguard.DefaultMTU)
	if err != nil {
		return err
	}

	tunDevice := device.NewDevice(tunIface, conn.NewDefaultBind(), device.NewLogger(device.LogLevelSilent, "dotshake: "))
	err = tunDevice.Up()
	if err != nil {
		return err
	}

	uapi, err := getUAPI(i.Tun)
	if err != nil {
		return err
	}

	go func() {
		for {
			conn, err := uapi.Accept()
			if err != nil {
				fmt.Printf("uapi accept failed with error: %v\n", err)
				continue
			}
			go tunDevice.IpcHandle(conn)
		}
	}()

	err = i.assignAddr(address)
	if err != nil {
		return err
	}

	return nil
}

func (i *Iface) RemoveRemotePeer(iface string, remoteip, remotePeerPubKey string) error {
	i.dotlog.Logger.Debugf("delete %s on this %s", remotePeerPubKey, i.Tun)

	peerKeyParsed, err := wgtypes.ParseKey(remotePeerPubKey)
	if err != nil {
		return err
	}

	peer := wgtypes.PeerConfig{
		Remove:    true,
		PublicKey: peerKeyParsed,
	}

	config := wgtypes.Config{
		Peers: []wgtypes.PeerConfig{peer},
	}

	return i.configureDevice(config)
}
