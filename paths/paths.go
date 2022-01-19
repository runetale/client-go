package paths

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

// json file to manage the wics server startup config.
func DefaultWicsConfigFile() string {
	return "/etc/wizy/wics.json"
}

// state file to manage the account state.
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

// state file to manage the secret information of the wics server.
func DefaultWicsServerStateFile() string {
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

func MkStateDir(dirPath string) error {
	if err := os.MkdirAll(dirPath, 0700); err != nil {
		return err
	}

	return checkStateDirPermission(dirPath)
}

func checkStateDirPermission(dir string) error {
	const (
		perm = 700
	)

	if filepath.Base(dir) != "wizy" {
		return nil
	}

	fi, err := os.Stat(dir)
	if err != nil {
		return err
	}

	if !fi.IsDir() {
		return fmt.Errorf("expected %q is a directory, but %v", dir, fi.Mode())
	}

	if fi.Mode().Perm() == perm {
		return nil
	}

	return os.Chmod(dir, perm)
}
