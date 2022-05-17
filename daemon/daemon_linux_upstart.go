package daemon

import (
	"fmt"
	"os"

	"github.com/Notch-Technologies/dotshake/wislog"
)

type upstartRecord struct {
	// binary path
	binPath string
	// daemon name
	serviceName string
	// daemon file path
	daemonFilePath string
	// daemon system config
	systemConfig string

	wislog *wislog.WisLog
}

func (d *upstartRecord) Install() (err error) {
	defer func() {
		if os.Getuid() != 0 && err != nil {
			d.wislog.Logger.Errorf("run it again with sudo privileges: %s", err.Error())
			err = fmt.Errorf("run it again with sudo privileges: %s", err.Error())
		}
	}()

	return nil
}

func (d *upstartRecord) Uninstall() error {
	return nil
}

func (d *upstartRecord) Load() error {
	return nil
}

func (d *upstartRecord) Unload() error {
	return nil
}

func (d *upstartRecord) Start() error {
	return nil
}

func (d *upstartRecord) Stop() error {
	return nil
}

func (d *upstartRecord) Status() error {
	return nil
}

func (d *upstartRecord) IsInstalled() bool {
	return false
}

func (d *upstartRecord) IsRunnning() (string, bool) {
	return "", false
}
