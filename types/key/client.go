package key

import (
	"go4.org/mem"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"

	"github.com/Notch-Technologies/wizy/types/structs"
)

const (
	clientPrivateKeyPrefix = "private_client_key:"
	clientPublicKeyPrefix  = "public_client_key:"
)

type WicsClientPrivateState struct {
	_          structs.Incomparable
	privateKey wgtypes.Key
}

func NewClientPrivateKey() (WicsClientPrivateState, error) {
	k, err := wgtypes.GeneratePrivateKey()
	if err != nil {
		return WicsClientPrivateState{}, err
	}

	return WicsClientPrivateState{
		privateKey: k,
	}, nil
}

func (s WicsClientPrivateState) MarshalText() ([]byte, error) {
	return toHex(s.privateKey[:], clientPrivateKeyPrefix), nil
}

func (s *WicsClientPrivateState) UnmarshalText(b []byte) error {
	return parseHex(s.privateKey[:], mem.B(b), mem.S(clientPrivateKeyPrefix))
}

func (s WicsClientPrivateState) PublicKey() string {
	pkey := s.privateKey.PublicKey().String()
	return pkey
}

func (s WicsClientPrivateState) PrivateKey() string {
	pkey := s.privateKey.String()
	return pkey
}
