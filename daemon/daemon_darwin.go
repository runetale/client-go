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
)

type daemon struct {
	targetPath  string
	serviceName string
	plistPath   string
	plistFile   string
}

func newDaemon(path, serviceName, plistPath, plistFile string) *daemon {
	return &daemon{
		targetPath:  path,
		serviceName: serviceName,
		plistPath:   plistPath,
		plistFile:   plistFile,
	}
}

// in effect, all it does is call load and start.
//
func (d *daemon) Install() (err error) {
	defer func() {
		if os.Getuid() != 0 && err != nil {
			err = fmt.Errorf("run it again with sudo privileges: %s", err.Error())
		}
	}()

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
		return err
	}

	exePath, err := os.Executable()
	if err != nil {
		return err
	}

	tmpBin := d.targetPath + ".tmp"
	f, err := os.Create(tmpBin)
	if err != nil {
		return err
	}

	exeFile, err := os.Open(exePath)
	if err != nil {
		f.Close()
		return err
	}

	_, err = io.Copy(f, exeFile)
	exeFile.Close()
	if err != nil {
		f.Close()
		return err
	}

	if err := f.Close(); err != nil {
		return err
	}

	if err := os.Chmod(tmpBin, 0755); err != nil {
		return err
	}

	if err := os.Rename(tmpBin, d.targetPath); err != nil {
		return err
	}

	err = d.Uninstall()
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(d.plistPath, []byte(d.plistFile), 0700); err != nil {
		return err
	}

	err = d.Load()
	if err != nil {
		return err
	}

	err = d.Start()
	if err != nil {
		return err
	}

	return nil
}

// in effect, all it does is call unload and stop.
//
func (d *daemon) Uninstall() (err error) {
	_, isRunnning := d.IsRunnning()
	if isRunnning {
		err := d.Stop()
		if err != nil {
			return err
		}
		err = d.Unload()
		if err != nil {
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
	if out, err := exec.Command("launchctl", "load", d.plistPath).CombinedOutput(); err != nil {
		fmt.Printf("failed to running launchctl load %s, because %s\n %s\n", d.plistPath, err.Error(), out)
		return err
	}
	return nil
}

func (d *daemon) Unload() error {
	out, err := exec.Command("launchctl", "unload", d.serviceName).CombinedOutput()
	if err != nil {
		fmt.Printf("failed to launchctl unload %s, because %v.\n %s\n", d.serviceName, err.Error(), out)
		return err
	}
	return nil
}

func (d *daemon) Start() error {
	if out, err := exec.Command("launchctl", "start", d.serviceName).CombinedOutput(); err != nil {
		fmt.Printf("failed to running launchctl start %s, because %s.\n %s\n", d.serviceName, err.Error(), out)
		return err
	}
	return nil
}

func (d *daemon) Stop() error {
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

	fmt.Print(mes)
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
			return fmt.Sprintf("%s is running. but cannot get pid. please report it", d.serviceName, data[1]), false
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
