package service

import (
	"bytes"
	"encloud/config"
	"encloud/pkg/storage/badger"
	"encloud/pkg/storage/couchbase"
	"encloud/pkg/types"
	"encoding/gob"
	"fmt"
)

// New func implements the storage interface
func NewDB(config *config.ConfYaml) *DB {
	return &DB{
		config: config,
	}
}

type DB struct {
	config *config.ConfYaml
}

func initBadgerDB() *badger.Storage {
	cfg, _ := config.LoadConf()

	badger := badger.New(cfg)
	err := badger.Init()
	if err != nil {
		fmt.Printf("Error in initializing badger db : %s\n", err)
	}

	return badger
}

func initCouchBaseDB() *couchbase.Storage {
	cfg, _ := config.LoadConf()

	couchbase := couchbase.New(cfg)
	err := couchbase.Init()
	if err != nil {
		fmt.Printf("Error in initializing couchbase db : %s\n", err)
	}

	return couchbase
}

func (d *DB) Store(key string, fileMetaData types.FileMetadata) {
	if d.config.Stat.StorageType == "couchbase" {
		couchbase := initCouchBaseDB()
		couchbase.Create(key, fileMetaData)
	} else {
		badger := initBadgerDB()
		var b bytes.Buffer
		e := gob.NewEncoder(&b)
		if err := e.Encode(fileMetaData); err != nil {
			panic(err)
		}

		badger.Create(key, b.Bytes())
		badger.Close()
	}
}

func (d *DB) Fetch(key string) types.FileData {
	var val types.FileData
	if d.config.Stat.StorageType == "couchbase" {
		couchbase := initCouchBaseDB()
		val = couchbase.Read(key)
	} else {
		badger := initBadgerDB()
		val = badger.Read(key)
		badger.Close()
	}
	return val
}

func (d *DB) FetchByCid(key string) types.FileMetadata {
	var val types.FileMetadata
	if d.config.Stat.StorageType == "couchbase" {
		couchbase := initCouchBaseDB()
		val = couchbase.ReadByCid(key)
	} else {
		badger := initBadgerDB()
		val = badger.ReadByCid(key)
		badger.Close()
	}

	return val
}
