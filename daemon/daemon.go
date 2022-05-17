package daemon

import "github.com/Notch-Technologies/dotshake/dotlog"

type Daemon interface {
	Install() error
	Uninstall() error

	Load() error
	Unload() error

	Start() error
	Stop() error

	Status() error

	IsInstalled() bool
	IsRunnning() (string, bool)
}

func NewDaemon(
	binPath, serviceName, daemonFilePath, systemConfig string,
	wl *dotlog.DotLog,
) Daemon {
	return newDaemon(binPath, serviceName, daemonFilePath, systemConfig, wl)
}
