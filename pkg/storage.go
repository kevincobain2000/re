package pkg

import (
	"fmt"
	"os"
	"sync"

	badger "github.com/dgraph-io/badger/v4"
)

var (
	storage     *badger.DB
	storageOnce sync.Once
)

func DB() *badger.DB {
	storageOnce.Do(func() {
		tmpDir := os.TempDir() + "/re"
		var err error
		options := badger.DefaultOptions(tmpDir).WithLoggingLevel(badger.ERROR)
		storage, err = badger.Open(options)
		if err != nil {
			fmt.Println(err.Error())
		}
	})
	return storage
}

type StorageHandler struct {
}

func NewStorageHandler() *StorageHandler {
	return &StorageHandler{}
}

func (h *StorageHandler) Get(key string) (string, error) {
	var valCopy []byte
	err := DB().View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}
		valCopy, err = item.ValueCopy(nil)
		return err
	})
	return string(valCopy), err
}

func (h *StorageHandler) Set(key, value string) error {
	err := DB().Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(key), []byte(value))
		return err
	})
	return err
}

func (h *StorageHandler) Delete(key string) error {
	err := DB().Update(func(txn *badger.Txn) error {
		err := txn.Delete([]byte(key))
		return err
	})
	return err
}
