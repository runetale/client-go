package iface

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/Notch-Technologies/wizy/wireguard"
	"golang.zx2c4.com/wireguard/wgctrl"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

func execCmd(command string) (string, error) {
	args := strings.Fields(command)
	out, err := exec.Command(args[0], args[1:]...).Output()
	return string(out), err
}

func isWireGuardModule() bool {
	_, err := execCmd("modinfo wireguard")
	if err != nil {
		return false
	}

	return true
}

func CreateIface(i *Iface, ip, cidr string) error {
	addr := ip + "/" + cidr

	if isWireGuardModule() {
		return createWithKernelSpace(i.Name, i.WgPrivateKey, addr)
	}

	return createWithUserSpace(i, addr)
}

func createWithKernelSpace(ifaceName, privateKey, address string) error {
	ipCmd, err := exec.LookPath("ip")
	if err != nil {
		return err
	}

	key, err := wgtypes.ParseKey(privateKey)
	if err != nil {
		return err
	}

	wgClient, err := wgctrl.New()
	if err != nil {
		return err
	}
	defer wgClient.Close()

	del, err := execCmd(ipCmd + " link delete dev " + ifaceName)
	if err != nil {
		fmt.Println(del)
	}

	link, err := execCmd(ipCmd + " link add dev " + ifaceName + " type wireguard ")
	if err != nil {
		fmt.Printf("%s, %v", link, err)
		return err
	}

	add, err := execCmd(ipCmd + " address add dev " + ifaceName + " " + address)
	if err != nil {
		fmt.Printf("%s, %v", add, err)
		return err
	}

	fMark := 0
	port := wireguard.WgPort
	wgConf := wgtypes.Config{
		PrivateKey:   &key,
		ReplacePeers: false,
		FirewallMark: &fMark,
		ListenPort:   &port,
	}

	_, err = wgClient.Device(ifaceName)
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = wgClient.ConfigureDevice(ifaceName, wgConf)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("device does not exist %s.", err.Error())
		} else {
			fmt.Println(err)
		}
		return err
	}

	if up, err := execCmd(ipCmd + " link set up dev " + ifaceName); err != nil {
		fmt.Printf("%s, %v", up, err)
		return err
	}

	return nil
}

func createWithUserSpace(i *Iface, address string) error {
	err := i.CreateWithUserSpace(address)
	if err != nil {
		return err
	}

	key, err := wgtypes.ParseKey(i.WgPrivateKey)
	if err != nil {
		return err
	}

	fwmark := 0
	port := wireguard.WgPort
	config := wgtypes.Config{
		PrivateKey:   &key,
		ReplacePeers: false,
		FirewallMark: &fwmark,
		ListenPort:   &port,
	}

	return i.configureDevice(config)
}
