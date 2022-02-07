package domain

import "errors"

var (
	ErrUserAlredyExists    = errors.New("ERR_USER_ALREADY_EXISTS")
	ErrNetworkAlredyExists = errors.New("ERR_NETWORK_ALREADY_EXISTS")
	ErrGroupAlredyExists   = errors.New("ERR_GROUP_ALREADY_EXISTS")
	ErrInvalidValue   	   = errors.New("ERR_INVALID_VALUE")
	ErrNotFound = errors.New("ERR_NOT_FOUND")
)
