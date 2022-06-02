package wireguard

import "time"

const (
	WgPort             = 51820
	DefaultMTU         = 1280
	DefaultWgKeepAlive = 25 * time.Second
)
