package usecase

import (
	"fmt"

	"github.com/Notch-Technologies/wizy/cmd/server/config"
	"github.com/Notch-Technologies/wizy/cmd/server/database"
	"github.com/Notch-Technologies/wizy/cmd/server/pb/peer"
	"github.com/Notch-Technologies/wizy/cmd/server/repository"
	"github.com/Notch-Technologies/wizy/store"
)

type PeerUsecaseCaller interface {
	InitialSync(clientMachinePubKey, wgPublicKey string) error
}

type PeerUsecase struct {
	peerRepository    repository.PeerRepositoryCaller
	networkRepository repository.NetworkRepositoryCaller
	serverStore       *store.ServerStore
	peerServer        peer.PeerService_SyncServer
	config            *config.ServerConfig
}

func NewPeerUsecase(
	db database.SQLExecuter, config *config.ServerConfig,
	server *store.ServerStore, peerServer peer.PeerService_SyncServer,
) PeerUsecaseCaller {
	return &PeerUsecase{
		peerRepository:    repository.NewPeerRepository(db),
		networkRepository: repository.NewNetworkRepository(db),
		serverStore:       server,
		peerServer:        peerServer,
		config:            config,
	}
}

func (p *PeerUsecase) InitialSync(clientMachinePubKey, wgPublicKey string) error {
	var (
		allowedIPsFormat = "%s/%d"
	)
	
	fmt.Println("Initial Sync")
	pe, err := p.peerRepository.FindByClientMachinePubKey(clientMachinePubKey)
	if err != nil {
		fmt.Println("can not find client machine pub key")
		return err
	}

	fmt.Println("Initial Sync")
	a, err := p.peerRepository.FindByWgPubKey(wgPublicKey)
	if err != nil {
		fmt.Println("can not find wg pub key")
		return err
	}
	fmt.Println("Initial Sync")

	fmt.Println("truth admin network id")
	fmt.Println(a.AdminNetworkID)

	fmt.Println("non truth admin network id")
	fmt.Println(pe.AdminNetworkID)

	peers, err := p.peerRepository.FindPeersByAdminNetworkID(pe.AdminNetworkID)
	if err != nil {
		fmt.Println("can not find peers")
		return err
	}

	_, err = p.networkRepository.FindByNetworkID(pe.NetworkID)
	if err != nil {
		fmt.Println("can not find networks")
		return err
	}

	remotePeers := []*peer.RemotePeer{}
	for _, rPeer := range peers {
		if pe.WgPubKey != rPeer.WgPubKey {
			remotePeers = append(remotePeers, &peer.RemotePeer{
				WgPubKey:   rPeer.WgPubKey,
				AllowedIps: []string{fmt.Sprintf(allowedIPsFormat, rPeer.IP, rPeer.CIDR)},
			})
		}
	}

	err = p.peerServer.Send(&peer.SyncResponse{
		PeerConfig:        &peer.PeerConfig{Address: pe.IP},
		RemotePeers:       remotePeers,
		RemotePeerIsEmpty: len(remotePeers) == 0,
		StunTurnConfig:    p.createStunTurnConfig(),
	})
	if err != nil {
		return err
	}

	return nil
}

func (p *PeerUsecase) createStunTurnConfig() *peer.StunTurnConfig {
	var (
		stuns []*peer.Host
		turns []*peer.Host
	)

	for _, stun := range p.config.Stuns {
		peer := &peer.Host{
			Url:      stun.URL,
			Username: *stun.Username,
			Password: *stun.Password,
		}
		stuns = append(stuns, peer)

	}

	for _, turn := range p.config.TURNConfig.Turns {
		peer := &peer.Host{
			Url:      turn.URL,
			Username: *turn.Username,
			Password: *turn.Password,
		}
		turns = append(turns, peer)
	}

	return &peer.StunTurnConfig{
		Stuns: stuns,
		TurnCredentials: &peer.TurnCredential{
			Turns:                turns,
			CredentialsTTL:       p.config.TURNConfig.CredentialsTTL.String(),
			Secret:               p.config.TURNConfig.Secret,
			TimeBasedCredentials: p.config.TURNConfig.TimeBasedCredentials,
		},
		Signal: &peer.Host{
			Url: p.config.Signal.URL,
		},
	}
}
