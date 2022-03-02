package paths

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

// json file to manage the server startup config.
func DefaultServerConfigFile() string {
	return "/etc/wissy/config.json"
}

// state file to manage the secret information of the server.
// do not disclose to the outside world.
func DefaultWissyServerStateFile() string {
	switch runtime.GOOS {
	case "freebsd", "openbsd":
		return "/var/db/wissy/server.state"
	case "linux":
		return "/var/lib/wissy/server.state"
	case "darwin":
		return "/Library/wissy/server.state"
	default:
		return ""
	}
}

// state file to manage the secret information of the server.
// do not disclose to the outside world.
func DefaultWicsClientStateFile() string {
	switch runtime.GOOS {
	case "freebsd", "openbsd":
		return "/var/db/wissy/client.state"
	case "linux":
		return "/var/lib/wissy/client.state"
	case "darwin":
		return "/Library/wissy/client.state"
	default:
		return ""
	}
}

func DefaultClientConfigFile() string {
	return "/etc/wissy/client.json"
}

func DefaultClientLogFile() string {
	return "/var/log/wissy/client.log"
}

func DefaultServerLogFile() string {
	return "/var/log/wissy/server.log"
}

func DefaultSignalingLogFile() string {
	return "/var/log/wissy/signaling.log"
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

	if filepath.Base(dir) != "wissy" {
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
