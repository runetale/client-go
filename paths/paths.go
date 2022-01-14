package paths

import (
	"runtime"
)

// File that manages the configuration of the management server.
func DefaultManagementFile() string {
	return "/etc/wizy/management.json"
}

// State file to manage the state of the Store.
func DefaultStoreStateFile() string {
	switch runtime.GOOS {
	case "freebsd", "openbsd":
		return "/var/db/wizy/management.state"
	case "linux":
		return "/var/lib/wizy/management.state"
	case "darwin":
		return "/Library/wizy/management.state"
	default:
		return ""
	}
}
