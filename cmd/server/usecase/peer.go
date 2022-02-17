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
	networkRepository *repository.NetworkRepository
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
		networkRepository: repository.NewNetworkRepository(db),
		serverStore:    server,
		peerServer:     peerServer,
	}
}

const AllowedIPsFormat = "%s/32"

func (p *PeerUsecase) InitialSync(clientPubKey string) error {
	pe, err := p.peerRepository.FindByClientPubKey(clientPubKey)
	if err != nil {
		fmt.Println("can not find pub key")
		return err
	}

	peers, err := p.peerRepository.FindPeersByClientPubKey(pe.ClientPubKey)
	if err != nil {
		fmt.Println("can not find peers")
		return err
	}

	network, err := p.networkRepository.FindByNetworkID(pe.NetworkID)
	if err != nil {
		fmt.Println("can not find networks")
		return err
	}

	fmt.Println("your ip")
	fmt.Println(network.IP)

	fmt.Println("Initial Sync Peers")
	fmt.Println(peers)

	remotePeers := []*peer.RemotePeer{}
	for _, rPeer := range peers {
		remotePeers = append(remotePeers, &peer.RemotePeer{
			WgPubKey:   rPeer.ClientPubKey,
			AllowedIps: []string{fmt.Sprintf(AllowedIPsFormat, "10.0.0.1")}, //todo /32
		})
	}
	err = p.peerServer.Send(&peer.SyncResponse{
		PeerConfig:        &peer.PeerConfig{Address: "", Dns: ""},
		RemotePeers:       remotePeers,
		RemotePeerIsEmpty: true,
	})
	if err != nil {
		return err
	}

	return nil
}
