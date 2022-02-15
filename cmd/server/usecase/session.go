package usecase

import (
	"errors"

	"github.com/Notch-Technologies/wizy/cmd/server/database"
	"github.com/Notch-Technologies/wizy/cmd/server/domain"
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
}

func NewSessionUsecase(
	db database.SQLExecuter,
	server *store.ServerStore,
) *SessionUsecase {
	return &SessionUsecase{
		setupKeyRepository: repository.NewSetupKeyRepository(db),
		userRepository:     repository.NewUserRepository(db),
		peerRepository:     repository.NewPeerRepository(db),
		serverStore:        server,
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

	peer, err := s.peerRepository.FindBySetupKeyID(sk.ID, clientPubKey)
	if err != nil {
		if errors.Is(err, domain.ErrNoRows) {
			peer := domain.NewPeer(sk.ID, user.NetworkID, user.UserGroupID, user.ID, user.OrganizationID, "", clientPubKey)
			err = s.peerRepository.CreatePeer(peer)
			if err != nil {
				return nil, err
			}
		}
		return nil, err
	}

	return peer, nil
}
