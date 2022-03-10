package daemon

import (
	"os"

	"github.com/Notch-Technologies/wizy/wislog"
)

func newDaemon(
	targetPath, serviceName, plistPath, plistFile string,
	wl *wislog.WisLog,
) Daemon {
	if _, err := os.Stat("/run/systemd/system"); err == nil {
		return &systemDRecord{
			targetPath:  targetPath,
			serviceName: serviceName,
			plistPath:   plistPath,
			plistFile:   plistFile,
    	
			wislog: wl,
		}
	}
	if _, err := os.Stat("/sbin/initctl"); err == nil {
		return &upstartRecord{
			targetPath:  targetPath,
			serviceName: serviceName,
			plistPath:   plistPath,
			plistFile:   plistFile,
    	
			wislog: wl,
		}
	}

	return &systemVRecord{
		targetPath:  targetPath,
		serviceName: serviceName,
		plistPath:   plistPath,
		plistFile:   plistFile,

		wislog: wl,
	}
}
