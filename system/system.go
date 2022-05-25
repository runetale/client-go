package system

type SysInfo struct {
	GoOS      string
	Kernel    string
	Core      string
	Platform  string
	OS        string
	OSVersion string
	Hostname  string
	CPUs      int
	Version   string
}

func NewSysInfo() *SysInfo {
	return GetInfo()
}
