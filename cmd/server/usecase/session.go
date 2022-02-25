package usecase

import (
	"errors"
	"fmt"

	"github.com/Notch-Technologies/wizy/cmd/server/channel"
	"github.com/Notch-Technologies/wizy/cmd/server/database"
	"github.com/Notch-Technologies/wizy/cmd/server/domain"
	"github.com/Notch-Technologies/wizy/cmd/server/pb/peer"
	"github.com/Notch-Technologies/wizy/cmd/server/repository"
	"github.com/Notch-Technologies/wizy/store"
)

type SessionUsecaseManager interface {
	CreatePeer(setupKey, clientPubKey, serverPubKey string) (*domain.Peer, error)
}

type SessionUsecase struct {
	setupKeyRepository *repository.SetupKeyRepository
	userRepository     *repository.UserRepository
	peerRepository     *repository.PeerRepository
	serverStore        *store.ServerStore
	peerUpdateManager *channel.PeersUpdateManager
}

func NewSessionUsecase(
	db database.SQLExecuter,
	server *store.ServerStore,
	peerUpdateManager *channel.PeersUpdateManager,
) *SessionUsecase {
	return &SessionUsecase{
		setupKeyRepository: repository.NewSetupKeyRepository(db),
		userRepository:     repository.NewUserRepository(db),
		peerRepository:     repository.NewPeerRepository(db),
		serverStore:        server,
		peerUpdateManager: peerUpdateManager,
	}
}

func (s *SessionUsecase) CreatePeer(setupKey, clientPubKey, serverPubKey string) (*domain.Peer, error) {
	if s.serverStore.GetPublicKey() != serverPubKey {
		return nil, errors.New(domain.ErrInvalidPublicKey.Error())
	}

	sk, err := s.setupKeyRepository.FindBySetupKey(setupKey)
	if err != nil {
		return nil, err
	}

	user, err := s.userRepository.FindByUserID(sk.UserID)
	if err != nil {
		return nil, err
	}

	pe, err := s.peerRepository.FindBySetupKeyID(sk.ID, clientPubKey)
	if err != nil {
		if errors.Is(err, domain.ErrNoRows) {
			newPeer := domain.NewPeer(sk.ID, user.NetworkID, user.UserGroupID, user.ID, user.OrganizationID, "", clientPubKey)
			err = s.peerRepository.CreatePeer(newPeer)
			if err != nil {
				return nil, err
			}
	
			peers, err := s.peerRepository.FindByOrganizationID(user.OrganizationID)
			if err != nil {
				return nil, err
			}

			peersToSend := []*peer.RemotePeer{}
			for _, remotePeer := range peers {
				for _, p := range peers {
					if remotePeer.ClientPubKey != p.ClientPubKey {
						peersToSend = append(peersToSend, &peer.RemotePeer{
							WgPubKey:   p.ClientPubKey,
							AllowedIps: []string{fmt.Sprintf(AllowedIPsFormat, "10.0.0.1"), fmt.Sprintf(AllowedIPsFormat, "10.0.0.2")}, //todo /32
						})
					}
				}
    		
				err := s.peerUpdateManager.SendUpdate(remotePeer.ClientPubKey, &channel.UpdateMessage{Update: &peer.SyncResponse{
					PeerConfig:        &peer.PeerConfig{Address: "", Dns: ""},
					RemotePeers:       peersToSend,
					RemotePeerIsEmpty: len(peersToSend) == 0,
				}})
				if err != nil {
					return nil, err
				}
				fmt.Println("Sending Update From Create Peer")
			}
			
		}
		return nil, err
	}

	return pe, nil
} 
