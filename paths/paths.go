package paths

import (
	"runtime"
)

// File that manages the configuration of the wics server.
func DefaultWicsFile() string {
	return "/etc/wizy/wics.json"
}

// State file to wics the state of the Store.
func DefaultWicsStateFile() string {
	switch runtime.GOOS {
	case "freebsd", "openbsd":
		return "/var/db/wizy/wics.state"
	case "linux":
		return "/var/lib/wizy/wics.state"
	case "darwin":
		return "/Library/wizy/wics.state"
	default:
		return ""
	}
}
