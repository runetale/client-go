package engine

import (
	"github.com/Notch-Technologies/wizy/wislog"
)

type Engine struct {
	wislog *wislog.WisLog
	Peer *Peer
}

func NewEngine(peer *Peer, log *wislog.WisLog) *Engine {
	return &Engine{
		Peer: peer,
		wislog: log,
	}
}
