package paths

import (
	"runtime"
)

// File that manages the configuration of the wics server.
func DefaultWicsConfigFile() string {
	return "/etc/wizy/wics.json"
}

// State file to wics the state of the Store.
func DefaultAccountStateFile() string {
	switch runtime.GOOS {
	case "freebsd", "openbsd":
		return "/var/db/wizy/account.state"
	case "linux":
		return "/var/lib/wizy/account.state"
	case "darwin":
		return "/Library/wizy/account.state"
	default:
		return ""
	}
}

// State file to wics the state of the Store.
func DefaultServerStateFile() string {
	switch runtime.GOOS {
	case "freebsd", "openbsd":
		return "/var/db/wizy/server.state"
	case "linux":
		return "/var/lib/wizy/server.state"
	case "darwin":
		return "/Library/wizy/server.state"
	default:
		return ""
	}
}
