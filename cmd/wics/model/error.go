package model

import "errors"

var (
	ErrUserAlredyExists = errors.New("USER_ALREADY_EXISTS")
	ErrNetworkAlredyExists = errors.New("NETWORK_ALREADY_EXISTS")
	ErrGroupAlredyExists = errors.New("GROUP_ALREADY_EXISTS")
)
