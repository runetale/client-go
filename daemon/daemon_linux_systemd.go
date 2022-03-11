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

type systemDRecord struct {
	// binary path
	targetPath  string
	// daemon name
	serviceName string
	// daemon system confi
	systemConfig string

	wislog *wislog.WisLog
}

// in effect, all it does is call load and start.
//
func (s *systemDRecord) Install() (err error) {
	defer func() {
		if os.Getuid() != 0 && err != nil {
			s.wislog.Logger.Errorf("run it again with sudo privileges: %s", err.Error())
			err = fmt.Errorf("run it again with sudo privileges: %s", err.Error())
		}
	}()

	err = s.checkPrivileges()
	if err != nil {
		return err
	}

	if s.IsInstalled() {
		return nil
	}

	if err := os.MkdirAll(filepath.Dir(s.targetPath), 0755); err != nil {
		s.wislog.Logger.Errorf("failed to create %s. because %s\n", s.targetPath, err.Error())
		return err
	}

	exePath, err := os.Executable()
	if err != nil {
		s.wislog.Logger.Errorf("failed to get executablePath. because %s\n", err.Error())
		return err
	}

	tmpBin := s.targetPath + ".tmp"
	f, err := os.Create(tmpBin)
	if err != nil {
		s.wislog.Logger.Errorf("failed to create %s. because %s\n", tmpBin, err.Error())
		return err
	}

	exeFile, err := os.Open(exePath)
	if err != nil {
		f.Close()
		s.wislog.Logger.Errorf("failed to open %s. because %s\n", exePath, err.Error())
		return err
	}

	_, err = io.Copy(f, exeFile)
	exeFile.Close()
	if err != nil {
		f.Close()
		s.wislog.Logger.Errorf("failed to copy %s to %s. because %s\n", f, exePath, err.Error())
		return err
	}

	if err := f.Close(); err != nil {
		s.wislog.Logger.Errorf("failed to close the %s. because %s\n", f.Name(), err.Error())
		return err
	}

	if err := os.Chmod(tmpBin, 0755); err != nil {
		s.wislog.Logger.Errorf("failed to grant permission for %s. because %s\n", tmpBin, err.Error())
		return err
	}

	if err := os.Rename(tmpBin, s.targetPath); err != nil {
		s.wislog.Logger.Errorf("failed to rename %s to %s. because %s\n", tmpBin, s.targetPath, err.Error())
		return err
	}

	err = s.Uninstall()
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(s.targetPath, []byte(s.systemConfig), 0700); err != nil {
		s.wislog.Logger.Errorf("failed to write %s to %s. because %s\n", s.targetPath, s.systemConfig, err.Error())
		return err
	}

	err = s.Load()
	if err != nil {
		return err
	}

	err = s.Start()
	if err != nil {
		return err
	}

	return nil
}

// in effect, all it does is call unload and stop.
//
func (s *systemDRecord) Uninstall() error {
	err := s.checkPrivileges()
	if err != nil {
		return err
	}

	_, isRunnning := s.IsRunnning()
	if isRunnning {
		err := s.Stop()
		if err != nil {
			s.wislog.Logger.Errorf("failed to stop %s. path is here %s. because %s\n", s.serviceName, s.targetPath, err.Error())
			return err
		}
		err = s.Unload()
		if err != nil {
			s.wislog.Logger.Errorf("failed to disable %s. path is here %s. because %s\n", s.serviceName, s.targetPath, err.Error())
			return err
		}
	}

	err = os.Remove(s.targetPath)
	if os.IsNotExist(err) {
		return nil
	}

	return err
}

func (s *systemDRecord) Load() error {
	err := s.checkPrivileges()
	if err != nil {
		return err
	}

	if out, err := exec.Command("systemctl", "daemon-reload").CombinedOutput(); err != nil {
		fmt.Printf("failed to running systemctl daemon-reload %s, because %s\n %s\n", s.targetPath, err.Error(), out)
		return err
	}

	if out, err := exec.Command("systemctl", "enable", s.serviceName + ".service").CombinedOutput(); err != nil {
		fmt.Printf("failed to running systemctl daemon-reload %s, because %s\n %s\n", s.targetPath, err.Error(), out)
		return err
	}

	return nil
}

func (s *systemDRecord) Unload() error {
	err := s.checkPrivileges()
	if err != nil {
		return err
	}

	if !s.IsInstalled() {
		return errors.New("not installed")
	}

	if _, isRunning := s.IsRunnning(); !isRunning {
		return errors.New("not running")
	}

	if out, err := exec.Command("systemctl", "disable", s.serviceName + ".service").CombinedOutput(); err != nil {
		fmt.Printf("failed to disable systemctl %s, because %s\n %s\n", s.targetPath, err.Error(), out)
		return err
	}

	return nil
}

func (s *systemDRecord) Start() error {
	err := s.checkPrivileges()
	if err != nil {
		return err
	}

	if _, isRunning := s.IsRunnning(); !isRunning {
		return errors.New("already running")
	}

	if out, err := exec.Command("systemctl", "start", s.serviceName + ".service").CombinedOutput(); err != nil {
		fmt.Printf("failed to running systemctl daemon-reload %s, because %s\n %s\n", s.targetPath, err.Error(), out)
		return err
	}

	return nil
}

func (s *systemDRecord) Stop() error {
	err := s.checkPrivileges()
	if err != nil {
		return err
	}

	if !s.IsInstalled() {
		return errors.New("not installed")
	}

	if _, isRunning := s.IsRunnning(); !isRunning {
		return errors.New("not running")
	}

	if out, err := exec.Command("systemctl", "stop", s.serviceName + ".service").CombinedOutput(); err != nil {
		fmt.Printf("failed to stop systemctl %s, because %s\n %s\n", s.targetPath, err.Error(), out)
		return err
	}

	return nil
}

func (s *systemDRecord) Status() error {
	err := s.checkPrivileges()
	if err != nil {
		return err
	}

	if !s.IsInstalled() {
		return fmt.Errorf("%s is not installed. please call the install command and try again", s.serviceName)
	}

	mes, isRunnning := s.IsRunnning()
	if !isRunnning {
		fmt.Println(mes)
		return err
	}

	fmt.Println(mes)
	return nil

	return nil
}

func (s *systemDRecord) IsInstalled() bool {
	if _, err := os.Stat(s.targetPath); err == nil {
		return true
	}

	return false
}

func (s *systemDRecord) IsRunnning() (string, bool) {
	output, err := exec.Command("systemctl", "status", s.serviceName + ".service").Output()
	if err == nil {
	    if matched, err := regexp.MatchString("Active: active", string(output)); err == nil && matched {
	        reg := regexp.MustCompile("Main PID: ([0-9]+)")
	        data := reg.FindStringSubmatch(string(output))
	        if len(data) > 1 {
				return fmt.Sprintf("%s is running on pid: %s", s.serviceName, data[1]), true
	        }
			return fmt.Sprintf("%s is running. but cannot get pid. please report it", s.serviceName), false
	    }
	}

	return fmt.Sprintf("%s is stopped", s.serviceName), false
}

func (s *systemDRecord) checkPrivileges() error {
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
