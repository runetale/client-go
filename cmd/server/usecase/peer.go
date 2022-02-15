package usecase

import (
	"fmt"

	"github.com/Notch-Technologies/wizy/cmd/server/database"
	"github.com/Notch-Technologies/wizy/cmd/server/pb/peer"
	"github.com/Notch-Technologies/wizy/cmd/server/repository"
	"github.com/Notch-Technologies/wizy/store"
)

type PeerUsecaseManger interface {
}

type PeerUsecase struct {
	peerRepository *repository.PeerRepository
	serverStore    *store.ServerStore
	peerServer     peer.PeerService_SyncServer
}

func NewPeerUsecase(
	db database.SQLExecuter,
	server *store.ServerStore,
	peerServer peer.PeerService_SyncServer,
) *PeerUsecase {
	return &PeerUsecase{
		peerRepository: repository.NewPeerRepository(db),
		serverStore:    server,
		peerServer:     peerServer,
	}
}

func (p *PeerUsecase) InitialSync(clientPubKey string) error {
	pe, err := p.peerRepository.FindByClientPubKey(clientPubKey)
	if err != nil {
		return err
	}

	peers, err := p.peerRepository.FindPeersByClientPubKey(pe.ClientPubKey)
	if err != nil {
		return err
	}

	fmt.Println(peers)

	err = p.peerServer.Send(&peer.SyncResponse{
		PeerConfig:        &peer.PeerConfig{Address: "", Dns: ""},
		RemotePeers:       []*peer.RemotePeer{},
		RemotePeerIsEmpty: true,
	})
	if err != nil {
		return err
	}

	return nil
}
