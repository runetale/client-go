package puncher

import "github.com/pion/ice/v2"

type StunTurnConfig struct {
	Stun *ice.URL
	Turn *ice.URL
}

func NewStunTurnConfig(
	stun *ice.URL,
	turn *ice.URL,
) *StunTurnConfig {
	return &StunTurnConfig{
		Stun: stun,
		Turn: turn,
	}
}
