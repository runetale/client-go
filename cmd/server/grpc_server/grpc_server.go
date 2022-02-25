package grpcserver

import (
	"github.com/Notch-Technologies/wizy/client"
	"github.com/Notch-Technologies/wizy/cmd/server/config"
	"github.com/Notch-Technologies/wizy/cmd/server/database"
	"github.com/Notch-Technologies/wizy/cmd/server/grpc_server/service"
	"github.com/Notch-Technologies/wizy/store"
	"github.com/Notch-Technologies/wizy/cmd/server/channel"
)

type Server struct {
	UserServiceServer         *service.UserServiceServer
	PeerServiceServer         *service.PeerServiceServer
	SessionServiceServer      *service.SessionServiceServer
	OrganizationServiceServer *service.OrganizationServiceServer
	NegotiationServer         *service.NegotiationServiceServer
}

func NewServer(
	db *database.Sqlite, config *config.Config,
	server *store.ServerStore, client *client.Auth0Client,
	peerUpdateManager *channel.PeersUpdateManager,
) (*Server, error) {
	return &Server{
		UserServiceServer:         service.NewUserServiceServer(db),
		PeerServiceServer:         service.NewPeerServiceServer(db, server, peerUpdateManager),
		SessionServiceServer:      service.NewSessionServiceServer(db, config, server, peerUpdateManager),
		OrganizationServiceServer: service.NewOrganizationServiceServer(db, client),
		NegotiationServer:         service.NewNegotiationServiceServer(db),
	}, nil
}
