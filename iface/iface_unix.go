//go:build linux || darwin
// +build linux darwin

package iface

import (
	"net"

	"golang.zx2c4.com/wireguard/ipc"
)

// getUAPI returns a Listener
func getUAPI(iface string) (net.Listener, error) {
	tunSock, err := ipc.UAPIOpen(iface)
	if err != nil {
		return nil, err
	}
	return ipc.UAPIListen(iface, tunSock)
}
