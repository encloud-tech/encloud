package service

import (
	"bytes"
	"encoding/gob"
	"encloud/config"
	"encloud/storage/badger"
	"encloud/types"
	"fmt"
)

func initDB() *badger.Storage {
	cfg, _ := config.LoadConf()

	badger := badger.New(cfg)
	err := badger.Init()
	if err != nil {
		fmt.Printf("ERROR saving to badger db : %s\n", err)
	}

	return badger
}

func Store(key string, fileMetaData types.FileMetadata) {
	badger := initDB()
	var b bytes.Buffer
	e := gob.NewEncoder(&b)
	if err := e.Encode(fileMetaData); err != nil {
		panic(err)
	}

	badger.Create(key, b.Bytes())
	badger.Close()
}

func FetchByCid(key string) types.FileMetadata {
	badger := initDB()
	val := badger.ReadByCid(key)
	var metaDataDecode types.FileMetadata
	d := gob.NewDecoder(bytes.NewReader(val))
	if err := d.Decode(&metaDataDecode); err != nil {
		panic(err)
	}

	badger.Close()
	return metaDataDecode
}

func Fetch(key string) types.FileData {
	badger := initDB()
	val := badger.Read(key)
	badger.Close()
	return val
}
