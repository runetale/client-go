package domain

import "errors"

var (
	ErrUserAlredyExists     = errors.New("ERR_USER_ALREADY_EXISTS")
	ErrNetworkAlredyExists  = errors.New("ERR_NETWORK_ALREADY_EXISTS")
	ErrGroupAlredyExists    = errors.New("ERR_GROUP_ALREADY_EXISTS")
	ErrInvalidValue         = errors.New("ERR_INVALID_VALUE")
	ErrNotFound             = errors.New("ERR_NOT_FOUND")
	ErrNoRows               = errors.New("ERR_NO_ROWS")
	ErrCanNotGetAccessToken = errors.New("ERR_CAN_NOT_GET_ACCESS_TOKEN")
	ErrInvalidHeader        = errors.New("ERR_INVALID_HEADER")
	ErrNotEnoughPermission  = errors.New("ERR_NOT_ENOUGH_PERMISSION")
	ErrInvalidPublicKey     = errors.New("ERR_INVALID_PUBLIC_KEY")
	ErrUnauthorizedIssOrAud = errors.New("ERR_UNAUTHORIZED_ISS_OR_AUD")
)
