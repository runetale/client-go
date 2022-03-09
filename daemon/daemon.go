package daemon

type Daemon interface {
	Install() error
	Uninstall() error

	Start() error
	Stop() error
}

func NewDaemon(path string, serviceName string, plistName string) Daemon {
	return newDaemon()
}
