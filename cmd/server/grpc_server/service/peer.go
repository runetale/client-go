package service

import (
	"fmt"
	"sync"

	"github.com/Notch-Technologies/wizy/cmd/server/database"
	"github.com/Notch-Technologies/wizy/cmd/server/pb/peer"
	"github.com/Notch-Technologies/wizy/cmd/server/usecase"
	"github.com/Notch-Technologies/wizy/store"
)

type UpdateMessage struct {
	Update *peer.SyncResponse
}

type PeersUpdateManager struct {
	peerChannels map[string]chan *UpdateMessage
	channelsMux  *sync.Mutex
}

// NewPeersUpdateManager returns a new instance of PeersUpdateManager
func NewPeersUpdateManager() *PeersUpdateManager {
	return &PeersUpdateManager{
		peerChannels: make(map[string]chan *UpdateMessage),
		channelsMux:  &sync.Mutex{},
	}
}

// SendUpdate sends update message to the peer's channel
func (p *PeersUpdateManager) SendUpdate(peer string, update *UpdateMessage) error {
	p.channelsMux.Lock()
	defer p.channelsMux.Unlock()
	if channel, ok := p.peerChannels[peer]; ok {
		channel <- update
		return nil
	}
	fmt.Printf("peer %s has no channel", peer)
	return nil
}

func (p *PeersUpdateManager) CreateChannel(peerKey string) chan *UpdateMessage {
	p.channelsMux.Lock()
	defer p.channelsMux.Unlock()

	if channel, ok := p.peerChannels[peerKey]; ok {
		delete(p.peerChannels, peerKey)
		close(channel)
	}
	//mbragin: todo shouldn't it be more? or configurable?
	channel := make(chan *UpdateMessage, 100)
	p.peerChannels[peerKey] = channel

	fmt.Printf("opened updates channel for a peer %s", peerKey)
	return channel
}

func (p *PeersUpdateManager) CloseChannel(peerKey string) {
	p.channelsMux.Lock()
	defer p.channelsMux.Unlock()
	if channel, ok := p.peerChannels[peerKey]; ok {
		delete(p.peerChannels, peerKey)
		close(channel)
	}

	fmt.Printf("closed updates channel of a peer %s", peerKey)
}


type PeerServiceServer struct {
	db *database.Sqlite
	serverStore *store.ServerStore
	peerUpdateManager *PeersUpdateManager

	peer.UnimplementedPeerServiceServer
}

func NewPeerServiceServer(
	db *database.Sqlite,
	server *store.ServerStore,
) *PeerServiceServer {
	return &PeerServiceServer{
		db: db,
		serverStore: server,
		peerUpdateManager: NewPeersUpdateManager(),
	}
}

func (pss *PeerServiceServer) Sync(req *peer.SyncMessage, srv peer.PeerService_SyncServer) error {
	pu := usecase.NewPeerUsecase(pss.db, pss.serverStore, srv)

	// rename client machine key -> clinet pub key
	err := pu.InitialSync(req.GetClientMachineKey())
	if err != nil {
		return err
	}

	updateChannel := pss.peerUpdateManager.CreateChannel(req.GetClientMachineKey())
	
	// TODO: create channel for connectivity state
	fmt.Println("channel online")

	// TODO: separate other package
	for {
		select {
		case update, open := <-updateChannel:
			if !open {
				fmt.Println("channel has been close")
				return nil
			}
			
			fmt.Printf("received an update for peer %s", req.GetClientMachineKey())

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
