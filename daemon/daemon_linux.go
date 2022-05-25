package daemon

import (
	"os"

	"github.com/Notch-Technologies/dotshake/dotlog"
)

func newDaemon(
	binPath, serviceName, daemonFilePath, systemConfig string,
	wl *dotlog.DotLog,
) Daemon {
	if _, err := os.Stat("/run/systemd/system"); err == nil {
		return &systemDRecord{
			binPath:        binPath,
			serviceName:    serviceName,
			daemonFilePath: daemonFilePath,
			systemConfig:   systemConfig,

			dotlog: wl,
		}
	}
	if _, err := os.Stat("/sbin/initctl"); err == nil {
		return &upstartRecord{
			binPath:        binPath,
			serviceName:    serviceName,
			daemonFilePath: daemonFilePath,
			systemConfig:   systemConfig,

			dotlog: wl,
		}
	}

	return &systemVRecord{
		binPath:        binPath,
		serviceName:    serviceName,
		daemonFilePath: daemonFilePath,
		systemConfig:   systemConfig,

		dotlog: wl,
	}
}
