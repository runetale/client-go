package store

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/Notch-Technologies/wizy/utils"
)

type FileStore struct {
	path string
}

func LoadFileStore(path string) (*FileStore, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			os.MkdirAll(filepath.Dir(path), 0755)
			store := FileStore{path}

			b, err := json.MarshalIndent(store, "", "\t")
			if err != nil {
				log.Fatal(err)
			}

			if err = utils.AtomicWriteFile(path, b, 0600); err != nil {
				return nil, err
			}
			return &store, nil
		}
		return nil, err
	}

	var fs FileStore

	if err := json.Unmarshal(b, &fs); err != nil {
		return nil, err
	}

	return &fs, nil
}
