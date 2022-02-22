package iface

import (
	"fmt"
	"net"
	"os/exec"
	"strings"

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

func CreateIface(ifaceName, privateKey, address string) error {
	err := CreateWithUserSpace(ifaceName, address)
	if err != nil {
		return err
	}

	key, err := wgtypes.ParseKey(privateKey)
	if err != nil {
		return err
	}

	fwmark := 0
	port := wgPort
	config := wgtypes.Config{
		PrivateKey:   &key,
		ReplacePeers: false,
		FirewallMark: &fwmark,
		ListenPort:   &port,
	}

	return configureDevice(ifaceName, config)
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
