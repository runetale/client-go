package daemon

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/Notch-Technologies/dotshake/dotlog"
)

type daemon struct {
	// binary path
	binPath string
	// daemon file path
	daemonFilePath string
	// daemon name
	serviceName string
	// daemon system confi
	systemConfig string

	dotlog *dotlog.DotLog
}

func newDaemon(
	binPath, serviceName, daemonFilePath, systemConfig string,
	dl *dotlog.DotLog,
) Daemon {
	return &daemon{
		binPath:        binPath,
		serviceName:    serviceName,
		daemonFilePath: daemonFilePath,
		systemConfig:   systemConfig,

		dotlog: dl,
	}
}

func (d *daemon) Install() (err error) {
	defer func() {
		if os.Getuid() != 0 && err != nil {
			d.dotlog.Logger.Errorf("run it again with sudo privileges: %s", err.Error())
			err = fmt.Errorf("run it again with sudo privileges: %s", err.Error())
		}
	}()

	err = d.Uninstall()
	if err != nil {
		return err
	}

	// seriously copy the binary
	// - create binary path => "/usr/local/bin/dotshake"
	// - execution path at build time => exeFile
	// - create tmp file => "/usr/local/bin/dotshake.tmp"
	// - copy exeFile to tmp file
	// - setting permisiion to tmpBin
	// - tmpBin to a real executable file
	//
	if err := os.MkdirAll(filepath.Dir(d.binPath), 0755); err != nil {
		d.dotlog.Logger.Errorf("failed to create %s. because %s\n", d.binPath, err.Error())
		return err
	}

	exePath, err := os.Executable()
	if err != nil {
		d.dotlog.Logger.Errorf("failed to get executablePath. because %s\n", err.Error())
		return err
	}

	tmpBin := d.binPath + ".tmp"
	f, err := os.Create(tmpBin)
	if err != nil {
		d.dotlog.Logger.Errorf("failed to create %s. because %s\n", tmpBin, err.Error())
		return err
	}

	exeFile, err := os.Open(exePath)
	if err != nil {
		f.Close()
		d.dotlog.Logger.Errorf("failed to open %s. because %s\n", exePath, err.Error())
		return err
	}

	_, err = io.Copy(f, exeFile)
	exeFile.Close()
	if err != nil {
		f.Close()
		d.dotlog.Logger.Errorf("failed to copy %s to %s. because %s\n", f, exePath, err.Error())
		return err
	}

	if err := f.Close(); err != nil {
		d.dotlog.Logger.Errorf("failed to close the %s. because %s\n", f.Name(), err.Error())
		return err
	}

	if err := os.Chmod(tmpBin, 0755); err != nil {
		d.dotlog.Logger.Errorf("failed to grant permission for %s. because %s\n", tmpBin, err.Error())
		return err
	}

	if err := os.Rename(tmpBin, d.binPath); err != nil {
		d.dotlog.Logger.Errorf("failed to rename %s to %s. because %s\n", tmpBin, d.binPath, err.Error())
		return err
	}

	if err := ioutil.WriteFile(d.daemonFilePath, []byte(d.systemConfig), 0700); err != nil {
		return err
	}

	if out, err := exec.Command("launchctl", "load", d.daemonFilePath).CombinedOutput(); err != nil {
		return fmt.Errorf("error running launchctl load %s: %v, %s", d.daemonFilePath, err, out)
	}

	if out, err := exec.Command("launchctl", "start", d.serviceName).CombinedOutput(); err != nil {
		return fmt.Errorf("error running launchctl start %s: %v, %s", d.serviceName, err, out)
	}

	return nil
}

func (d *daemon) Uninstall() (ret error) {
	plist, err := exec.Command("launchctl", "list", d.serviceName).Output()
	_ = plist
	running := err == nil

	if running {
		out, err := exec.Command("launchctl", "stop", d.serviceName).CombinedOutput()
		if err != nil {
			fmt.Printf("launchctl stop %s: %v, %s\n", d.serviceName, err, out)
			ret = err
		}
		out, err = exec.Command("launchctl", "unload", d.daemonFilePath).CombinedOutput()
		if err != nil {
			fmt.Printf("launchctl unload %s: %v, %s\n", d.daemonFilePath, err, out)
			if ret == nil {
				ret = err
			}
		}
	}

	err = os.Remove(d.daemonFilePath)
	if os.IsNotExist(err) {
		err = nil
		if ret == nil {
			ret = err
		}
	}

	err = os.Remove(d.binPath)
	if os.IsNotExist(err) {
		err = nil
		if ret == nil {
			ret = err
		}
	}
	return ret
}
