package state

import (
	"crypto/rand"
	"fmt"
	"encoding/hex"
	"go4.org/mem"
	"errors"

	"github.com/Notch-Technologies/wizy/types/key"
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
	return toHex(s.key[:], serverPrivateKeyPrefix), nil
}

func (s *WicsServerPrivateState) UnmarshalText(b []byte) error {
	return parseHex(s.key[:], mem.B(b), mem.S(serverPrivateKeyPrefix))
}

func toHex(k []byte, prefix string) []byte {
	ret := make([]byte, len(prefix)+len(k)*2)
	copy(ret, prefix)
	hex.Encode(ret[len(prefix):], k)
	return ret
}

func parseHex(out []byte, in, prefix mem.RO) error {
	if !mem.HasPrefix(in, prefix) {
		return fmt.Errorf("key hex string doesn't have expected type prefix %s", prefix.StringCopy())
	}
	in = in.SliceFrom(prefix.Len())
	if want := len(out) * 2; in.Len() != want {
		return fmt.Errorf("key hex has the wrong size, got %d want %d", in.Len(), want)
	}
	for i := range out {
		a, ok1 := fromHexChar(in.At(i*2 + 0))
		b, ok2 := fromHexChar(in.At(i*2 + 1))
		if !ok1 || !ok2 {
			return errors.New("invalid hex character in key")
		}
		out[i] = (a << 4) | b
	}

	return nil
}

func fromHexChar(c byte) (byte, bool) {
	switch {
	case '0' <= c && c <= '9':
		return c - '0', true
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10, true
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10, true
	}

	return 0, false
}

