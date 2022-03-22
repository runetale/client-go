package service

import (
	"fmt"

	"github.com/Notch-Technologies/wizy/cmd/server/channel"
	"github.com/Notch-Technologies/wizy/cmd/server/config"
	"github.com/Notch-Technologies/wizy/cmd/server/database"
	"github.com/Notch-Technologies/wizy/cmd/server/pb/peer"
	"github.com/Notch-Technologies/wizy/cmd/server/usecase"
	"github.com/Notch-Technologies/wizy/store"
)

type PeerServerServiceCaller interface {
	Sync(*peer.SyncRequest, peer.PeerService_SyncServer) error
}

type PeerServerService struct {
	db                *database.Sqlite
	serverStore       *store.ServerStore
	peerUpdateManager *channel.PeersUpdateManager
	config            *config.ServerConfig

	peer.UnimplementedPeerServiceServer
}

func NewPeerServerService(
	db *database.Sqlite, config *config.ServerConfig,
	server *store.ServerStore, peerUpdateManager *channel.PeersUpdateManager,
) PeerServerServiceCaller {
	return &PeerServerService{
		db:                db,
		serverStore:       server,
		peerUpdateManager: peerUpdateManager,
		config:            config,
	}
}

func (p *PeerServerService) Sync(req *peer.SyncRequest, srv peer.PeerService_SyncServer) error {
	pu := usecase.NewPeerUsecase(p.db, p.config, p.serverStore, srv)

	err := pu.InitialSync(req.GetClientMachineKey(), req.GetWgPublicKey())
	if err != nil {
		return err
	}

	updateChannel := p.peerUpdateManager.CreateChannel(req.GetClientMachineKey())

	for {
		select {
		case update, open := <-updateChannel:
			fmt.Println("coming updatechannel")
			fmt.Println(update)
			if !open {
				fmt.Println("channel has been close")
				return nil
			}

			fmt.Printf("received an update for peer %s", req.GetClientMachineKey())

			fmt.Println("send update message")
			err = srv.SendMsg(update.Update)
			if err != nil {
				fmt.Println("failed to sending update message")
				return err
			}

			fmt.Printf("send an update to peer %s", req.GetClientMachineKey())
		case <-srv.Context().Done():
			p.peerUpdateManager.CloseChannel(req.GetClientMachineKey())
			fmt.Println("channel offline")
			return srv.Context().Err()
		}
	}
}
