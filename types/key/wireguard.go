package key

import (
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
)

func NewGenerateKey() (string, error) {
	key, err := wgtypes.GenerateKey()
	if err != nil {
		panic(err)
	}
	return key.String(), nil
}
