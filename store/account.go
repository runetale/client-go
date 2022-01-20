package store

// Management Account Store

import "sync"

type AccountManager interface {
	CreateSetupKey()
}

type AccountStore struct {
	FileStoreManager FileStoreManager
	mu               sync.Mutex
}

func NewAccountStore(f FileStoreManager) *AccountStore {
	return &AccountStore{
		FileStoreManager: f,
		mu:               sync.Mutex{},
	}
}

func (as *AccountStore) CreateSetupKey() {
	panic("not implement CreateSetupKey")
}
