package key

import (
	"go4.org/mem"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"

	"github.com/Notch-Technologies/wizy/types/structs"
)

const (
	serverPrivateKeyPrefix = "private_server_key:"
	serverPublicKeyPrefix  = "public_server_key:"
)

type WicsServerPrivateState struct {
	_          structs.Incomparable
	privateKey wgtypes.Key
}

func NewServerPrivateKey() (WicsServerPrivateState, error) {
	k, err := wgtypes.GeneratePrivateKey()
	if err != nil {
		return WicsServerPrivateState{}, err
	}

	return WicsServerPrivateState{
		privateKey: k,
	}, nil
}

func (s WicsServerPrivateState) MarshalText() ([]byte, error) {
	return toHex(s.privateKey[:], serverPrivateKeyPrefix), nil
}

func (s *WicsServerPrivateState) UnmarshalText(b []byte) error {
	return parseHex(s.privateKey[:], mem.B(b), mem.S(serverPrivateKeyPrefix))
}

func (s WicsServerPrivateState) PublicKey() string {
	pkey := s.privateKey.PublicKey().String()
	return pkey
}

func (s WicsServerPrivateState) PrivateKey() string {
	pkey := s.privateKey.String()
	return pkey
}
