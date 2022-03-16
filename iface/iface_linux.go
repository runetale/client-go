package iface

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/Notch-Technologies/wizy/version"
	"github.com/Notch-Technologies/wizy/wireguard"
	"github.com/Notch-Technologies/wizy/wislog"
	"golang.zx2c4.com/wireguard/wgctrl"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

func execCmd(command string) (string, error) {
	args := strings.Fields(command)
	out, err := exec.Command(args[0], args[1:]...).Output()
	return string(out), err
}

func isWireGuardModule(
	wislog *wislog.WisLog,
) bool {
	_, err := execCmd("modinfo wireguard")
	if err != nil {
		wislog.Logger.Infof("cannot get modinfo wireguard: %s", err.Error())
		return false
	}

	wislog.Logger.Infof("get modinfo wireguard: %s")
	return true
}

func CreateIface(
	i *Iface, ip, cidr string,
	wislog *wislog.WisLog,
) error {
	addr := ip + "/" + cidr

	if version.Get() == version.NixOS {
		return createWithKernelSpace(i.Name, i.WgPrivateKey, addr, wislog)
	}

	if isWireGuardModule(wislog) {
		wislog.Logger.Infof("wireguard in the kernel space.")
		return createWithKernelSpace(i.Name, i.WgPrivateKey, addr, wislog)
	}

	wislog.Logger.Infof("wireguard in the user space.")
	return createWithUserSpace(i, addr)
}

func createWithKernelSpace(
	ifaceName, privateKey, address string,
	wislog *wislog.WisLog,
) error {
	ipCmd, err := exec.LookPath("ip")
	if err != nil {
		wislog.Logger.Errorf("failed to ip command: %s", err.Error())
		return err
	}

	key, err := wgtypes.ParseKey(privateKey)
	if err != nil {
		wislog.Logger.Errorf("failed to parsing private key: %s", err.Error())
		return err
	}

	wgClient, err := wgctrl.New()
	if err != nil {
		wislog.Logger.Errorf("failed to wireguard client: %s", err.Error())
		return err
	}
	defer wgClient.Close()

	del, err := execCmd(ipCmd + " link delete dev " + ifaceName)
	if err != nil {
		wislog.Logger.Errorf("failed to link delete: %s", err.Error())
		fmt.Println(del)
	}

	link, err := execCmd(ipCmd + " link add dev " + ifaceName + " type wireguard ")
	if err != nil {
		wislog.Logger.Errorf("failed to link add dev. ifaceName: [%s]", ifaceName)
		fmt.Printf("%s, %v", link, err)
		return err
	}

	add, err := execCmd(ipCmd + " address add dev " + ifaceName + " " + address)
	if err != nil {
		wislog.Logger.Errorf("failed to address add dev. ifaceName: [%s], address: [%s]", ifaceName, address)
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
		wislog.Logger.Errorf("failed to create wireguard device. ifaceName: [%s]", ifaceName)
		fmt.Println(err)
		return err
	}

	err = wgClient.ConfigureDevice(ifaceName, wgConf)
	if err != nil {
		if os.IsNotExist(err) {
			wislog.Logger.Errorf("device does not exist %s.", ifaceName)
			fmt.Printf("device does not exist %s.", err.Error())
		} else {
			wislog.Logger.Errorf("%s.", err.Error())
			fmt.Println(err)
		}
		return err
	}

	if up, err := execCmd(ipCmd + " link set up dev " + ifaceName); err != nil {
		wislog.Logger.Errorf("%s, %s", ifaceName, err.Error())
		fmt.Printf("%s, %s", up, err.Error())
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
