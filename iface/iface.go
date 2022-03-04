package iface

import (
	"fmt"
	"net"
	"os/exec"
	"strings"
	"time"

	"github.com/Notch-Technologies/wizy/wireguard"
	"github.com/Notch-Technologies/wizy/wislog"
	"golang.zx2c4.com/wireguard/conn"
	"golang.zx2c4.com/wireguard/device"
	"golang.zx2c4.com/wireguard/tun"
	"golang.zx2c4.com/wireguard/wgctrl"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

type Iface struct {
	// your wireguard interface name
	Name string
	// your wireguard private key
	WgPrivateKey string
	// your ip
	IP string
	// your cidr range
	CIDR string

	wislog *wislog.WisLog
}

func NewIface(
	tunName, wgPrivateKey, ip, cidr string,
	wislog *wislog.WisLog,
) *Iface {
	return &Iface{
		Name:         tunName,
		WgPrivateKey: wgPrivateKey,
		IP:           ip,
		CIDR:         cidr,

		wislog: wislog,
	}
}

func (i *Iface) UpdatePeer(
	remotePeerPubKey string, allowedIps string,
	keepAlive time.Duration, endpoint string,
	preSharedKey *wgtypes.Key) error {

	fmt.Printf("updating interface %s peer %s: endpoint %s\n", i.Name, remotePeerPubKey, endpoint)
	_, ipNet, err := net.ParseCIDR(allowedIps)
	if err != nil {
		return err
	}

	peerKeyParsed, err := wgtypes.ParseKey(remotePeerPubKey)
	if err != nil {
		return err
	}
	peer := wgtypes.PeerConfig{
		PublicKey:                   peerKeyParsed,
		ReplaceAllowedIPs:           true,
		AllowedIPs:                  []net.IPNet{*ipNet},
		PersistentKeepaliveInterval: &keepAlive,
		PresharedKey:                preSharedKey,
	}

	config := wgtypes.Config{
		Peers: []wgtypes.PeerConfig{peer},
	}

	err = i.configureDevice(config)
	if err != nil {
		return err
	}

	if endpoint != "" {
		return i.updatePeerEndpoint(remotePeerPubKey, endpoint)
	}

	return nil
}

func (i *Iface) updatePeerEndpoint(remotePeerPubKey string,
	newEndpoint string) error {

	fmt.Printf("updating peer [%s] endpoint [%s]\n", remotePeerPubKey, newEndpoint)

	peerAddr, err := net.ResolveUDPAddr("udp4", newEndpoint)
	if err != nil {
		return err
	}

	fmt.Printf("parsed peer endpoint [%s]\n", peerAddr.String())

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

// RemovePeer removes a Wireguard Peer from the interface iface
func (i *Iface) RemovePeer(iface string, peerKey string) error {
	fmt.Printf("Removing peer %s from interface %s ", peerKey, iface)

	peerKeyParsed, err := wgtypes.ParseKey(peerKey)
	if err != nil {
		return err
	}

	peer := wgtypes.PeerConfig{
		PublicKey: peerKeyParsed,
		Remove:    true,
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

	_, err = wg.Device(i.Name)
	if err != nil {
		return err
	}

	fmt.Printf("create Wireguard device %s\n", i.Name)

	return wg.ConfigureDevice(i.Name, config)
}

func (i *Iface) addRoute(ipNet *net.IPNet) error {
	cmd := exec.Command("route", "add", "-net", ipNet.String(), "-interface", i.Name)
	if out, err := cmd.CombinedOutput(); err != nil {
		fmt.Printf("Command: %v failed with output %s and error: %v", cmd.String(), out, err)
		return err
	}

	return nil
}

func (i *Iface) assignAddr(address string) error {
	ip := strings.Split(address, "/")
	cmd := exec.Command("ifconfig", i.Name, "inet", address, ip[0])
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
	tunIface, err := tun.CreateTUN(i.Name, wireguard.DefaultMTU)
	if err != nil {
		return err
	}

	tunDevice := device.NewDevice(tunIface, conn.NewDefaultBind(), device.NewLogger(device.LogLevelSilent, "wissy: "))
	err = tunDevice.Up()
	if err != nil {
		return err
	}

	uapi, err := getUAPI(i.Name)
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
