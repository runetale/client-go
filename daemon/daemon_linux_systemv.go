package daemon

import (
	"fmt"
	"os"

	"github.com/Notch-Technologies/dotshake/dotlog"
)

type systemVRecord struct {
	// binary path
	binPath string
	// daemon name
	serviceName string
	// daemon file path
	daemonFilePath string
	// daemon system config
	systemConfig string

	dotlog *dotlog.DotLog
}

func (d *systemVRecord) Install() (err error) {
	defer func() {
		if os.Getuid() != 0 && err != nil {
			d.dotlog.Logger.Errorf("run it again with sudo privileges: %s", err.Error())
			err = fmt.Errorf("run it again with sudo privileges: %s", err.Error())
		}
	}()

	return nil
}

func (d *systemVRecord) Uninstall() error {
	return nil
}

func (d *systemVRecord) Load() error {
	return nil
}

func (d *systemVRecord) Unload() error {
	return nil
}

func (d *systemVRecord) Start() error {
	return nil
}

func (d *systemVRecord) Stop() error {
	return nil
}

func (d *systemVRecord) Status() error {
	return nil
}

func (d *systemVRecord) IsInstalled() bool {
	return false
}

func (d *systemVRecord) IsRunnning() (string, bool) {
	return "", false
}
