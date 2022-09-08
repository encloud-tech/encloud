package service

import (
	"bytes"
	"encoding/gob"
	"filecoin-encrypted-data-storage/config"
	"fmt"

	"filecoin-encrypted-data-storage/storage/badger"
)

type FileMetadata struct {
	Timestamp int64
	Name      string
	Size      int
	FileType  string
	Cid       string
	Dek       []byte
}

func initDB() *badger.Storage {
	cfg, _ := config.LoadConf()

	badger := badger.New(cfg)
	err := badger.Init()
	if err != nil {
		fmt.Printf("ERROR saving to badger db : %s\n", err)
	}

	return badger
}

func Store(key string, fileMetaData FileMetadata) {
	badger := initDB()
	var b bytes.Buffer
	e := gob.NewEncoder(&b)
	if err := e.Encode(fileMetaData); err != nil {
		panic(err)
	}

	badger.Create(key, b.Bytes())
	badger.Close()
}

func Fetch(key string) FileMetadata {
	badger := initDB()
	val := badger.Read(key)
	var metaDataDecode FileMetadata
	d := gob.NewDecoder(bytes.NewReader(val))
	if err := d.Decode(&metaDataDecode); err != nil {
		panic(err)
	}

	badger.Close()
	return metaDataDecode
}
