package service

import (
	"bytes"
	"encoding/gob"
	"fmt"

	"github.com/encloud-tech/encloud/pkg/storage/badger"
	"github.com/encloud-tech/encloud/pkg/storage/couchbase"
	"github.com/encloud-tech/encloud/pkg/types"
)

// New func implements the storage interface
func NewDB(config types.ConfYaml) *DB {
	return &DB{
		config: config,
	}
}

type DB struct {
	config types.ConfYaml
}

func initBadgerDB(cfg *types.ConfYaml) *badger.Storage {
	badger := badger.New(cfg)
	err := badger.Init()
	if err != nil {
		fmt.Printf("Error in initializing badger db : %s\n", err)
	}

	return badger
}

func initCouchBaseDB(cfg *types.ConfYaml) *couchbase.Storage {
	couchbase := couchbase.New(cfg)
	err := couchbase.Init()
	if err != nil {
		fmt.Printf("Error in initializing couchbase db : %s\n", err)
	}

	return couchbase
}

func (d *DB) Store(key string, fileMetaData types.FileMetadata) {
	if d.config.Stat.StorageType == "couchbase" {
		couchbase := initCouchBaseDB(&d.config)
		couchbase.Create(key, fileMetaData)
	} else {
		badger := initBadgerDB(&d.config)
		var b bytes.Buffer
		e := gob.NewEncoder(&b)
		if err := e.Encode(fileMetaData); err != nil {
			panic(err)
		}

		badger.Create(key, b.Bytes())
		badger.Close()
	}
}

func (d *DB) FetchKeys() types.ListKeys {
	var val types.ListKeys
	couchbase := initCouchBaseDB(&d.config)
	val = couchbase.FetchKeys()
	return val
}

func (d *DB) Fetch(key string) types.FileData {
	var val types.FileData
	if d.config.Stat.StorageType == "couchbase" {
		couchbase := initCouchBaseDB(&d.config)
		val = couchbase.Read(key)
	} else {
		badger := initBadgerDB(&d.config)
		val = badger.Read(key)
		badger.Close()
	}
	return val
}

func (d *DB) FetchByCid(key string) types.FileMetadata {
	var val types.FileMetadata
	if d.config.Stat.StorageType == "couchbase" {
		couchbase := initCouchBaseDB(&d.config)
		val = couchbase.ReadByCid(key)
	} else {
		badger := initBadgerDB(&d.config)
		val = badger.ReadByCid(key)
		badger.Close()
	}

	return val
}
