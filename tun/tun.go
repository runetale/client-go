package tun

import "runtime"

func TunName() string {
	switch runtime.GOOS {
	case "openbsd":
		return "tun"
	case "linux":
		return "ds0"
	case "darwin":
		return "utun100"
	case "windows":
		return "dotshake"
	}
	return "ds0"
}
