package iface

import (
	"fmt"
	"net"
	"time"

	"golang.zx2c4.com/wireguard/wgctrl"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

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

	fmt.Printf("got Wireguard device %s\n", iface)

	return wg.ConfigureDevice(iface, config)
}

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
		PublicKey: peerKeyParsed,
		ReplaceAllowedIPs: true,
		AllowedIPs: []net.IPNet{*ipNet},
		PersistentKeepaliveInterval: &keepAlive,
		PresharedKey: preSharedKey,
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
