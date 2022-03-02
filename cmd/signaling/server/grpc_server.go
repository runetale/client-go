package server

import (
	"github.com/Notch-Technologies/wizy/cmd/signaling/server/service"
	"github.com/Notch-Technologies/wizy/wislog"
)

type Server struct {
	NegotiationServer *service.NegotiationServerService

	wislog *wislog.WisLog
}

func NewServer(wl *wislog.WisLog) *Server {
	return &Server{
		NegotiationServer: service.NewNegotiationServerService(),
		wislog:            wl,
	}
}
