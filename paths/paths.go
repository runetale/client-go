package paths

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

// json file to manage the wics server startup config.
func DefaultWicsConfigFile() string {
	return "/etc/wissy/wics.json"
}

// state file to manage the secret information of the wics server.
// do not disclose to the outside world.
func DefaultWicsServerStateFile() string {
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

func DefaultClientConfigFile() string {
	return "/etc/wissy/client.json"
}

func DefaultClientLogFile() string {
	return "/var/log/wissy/client.log"
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
