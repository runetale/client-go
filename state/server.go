package state

import (
	"crypto/rand"
	"go4.org/mem"

	key "github.com/Notch-Technologies/wizy/types/key"
	"github.com/Notch-Technologies/wizy/types/structs"
)

const (
	serverPrivateKeyPrefix = "private_server_key:"
	serverPublicKeyPrefix = "public_server_key:"
)

type WicsServerPrivateState struct {
	_ structs.Incomparable
	key key.Key
}

func NewPreshared() (*key.Key, error) {
	var k [32]byte
	_, err := rand.Read(k[:])
	if err != nil {
		return nil, err
	}
	return (*key.Key)(&k), nil
}

func NewServerPrivateKey() (WicsServerPrivateState, error) {
	k, err := NewPreshared()
	if err != nil {
		return WicsServerPrivateState{}, err
	}

	k[0] &= 248
	k[31] = (k[31] & 127) | 64
	privateKey := (key.Key)(*k)

	return WicsServerPrivateState{
		key: privateKey,
	}, nil
}

func (s WicsServerPrivateState) MarshalText() ([]byte, error) {
	return key.ToHex(s.key[:], serverPrivateKeyPrefix), nil
}

func (s *WicsServerPrivateState) UnmarshalText(b []byte) error {
	return key.ParseHex(s.key[:], mem.B(b), mem.S(serverPrivateKeyPrefix))
}

