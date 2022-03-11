package daemon

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/Notch-Technologies/wizy/wislog"
)

type daemon struct {
	targetPath  string
	serviceName string
	plistPath   string
	plistFile   string

	wislog *wislog.WisLog
}

func newDaemon(
	targetPath, serviceName, plistPath, plistFile string,
	wl *wislog.WisLog,
) Daemon {
	return &daemon{
		targetPath:  targetPath,
		serviceName: serviceName,
		plistPath:   plistPath,
		plistFile:   plistFile,

		wislog: wl,
	}
}

// in effect, all it does is call load and start.
//
func (d *daemon) Install() (err error) {
	defer func() {
		if os.Getuid() != 0 && err != nil {
			d.wislog.Logger.Errorf("run it again with sudo privileges: %s", err.Error())
			err = fmt.Errorf("run it again with sudo privileges: %s", err.Error())
		}
	}()

	// TODO: added check privileges and is installed

	// seriously copy the binary
	// - create binary path => "/usr/local/bin/wissy"
	// - execution path at build time => exeFile
	// - create tmp file => "/usr/local/bin/wissy.tmp"
	// - copy exeFile to tmp file
	// - setting permisiion to tmpBin
	// - tmpBin to a real executable file
	// good luck
	//
	if err := os.MkdirAll(filepath.Dir(d.targetPath), 0755); err != nil {
		d.wislog.Logger.Errorf("failed to create %s. because %s\n", d.targetPath, err.Error())
		return err
	}

	exePath, err := os.Executable()
	if err != nil {
		d.wislog.Logger.Errorf("failed to get executablePath. because %s\n", err.Error())
		return err
	}

	tmpBin := d.targetPath + ".tmp"
	f, err := os.Create(tmpBin)
	if err != nil {
		d.wislog.Logger.Errorf("failed to create %s. because %s\n", tmpBin, err.Error())
		return err
	}

	exeFile, err := os.Open(exePath)
	if err != nil {
		f.Close()
		d.wislog.Logger.Errorf("failed to open %s. because %s\n", exePath, err.Error())
		return err
	}

	_, err = io.Copy(f, exeFile)
	exeFile.Close()
	if err != nil {
		f.Close()
		d.wislog.Logger.Errorf("failed to copy %s to %s. because %s\n", f, exePath, err.Error())
		return err
	}

	if err := f.Close(); err != nil {
		d.wislog.Logger.Errorf("failed to close the %s. because %s\n", f.Name(), err.Error())
		return err
	}

	if err := os.Chmod(tmpBin, 0755); err != nil {
		d.wislog.Logger.Errorf("failed to grant permission for %s. because %s\n", tmpBin, err.Error())
		return err
	}

	if err := os.Rename(tmpBin, d.targetPath); err != nil {
		d.wislog.Logger.Errorf("failed to rename %s to %s. because %s\n", tmpBin, d.targetPath, err.Error())
		return err
	}

	err = d.Uninstall()
	if err != nil {
		d.wislog.Logger.Errorf("uninstallation of %s failed. plist file is here %s. because %s\n", d.serviceName, d.plistPath, err.Error())
		return err
	}

	if err := ioutil.WriteFile(d.plistPath, []byte(d.plistFile), 0700); err != nil {
		d.wislog.Logger.Errorf("failed to write %s to %s. because %s\n", d.plistPath, d.plistFile, err.Error())
		return err
	}

	err = d.Load()
	if err != nil {
		d.wislog.Logger.Errorf("failed to load %s. plist paht is here %s. because %s\n", d.serviceName, d.plistPath, err.Error())
		return err
	}

	err = d.Start()
	if err != nil {
		d.wislog.Logger.Errorf("failed to start %s. plist path is here %s. because %s\n", d.serviceName, d.plistPath, err.Error())
		return err
	}

	return nil
}

// in effect, all it does is call unload and stop.
//
func (d *daemon) Uninstall() (err error) {
	err := d.checkPrivileges()
	if err != nil {
		return err
	}

	_, isRunnning := d.IsRunnning()
	if isRunnning {
		err := d.Stop()
		if err != nil {
			d.wislog.Logger.Errorf("failed to stop %s. plist path is here %s. because %s\n", d.serviceName, d.plistPath, err.Error())
			return err
		}
		err = d.Unload()
		if err != nil {
			d.wislog.Logger.Errorf("failed to unload %s. plist paht is here %s. because %s\n", d.serviceName, d.plistPath, err.Error())
			return err
		}
	}

	err = os.Remove(d.plistPath)
	if os.IsNotExist(err) {
		return nil
	}

	return err
}

func (d *daemon) Load() error {
	err := d.checkPrivileges()
	if err != nil {
		return err
	}

	if out, err := exec.Command("launchctl", "load", d.plistPath).CombinedOutput(); err != nil {
		fmt.Printf("failed to running launchctl load %s, because %s\n %s\n", d.plistPath, err.Error(), out)
		return err
	}

	return nil
}

func (d *daemon) Unload() error {
	err := d.checkPrivileges()
	if err != nil {
		return err
	}

	if !d.IsInstalled() {
		return errors.New("not installed")
	}

	if _, isRunning := d.IsRunnning(); !isRunning {
		return errors.New("not running")
	}

	out, err := exec.Command("launchctl", "unload", d.serviceName).CombinedOutput()
	if err != nil {
		fmt.Printf("failed to launchctl unload %s, because %v.\n %s\n", d.serviceName, err.Error(), out)
		return err
	}

	return nil
}

func (d *daemon) Start() error {
	err := d.checkPrivileges()
	if err != nil {
		return err
	}

	if _, isRunning := d.IsRunnning(); !isRunning {
		return errors.New("already running")
	}

	if out, err := exec.Command("launchctl", "start", d.serviceName).CombinedOutput(); err != nil {
		fmt.Printf("failed to running launchctl start %s, because %s.\n %s\n", d.serviceName, err.Error(), out)
		return err
	}
	return nil
}

func (d *daemon) Stop() error {
	err := d.checkPrivileges()
	if err != nil {
		return err
	}

	if !d.IsInstalled() {
		return errors.New("not installed")
	}

	if _, isRunning := d.IsRunnning(); !isRunning {
		return errors.New("not running")
	}

	out, err := exec.Command("launchctl", "stop", d.serviceName).CombinedOutput()
	if err != nil {
		fmt.Printf("failed to launchctl stop %s, because %v.\n %s\n", d.serviceName, err.Error(), out)
		return err
	}
	return nil
}

func (d *daemon) Status() error {
	err := d.checkPrivileges()
	if err != nil {
		return err
	}

	if !d.IsInstalled() {
		return fmt.Errorf("%s is not installed. please call the install command and try again", d.serviceName)
	}

	mes, isRunnning := d.IsRunnning()
	if !isRunnning {
		fmt.Println(mes)
		return err
	}

	fmt.Println(mes)
	return nil
}

func (d *daemon) IsInstalled() bool {
	if _, err := os.Stat(d.plistPath); err == nil {
		return true
	}
	return false
}

func (d *daemon) IsRunnning() (string, bool) {
	out, err := exec.Command("launchctl", "list", d.serviceName).CombinedOutput()
	if err == nil {
		if matched, err := regexp.MatchString(d.serviceName, string(out)); err == nil && matched {
			reg := regexp.MustCompile("PID\" = ([0-9]+);")
			data := reg.FindStringSubmatch(string(out))
			if len(data) > 1 {
				return fmt.Sprintf("%s is running on pid: %s", d.serviceName, data[1]), true
			}
			return fmt.Sprintf("%s is running. but cannot get pid. please report it", d.serviceName), false
		}
	}

	return fmt.Sprintf("%s is stopped", d.serviceName), false
}

func (d *daemon) checkPrivileges() error {
	if out, err := exec.Command("id", "-g").CombinedOutput(); err == nil {
		if gid, parseErr := strconv.ParseUint(strings.TrimSpace(string(out)), 10, 32); parseErr == nil {
			if gid == 0 {
				return nil
			}
			return errors.New("run with root privileges")
		}
	}
	return errors.New("unsupport system")
}
