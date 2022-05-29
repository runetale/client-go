package daemon

import (
	"github.com/Notch-Technologies/dotshake/dotlog"
)

const (
	success      = "\t[\033[32mOK\033[0m]"
	failed       = "\t[\033[31mFAILED\033[0m]"
	notinstalled = "\t[\033[31mNOT_INSTALLED\033[0m]"
	notrunning   = "\t[\033[31mNOT_RUNNING\033[0m]"
)

type Daemon interface {
	Install() error
	Uninstall() error
	Status() string
}

func NewDaemon(
	binPath, serviceName, daemonFilePath, systemConfig string,
	wl *dotlog.DotLog,
) Daemon {
	return newDaemon(binPath, serviceName, daemonFilePath, systemConfig, wl)
}
