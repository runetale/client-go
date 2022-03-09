package daemon

type daemon struct {
	binPath string
	serviceName string
	plistName string
}

func newDaemon() *daemon {
	return &daemon{}
}

func (d *daemon) Install() error {
	return nil
}

func (d *daemon) Uninstall() error {
	return nil
}

func (d *daemon) Start() error {
	return nil
}

func (d *daemon) Stop() error {
	return nil
}
