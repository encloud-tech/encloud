package badger

import (
	"bytes"
	"encoding/gob"
	"filecoin-encrypted-data-storage/config"
	"filecoin-encrypted-data-storage/types"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/dgraph-io/badger/v3"
)

// New func implements the storage interface
func New(config *config.ConfYaml) *Storage {
	return &Storage{
		config: config,
	}
}

// Storage is interface structure
type Storage struct {
	config *config.ConfYaml
	opts   badger.Options
	name   string
	db     *badger.DB

	lock sync.RWMutex
}

// Init client storage.
func (s *Storage) Init() error {
	var err error
	s.name = "badger"
	dbPath := s.config.Stat.BadgerDB.Path
	if dbPath == "" {
		dbPath = os.TempDir() + "badger"
	}
	s.opts = badger.DefaultOptions(dbPath)

	s.db, err = badger.Open(s.opts)

	return err
}

// Close the storage connection
func (s *Storage) Close() error {
	if s.db == nil {
		return nil
	}

	return s.db.Close()
}

func (s *Storage) Create(key string, val []byte) {
	err := s.db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(key), val)
		return err
	})

	if err != nil {
		fmt.Printf("ERROR saving to badger db : %s\n", err)
	}
}

func (s *Storage) Read(key string) types.FileData {
	var ival types.FileData
	err := s.db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()
		prefix := []byte(key)
		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			item := it.Item()
			err := item.Value(func(v []byte) error {
				var metaDataDecode types.FileMetadata
				d := gob.NewDecoder(bytes.NewReader(v))
				if err := d.Decode(&metaDataDecode); err != nil {
					panic(err)
				}
				ival = append(ival, metaDataDecode)
				return nil
			})

			if err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		log.Println("Failed to read data from the cache.", "key", string(key), "error", err)
	}

	return ival
}

func (s *Storage) ReadByCid(key string) []byte {
	var ival []byte
	err := s.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}
		ival, err = item.ValueCopy(nil)
		return err
	})

	if err != nil {
		log.Println("Failed to read data from the cache.", "key", string(key), "error", err)
	}

	return ival
}
