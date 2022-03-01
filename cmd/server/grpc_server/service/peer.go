package service

import (
	"fmt"

	"github.com/Notch-Technologies/wizy/cmd/server/channel"
	"github.com/Notch-Technologies/wizy/cmd/server/database"
	"github.com/Notch-Technologies/wizy/cmd/server/pb/peer"
	"github.com/Notch-Technologies/wizy/cmd/server/usecase"
	"github.com/Notch-Technologies/wizy/store"
)

type PeerServiceServer struct {
	db                *database.Sqlite
	serverStore       *store.ServerStore
	peerUpdateManager *channel.PeersUpdateManager

	peer.UnimplementedPeerServiceServer
}

func NewPeerServiceServer(
	db *database.Sqlite,
	server *store.ServerStore,
	peerUpdateManager *channel.PeersUpdateManager,
) *PeerServiceServer {
	return &PeerServiceServer{
		db:                db,
		serverStore:       server,
		peerUpdateManager: peerUpdateManager,
	}
}

func (pss *PeerServiceServer) Sync(req *peer.SyncMessage, srv peer.PeerService_SyncServer) error {
	pu := usecase.NewPeerUsecase(pss.db, pss.serverStore, srv)

	err := pu.InitialSync(req.GetClientMachineKey())
	if err != nil {
		fmt.Println(err)
		return err
	}

	updateChannel := pss.peerUpdateManager.CreateChannel(req.GetClientMachineKey())

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
			pss.peerUpdateManager.CloseChannel(req.GetClientMachineKey())
			fmt.Println("channel offline")
			return srv.Context().Err()
		}
	}
}
