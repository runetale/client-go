package store

import "errors"

type StateKey string

var ErrStateNotFound = errors.New("state not found")

const (
	ClientPrivateKeyStateKey = StateKey("client-private-key")
)
