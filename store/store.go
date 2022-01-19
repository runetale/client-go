package store

import "errors"

type StateKey string

var ErrStateNotFound = errors.New("state not found")

const (
	ServerPrivateKeyStateKey = StateKey("server-private-key")
)
