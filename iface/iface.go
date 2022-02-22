package iface

import (
	"fmt"
	"net"
	"os/exec"
	"strings"
	"time"

	"golang.zx2c4.com/wireguard/conn"
	"golang.zx2c4.com/wireguard/device"
	"golang.zx2c4.com/wireguard/ipc"
	"golang.zx2c4.com/wireguard/tun"
	"golang.zx2c4.com/wireguard/wgctrl"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

const (
	defaultMTU = 1280
)

const wgPort = 51820

func UpdatePeer(
	iface string, peerKey string, allowedIps string,
	keepAlive time.Duration, endpoint string,
	preSharedKey *wgtypes.Key) error {

	fmt.Printf("updating interface %s peer %s: endpoint %s\n", iface, peerKey, endpoint)
	_, ipNet, err := net.ParseCIDR(allowedIps)
	if err != nil {
		return err
	}

	peerKeyParsed, err := wgtypes.ParseKey(peerKey)
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

	err = configureDevice(iface, config)
	if err != nil {
		return err
	}

	if endpoint != "" {
		return UpdatePeerEndpoint(iface, peerKey, endpoint)
	}

	return nil
}

func UpdatePeerEndpoint(iface string, peerKey string,
	newEndpoint string) error {

	fmt.Printf("updating peer %s endpoint %s\n", peerKey, newEndpoint)

	peerAddr, err := net.ResolveUDPAddr("udp4", newEndpoint)
	if err != nil {
		return err
	}

	fmt.Printf("parsed peer endpoint [%s]\n", peerAddr.String())

	peerKeyParsed, err := wgtypes.ParseKey(peerKey)
	if err != nil {
		return err
	}

	peer := wgtypes.PeerConfig{
		PublicKey:         peerKeyParsed,
		ReplaceAllowedIPs: false,
		UpdateOnly:        true,
		Endpoint:          peerAddr,
	}
	config := wgtypes.Config{
		Peers: []wgtypes.PeerConfig{peer},
	}
	return configureDevice(iface, config)
}

// RemovePeer removes a Wireguard Peer from the interface iface
func RemovePeer(iface string, peerKey string) error {
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

	return configureDevice(iface, config)
}

func configureDevice(iface string, config wgtypes.Config) error {
	wg, err := wgctrl.New()
	if err != nil {
		return err
	}
	defer wg.Close()

	_, err = wg.Device(iface)
	if err != nil {
		return err
	}
	fmt.Printf("create Wireguard device %s\n", iface)

	return wg.ConfigureDevice(iface, config)
}

// getUAPI returns a Listener
func getUAPI(iface string) (net.Listener, error) {
	tunSock, err := ipc.UAPIOpen(iface)
	if err != nil {
		return nil, err
	}
	return ipc.UAPIListen(iface, tunSock)
}

// assignAddr Adds IP address to the tunnel interface and network route based on the range provided
func assignAddr(address string, ifaceName string) error {
	ip := strings.Split(address, "/")
	cmd := exec.Command("ifconfig", ifaceName, "inet", address, ip[0])
	if out, err := cmd.CombinedOutput(); err != nil {
		fmt.Printf("Command: %v failed with output %s and error: %v", cmd.String(), out, err)
		return err
	}
	_, resolvedNet, err := net.ParseCIDR(address)
	err = addRoute(ifaceName, resolvedNet)
	if err != nil {
		fmt.Printf("Adding route failed with error: %v", err)
	}
	return nil
}

// addRoute Adds network route based on the range provided
func addRoute(iface string, ipNet *net.IPNet) error {
	cmd := exec.Command("route", "add", "-net", ipNet.String(), "-interface", iface)
	if out, err := cmd.CombinedOutput(); err != nil {
		fmt.Printf("Command: %v failed with output %s and error: %v", cmd.String(), out, err)
		return err
	}
	return nil
}

func CreateWithUserSpace(iface, address string) error {
	tunIface, err := tun.CreateTUN(iface, defaultMTU)
	if err != nil {
		return err
	}

	tunDevice := device.NewDevice(tunIface, conn.NewDefaultBind(), device.NewLogger(device.LogLevelSilent, "wissy: "))
	err = tunDevice.Up()
	if err != nil {
		return err
	}

	uapi, err := getUAPI(iface)
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

	fmt.Println("uapi handler started")
	err = assignAddr(address, iface)
	if err != nil {
		return err
	}
	return nil
}
