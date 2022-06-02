package webrtc

type Credentials struct {
	UserName string
	Pwd      string
}

func NewCredentials(uname, pwd string) *Credentials {
	return &Credentials{
		UserName: uname,
		Pwd:      pwd,
	}
}
