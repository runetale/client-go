package store

import "sync"

type AccountManager interface {
	CreateSetupKey()
}

type Account struct {
	StoreManager StoreManager
	mu sync.Mutex
}

func NewAccount(sm StoreManager) *Account {
	return &Account{
		StoreManager: sm,
		mu: sync.Mutex{},
	}
}

func (ac *Account) CreateSetupKey() {
	panic("not implement CreateSetupKey")
}
