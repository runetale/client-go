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
	peerUpdateManager  *channel.PeersUpdateManager
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
		peerUpdateManager:  peerUpdateManager,
	}
}

func (s *SessionUsecase) CreatePeer(setupKey, clientMachinePubKey, serverMachinePubKey, wgPubKey, ip string) (*domain.Peer, error) {
	if s.serverStore.GetPublicKey() != serverMachinePubKey {
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

	pe, err := s.peerRepository.FindBySetupKeyID(sk.ID, clientMachinePubKey)
	if err != nil {
		if errors.Is(err, domain.ErrNoRows) {
			newPeer := domain.NewPeer(
				sk.ID,
				user.NetworkID,
				user.UserGroupID,
				user.ID,
				user.OrganizationID,
				ip,
				clientMachinePubKey,
				wgPubKey,
			)

			err = s.peerRepository.CreatePeer(newPeer)
			if err != nil {
				return nil, err
			}

			peers, err := s.peerRepository.FindByOrganizationID(user.OrganizationID)
			if err != nil {
				return nil, err
			}

			for _, remotePeer := range peers {
				peersToSend := []*peer.RemotePeer{}
				for _, p := range peers {
					if remotePeer.WgPubKey != p.WgPubKey {
						peersToSend = append(peersToSend, &peer.RemotePeer{
							WgPubKey:   p.ClientPubKey,
							AllowedIps: []string{fmt.Sprintf(AllowedIPsFormat, p.IP)},
						})
					}
				}

				fmt.Printf("send peer information to the %s update channel\n", remotePeer.ClientPubKey)
				fmt.Println(peersToSend)
				err := s.peerUpdateManager.SendUpdate(remotePeer.ClientPubKey, &channel.UpdateMessage{Update: &peer.SyncResponse{
					PeerConfig:        &peer.PeerConfig{Address: newPeer.IP},
					RemotePeers:       peersToSend,
					RemotePeerIsEmpty: len(peersToSend) == 0,
				}})
				if err != nil {
					return nil, err
				}
				fmt.Println("send updates that will be sent upon initial Peer registration")
			}

		}
		return nil, err
	}

	return pe, nil
}
