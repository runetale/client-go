package grpcserver

import (
	"github.com/Notch-Technologies/wizy/client"
	"github.com/Notch-Technologies/wizy/cmd/server/config"
	"github.com/Notch-Technologies/wizy/cmd/server/database"
	"github.com/Notch-Technologies/wizy/cmd/server/grpc_server/service"
	"github.com/Notch-Technologies/wizy/store"
)

type Server struct {
	UserServiceServer    *service.UserServiceServer
	PeerServiceServer    *service.PeerServiceServer
	SessionServiceServer *service.SessionServiceServer
	OrganizationServiceServer *service.OrganizationServiceServer
}

func NewServer(
	db *database.Sqlite, config *config.Config,
	server *store.ServerStore, client *client.Auth0Client,
) (*Server, error) {
	return &Server{
		UserServiceServer:    service.NewUserServiceServer(db),
		PeerServiceServer:    service.NewPeerServiceServer(db),
		SessionServiceServer: service.NewSessionServiceServer(db, config, server),
		OrganizationServiceServer: service.NewOrganizationServiceServer(db, client),
	}, nil
}
