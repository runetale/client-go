package daemon

type Daemon interface {
	Install() error
	Uninstall() error

	Load() error
	Unload() error

	Start() error
	Stop() error

	Status() error

	IsInstalled() bool
	IsRunnning() (string, bool)
}

func NewDaemon(path, serviceName, plistName, plistFile string) Daemon {
	return newDaemon(path, serviceName, plistName, plistFile)
}
