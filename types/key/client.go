package key

import (
	"go4.org/mem"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"

	"github.com/Notch-Technologies/dotshake/types/structs"
)

const (
	clientPrivateKeyPrefix = "private_client_key:"
	clientPublicKeyPrefix  = "public_client_key:"
)

type DotshakeClientPrivateState struct {
	_          structs.Incomparable
	privateKey wgtypes.Key
}

func NewClientPrivateKey() (DotshakeClientPrivateState, error) {
	k, err := wgtypes.GeneratePrivateKey()
	if err != nil {
		return DotshakeClientPrivateState{}, err
	}

	return DotshakeClientPrivateState{
		privateKey: k,
	}, nil
}

func (s DotshakeClientPrivateState) MarshalText() ([]byte, error) {
	return toHex(s.privateKey[:], clientPrivateKeyPrefix), nil
}

func (s *DotshakeClientPrivateState) UnmarshalText(b []byte) error {
	return parseHex(s.privateKey[:], mem.B(b), mem.S(clientPrivateKeyPrefix))
}

func (s DotshakeClientPrivateState) PublicKey() string {
	pkey := s.privateKey.PublicKey().String()
	return pkey
}

func (s DotshakeClientPrivateState) PrivateKey() string {
	pkey := s.privateKey.String()
	return pkey
}
