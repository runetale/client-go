package daemon

import (
	"os"

	"github.com/Notch-Technologies/dotshake/wislog"
)

func newDaemon(
	binPath, serviceName, daemonFilePath, systemConfig string,
	wl *wislog.WisLog,
) Daemon {
	if _, err := os.Stat("/run/systemd/system"); err == nil {
		return &systemDRecord{
			binPath:        binPath,
			serviceName:    serviceName,
			daemonFilePath: daemonFilePath,
			systemConfig:   systemConfig,

			wislog: wl,
		}
	}
	if _, err := os.Stat("/sbin/initctl"); err == nil {
		return &upstartRecord{
			binPath:        binPath,
			serviceName:    serviceName,
			daemonFilePath: daemonFilePath,
			systemConfig:   systemConfig,

			wislog: wl,
		}
	}

	return &systemVRecord{
		binPath:        binPath,
		serviceName:    serviceName,
		daemonFilePath: daemonFilePath,
		systemConfig:   systemConfig,

		wislog: wl,
	}
}
