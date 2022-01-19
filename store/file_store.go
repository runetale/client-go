package store

// this package when used write or read state file
//

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	"github.com/Notch-Technologies/wizy/paths"
	"github.com/Notch-Technologies/wizy/utils"
)

type FileStoreManager interface {
	WriteState(id StateKey, bs []byte) error
	ReadState(id StateKey) ([]byte, error)
}

type FileStore struct {
	path string

	mu sync.RWMutex
	cache map[StateKey][]byte
}

func NewFileStore(path string) (*FileStore, error) {
	if err := paths.MkStateDir(filepath.Dir(path)); err != nil {
		return nil, fmt.Errorf("does not creating state directory: %w", err)
	}

	b, err := ioutil.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			if err = utils.AtomicWriteFile(path, []byte("{}"), 0600); err != nil {
				return nil, err
			}
			return &FileStore{
				path: path, 
				cache: make(map[StateKey][]byte),
			}, nil
		}
		return nil, err
	}

	fs := &FileStore{
		path: path,
		cache: make(map[StateKey][]byte),
	}

	if err := json.Unmarshal(b, &fs.cache); err != nil {
		return nil, err
	}

	return fs, nil
}

func (s *FileStore) WriteState(id StateKey, bs []byte) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if bytes.Equal(s.cache[id], bs) {
		return nil
	}
	s.cache[id] = append([]byte(nil), bs...)
	bs, err := json.MarshalIndent(s.cache, "", "  ")
	if err != nil {
		return err
	}
	return utils.AtomicWriteFile(s.path, bs, 0600)
}

func (s *FileStore) ReadState(id StateKey) ([]byte, error) {
    s.mu.RLock()
    defer s.mu.RUnlock()
    bs, ok := s.cache[id]
    if !ok {
        return nil, ErrStateNotFound
    }
    return bs, nil
}

