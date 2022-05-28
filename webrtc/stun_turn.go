package pion

import "github.com/pion/ice/v2"

type StunTurnConfig struct {
	Stun *ice.URL
	Turn *ice.URL
}
