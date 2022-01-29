package tun

import "runtime"


func TunName() string {
	switch runtime.GOOS {
	case "openbsd":
		return "tun"
	case "linux":
		return "ws0"
	case "darwin":
		return "utun"
	case "windows":
		return "Wissy"
	}
	return "ws0"
}
