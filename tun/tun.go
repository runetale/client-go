package tun

import "runtime"

func TunName() string {
	switch runtime.GOOS {
	case "openbsd":
		return "tun"
	case "linux":
		return "ws0"
	case "darwin":
		return "utun100"
	case "windows":
		return "dotshake"
	}
	return "ws0"
}
