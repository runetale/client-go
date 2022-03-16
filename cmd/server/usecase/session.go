package usecase

import (
	"errors"
	"fmt"
	"net"

	"github.com/Notch-Technologies/wizy/cmd/server/channel"
	"github.com/Notch-Technologies/wizy/cmd/server/config"
	"github.com/Notch-Technologies/wizy/cmd/server/database"
	"github.com/Notch-Technologies/wizy/cmd/server/domain"
	"github.com/Notch-Technologies/wizy/cmd/server/ip"
	"github.com/Notch-Technologies/wizy/cmd/server/pb/peer"
	"github.com/Notch-Technologies/wizy/cmd/server/repository"
	"github.com/Notch-Technologies/wizy/store"
	"github.com/Notch-Technologies/wizy/types/key"
)

type SessionUsecaseCaller interface {
	CreatePeer(setupKey, clientMachinePubKey, serverMachinePubKey, wgPubKey string) (*domain.Peer, error)
	ValidateSetupKey(setupKey string) error
}

type SessionUsecase struct {
	setupKeyRepository *repository.SetupKeyRepository
	userRepository     *repository.UserRepository
	peerRepository     *repository.PeerRepository
	networkRepository  *repository.NetworkRepository
	serverStore        *store.ServerStore
	peerUpdateManager  *channel.PeersUpdateManager
	config             *config.ServerConfig
}

func NewSessionUsecase(
	db database.SQLExecuter, config *config.ServerConfig,
	server *store.ServerStore,
	peerUpdateManager *channel.PeersUpdateManager,
) SessionUsecaseCaller {
	return &SessionUsecase{
		setupKeyRepository: repository.NewSetupKeyRepository(db),
		userRepository:     repository.NewUserRepository(db),
		peerRepository:     repository.NewPeerRepository(db),
		networkRepository:  repository.NewNetworkRepository(db),
		serverStore:        server,
		peerUpdateManager:  peerUpdateManager,
		config:             config,
	}
}

func (s *SessionUsecase) createStunTurnConfig() *peer.StunTurnConfig {
	var (
		stuns []*peer.Host
		turns []*peer.Host
	)

	for _, stun := range s.config.Stuns {
		p := &peer.Host{
			Url:      stun.URL,
			Username: *stun.Username,
			Password: *stun.Password,
		}
		stuns = append(stuns, p)

	}

	for _, turn := range s.config.TURNConfig.Turns {
		p := &peer.Host{
			Url:      turn.URL,
			Username: *turn.Username,
			Password: *turn.Password,
		}
		turns = append(turns, p)
	}

	return &peer.StunTurnConfig{
		Stuns: stuns,
		TurnCredentials: &peer.TurnCredential{
			Turns:                turns,
			CredentialsTTL:       s.config.TURNConfig.CredentialsTTL.String(),
			Secret:               s.config.TURNConfig.Secret,
			TimeBasedCredentials: s.config.TURNConfig.TimeBasedCredentials,
		},
		Signal: &peer.Host{
			Url: s.config.Signal.URL,
		},
	}
}

func (s *SessionUsecase) CreatePeer(
	setupKey, clientMachinePubKey,
	serverMachinePubKey, wgPubKey string,
) (*domain.Peer, error) {
	var (
		allowedIPsFormat = "%s/%d"
	)

	if s.serverStore.GetPublicKey() != serverMachinePubKey {
		return nil, errors.New(domain.ErrInvalidPublicKey.Error())
	}

	// if there is no setup key, the machine key associated with the peer's 
	// client is used because the peer is starting up for the first time
	//
	if setupKey == "" {
		pe, err := s.peerRepository.FindByClientPubKey(clientMachinePubKey)
		if err != nil {
			return nil, err
		}
		return pe, nil
	}

	sk, err := s.setupKeyRepository.FindBySetupKey(setupKey)
	if err != nil {
		return nil, err
	}

	user, err := s.userRepository.FindByUserID(sk.UserID)
	if err != nil {
		return nil, err
	}

	network, err := s.networkRepository.FindByNetworkID(user.NetworkID)
	if err != nil {
		return nil, err
	}

	ipNet := ip.ParseCIDRMaskToIPNet(network.IP, int(network.CIDR), 32)

	pe, err := s.peerRepository.FindBySetupKeyID(sk.ID, clientMachinePubKey)
	if err != nil {
		// if the peer is registering for the first time
		//
		if errors.Is(err, domain.ErrNoRows) {
			peers, err := s.peerRepository.FindPeersByOrganizationID(user.OrganizationID)
			if err != nil {
				return nil, err
			}

			// new allocate ip
			//
			var issIps []net.IP

			for _, p := range peers {
				ip := net.ParseIP(p.IP)
				issIps = append(issIps, ip)
			}

			allocIP, err := ip.NewAllocateIP(ipNet, issIps)
			if err != nil {
				return nil, err
			}

			// create new peer
			//
			newPeer := domain.NewPeer(
				sk.ID,
				user.NetworkID,
				user.UserGroupID,
				user.ID,
				user.OrganizationID,
				allocIP.String(),
				network.CIDR,
				clientMachinePubKey,
				wgPubKey,
			)

			err = s.peerRepository.CreatePeer(newPeer)
			if err != nil {
				return nil, err
			}

			// return already registered Peers
			//
			peers, err = s.peerRepository.FindPeersByOrganizationID(user.OrganizationID)
			if err != nil {
				return nil, err
			}

			for _, remotePeer := range peers {
				peersToSend := []*peer.RemotePeer{}
				for _, p := range peers {
					if remotePeer.WgPubKey != p.WgPubKey {
						peersToSend = append(peersToSend, &peer.RemotePeer{
							WgPubKey:   p.ClientPubKey,
							AllowedIps: []string{fmt.Sprintf(allowedIPsFormat, p.IP, p.CIDR)},
						})
					}
				}

				fmt.Printf("send peer information to the %s update channel\n", remotePeer.ClientPubKey)

				// send update message to remote peer
				//
				err := s.peerUpdateManager.SendUpdate(
					remotePeer.ClientPubKey,
					&channel.UpdateMessage{
						Update: &peer.SyncResponse{
							PeerConfig:        &peer.PeerConfig{Address: newPeer.IP},
							RemotePeers:       peersToSend,
							RemotePeerIsEmpty: len(peersToSend) == 0,
							StunTurnConfig:    s.createStunTurnConfig(),
						},
					},
				)
				if err != nil {
					return nil, err
				}

				fmt.Println("send updates that will be sent upon initial Peer registration")
			}
			return newPeer, nil
		}
		return nil, err
	}

	return pe, nil
}

func (s *SessionUsecase) ValidateSetupKey(setupKey string) error {
	customClaims, err := key.GetCustomClaims(setupKey, s.config.JwtConfig.Secret)
	if err != nil {
		return err
	}

	if customClaims.Audience != s.config.JwtConfig.Aud || customClaims.Issuer != s.config.JwtConfig.Iss {
		return errors.New(domain.ErrUnauthorizedIssOrAud.Error())
	}

	return nil
}
