package grpcserver

import (
	client "github.com/Notch-Technologies/wizy/auth0"
	"github.com/Notch-Technologies/wizy/cmd/server/channel"
	"github.com/Notch-Technologies/wizy/cmd/server/config"
	"github.com/Notch-Technologies/wizy/cmd/server/database"
	"github.com/Notch-Technologies/wizy/cmd/server/server/service"
	"github.com/Notch-Technologies/wizy/store"
)

type Server struct {
	UserServerService         service.UserServerServiceCaller
	PeerServerService         service.PeerServerServiceCaller
	SessionServerService      service.SessionServiceServerCaller
	OrganizationServerService service.OrganizationServerServiceCaller
}

func NewServer(
	db *database.Sqlite, config *config.ServerConfig,
	server *store.ServerStore, client *client.Auth0Client,
	peerUpdateManager *channel.PeersUpdateManager,
) *Server {
	return &Server{
		UserServerService:         service.NewUserServerService(db),
		PeerServerService:         service.NewPeerServerService(db, server, peerUpdateManager),
		SessionServerService:      service.NewSessionServerService(db, config, server, peerUpdateManager),
		OrganizationServerService: service.NewOrganizationServerService(db, client),
	}
}
