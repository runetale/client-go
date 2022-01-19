package store

import "sync"

type ServerManager inteface {
}

type Server struct {
	FileStoreManager FileStoreManager
	mu sync.Mutex
}
