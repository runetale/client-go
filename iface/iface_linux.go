package iface

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"golang.zx2c4.com/wireguard/wgctrl"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

func execCmd(command string) (string, error) {
	args := strings.Fields(command)
	out, err := exec.Command(args[0], args[1:]...).Output()
	return string(out), err
}

func isWireGuardModule() bool {
	out, err := execCmd("modinfo wireguard")
	if err != nil {
		return false
	}

	fmt.Println(out)

	return true
}

func CreateIface(ifaceName, privateKey, address string) {
	if isWireGuardModule() {
		createWithKernelSpace(ifaceName, privateKey, address)
	}

	createWithUserSpace()
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

	_, err = execCmd(ipCmd + " link delete dev " + ifaceName)
	if err != nil {
		return err
	}

	link, err := execCmd(ipCmd + " link add dev " + ifaceName + " type wireguard ")
	if err != nil {
		fmt.Println(link)
		return err
	}

	add, err := execCmd(ipCmd + " address add dev " + ifaceName + " " + address + "/24")
	if err != nil {
		fmt.Println(add)
		return err
	}

	fMark := 0
	port := 51820
	wgConf := wgtypes.Config{
		PrivateKey:   &key,
		ReplacePeers: false,
		FirewallMark: &fMark,
		ListenPort:   &port,
	}

	_, err = wgClient.Device(ifaceName)
	if err != nil {
		return err
	}

	err = wgClient.ConfigureDevice(ifaceName, wgConf)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Device does not exist: ")
			fmt.Println(err)
		} else {
			fmt.Printf("This is inconvenient: %v", err)
		}
	}

	return nil
}

func createWithUserSpace() {}
