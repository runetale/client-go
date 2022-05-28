package paths

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

// state file to manage the secret information of the server.
// do not disclose to the outside world.
func DefaultWissyServerStateFile() string {
	switch runtime.GOOS {
	case "freebsd", "openbsd":
		return "/var/db/dotshake/server.state"
	case "linux":
		return "/var/lib/dotshake/server.state"
	case "darwin":
		return "/Library/dotshake/server.state"
	default:
		return ""
	}
}

// state file to manage the secret information of the server.
// do not disclose to the outside world.
func DefaultDotshakeClientStateFile() string {
	switch runtime.GOOS {
	case "freebsd", "openbsd":
		return "/var/db/dotshake/client.state"
	case "linux":
		return "/var/lib/dotshake/client.state"
	case "darwin":
		return "/Library/dotshake/client.state"
	default:
		return ""
	}
}

func DefaultClientConfigFile() string {
	return "/etc/dotshake/client.json"
}

// json file to manage the server startup config.
func DefaultServerConfigFile() string {
	return "/etc/dotshake/server.json"
}

func DefaultClientLogFile() string {
	return "/var/log/dotshake/client.log"
}

func DefaultServerLogFile() string {
	return "/var/log/dotshake/server.log"
}

func DefaultSignalingLogFile() string {
	return "/var/log/dotshake/signaling.log"
}

func DefaultLetsEncryptDir() string {
	switch runtime.GOOS {
	case "freebsd", "openbsd":
		return "/var/db/dotshake/letsencrypt"
	case "linux":
		return "/var/lib/dotshake/letsencrypt"
	case "darwin":
		return "/Library/dotshake/letsencrypt"
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

	if filepath.Base(dir) != "dotshake" {
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