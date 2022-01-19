package store

// Management Account Store

import "sync"

type AccountManager interface {
	CreateSetupKey()
}

type Account struct {
	FileStoreManager FileStoreManager
	mu               sync.Mutex
}

func NewAccount(f FileStoreManager) *Account {
	return &Account{
		FileStoreManager: f,
		mu:               sync.Mutex{},
	}
}

func (ac *Account) CreateSetupKey() {
	panic("not implement CreateSetupKey")
}
