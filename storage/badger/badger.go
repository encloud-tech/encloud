package badger

import (
	"filecoin-encrypted-data-storage/config"
	"fmt"
	"log"
	"os"

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

	// lock sync.RWMutex
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

func (s *Storage) Read(key string) []byte {
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
