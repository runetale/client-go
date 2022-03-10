package daemon

import (
	"fmt"
	"os"

	"github.com/Notch-Technologies/wizy/wislog"
)

type systemDRecord struct {
	targetPath  string
	serviceName string
	plistPath   string
	plistFile   string

	wislog *wislog.WisLog
}

func (d *systemDRecord) Install() (err error) {
	defer func() {
		if os.Getuid() != 0 && err != nil {
			d.wislog.Logger.Errorf("run it again with sudo privileges: %s", err.Error())
			err = fmt.Errorf("run it again with sudo privileges: %s", err.Error())
		}
	}()

	return nil
}

func (d *systemDRecord) Uninstall() error {
	return nil
}

func (d *systemDRecord) Load() error {
	return nil
}

func (d *systemDRecord) Unload() error {
	return nil
}

func (d *systemDRecord) Start() error {
	return nil
}

func (d *systemDRecord) Stop() error {
	return nil
}

func (d *systemDRecord) Status() error {
	return nil
}

func (d *systemDRecord) IsInstalled() bool {
	return false
}

func (d *systemDRecord) IsRunnning() (string, bool) {
	return "", false
}
