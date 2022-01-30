package key

import (
	"crypto/rand"
	"encoding/base64"

	"go4.org/mem"

	"github.com/Notch-Technologies/wizy/types/structs"
)

const (
	serverPrivateKeyPrefix = "private_server_key:"
	serverPublicKeyPrefix  = "public_server_key:"
)

type WicsServerPrivateState struct {
	_   structs.Incomparable
	key Key
}

// TODO: (shintard) WireGuard no tukau
func NewPresharedKey() (*Key, error) {
	var k [32]byte
	_, err := rand.Read(k[:])
	if err != nil {
		return nil, err
	}
	return (*Key)(&k), nil
}

func NewServerPrivateKey() (WicsServerPrivateState, error) {
	k, err := NewPresharedKey()
	if err != nil {
		return WicsServerPrivateState{}, err
	}

	k[0] &= 248
	k[31] &= 127
	k[31] |= 64

	privateKey := (Key)(*k)

	return WicsServerPrivateState{
		key: privateKey,
	}, nil
}

func (s *WicsServerPrivateState) String() string { return base64.StdEncoding.EncodeToString(s.key[:]) }

func (s WicsServerPrivateState) MarshalText() ([]byte, error) {
	return toHex(s.key[:], serverPrivateKeyPrefix), nil
}

func (s *WicsServerPrivateState) UnmarshalText(b []byte) error {
	return parseHex(s.key[:], mem.B(b), mem.S(serverPrivateKeyPrefix))
}

func (s WicsServerPrivateState) PublicKey() string {
	pkey := s.key.Public().HexString()
	return pkey
}

func (s WicsServerPrivateState) PrivateKey() string {
	pkey := s.key.HexString()
	return pkey
}

func (s *WicsServerPrivateState) IsZero() bool {
	return s.key.IsZero()
}
