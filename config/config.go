package config

import (
	"encloud/pkg/types"
	"io/ioutil"
	"log"

	"github.com/adrg/xdg"
	"gopkg.in/yaml.v2"
)

var DotKeys = xdg.ConfigHome + "/encloud/.keys"
var IdRsa = xdg.ConfigHome + "/encloud/.keys/.idRsa"
var IdRsaPub = xdg.ConfigHome + "/encloud/.keys/.idRsaPub"
var Assets = xdg.ConfigHome + "/encloud/assets"
var TestDir = xdg.ConfigHome + "/encloud/testdata"
var KeySize = 3072

var SaltSize = 32                  // in bytes
var NonceSize = 24                 // in bytes. taken from aead.NonceSize()
var EncryptionKeySize = uint32(32) // KeySize is 32 bytes (256 bits).
var KeyTime = uint32(5)
var KeyMemory = uint32(1024 * 64) // KeyMemory in KiB. here, 64 MiB.
var KeyThreads = uint8(4)
var ChunkSize = 1024 * 32 // chunkSize in bytes. here, 32 KiB.

// LoadConf load default config
func LoadDefaultConf() error {
	configFilePath, err := xdg.ConfigFile("encloud/config.yaml")
	if err != nil {
		return err
	}

	log.Println("Save the config file at:", configFilePath)

	conf := types.ConfYaml{
		Estuary: types.SectionEstuary{
			BaseApiUrl:    "https://api.estuary.tech",
			UploadApiUrl:  "https://edge.estuary.tech/api/v1",
			GatewayApiUrl: "https://edge.estuary.tech",
			CdnApiUrl:     "https://cdn.estuary.tech",
			Token:         "EST6315eb22-5c76-4d47-9b75-1acb4a954070ARY",
		},
		Email: types.EmailStat{
			EmailType: "mailersend",
			From:      "contact@encloud.tech",
			SMTP: types.SectionSMTP{
				Server:   "sandbox.smtp.mailtrap.io",
				Port:     2525,
				Username: "ac984e52bfd35d",
				Password: "861b495c076713",
			},
		},
		Stat: types.SectionStat{
			KekType:     "ecies",
			StorageType: "badgerdb",
			BadgerDB: types.SectionBadgerDB{
				Path: "badger.db",
			},
			Couchbase: types.SectionCouchbase{
				Host:     "localhost",
				Username: "Administrator",
				Password: "Encloud@2022",
				Bucket: types.SectionBucket{
					Name:       "encloud",
					Scope:      "file",
					Collection: "metadata",
				},
			},
		},
	}

	yamlData, err := yaml.Marshal(&conf)

	if err != nil {
		return err
	}

	err = ioutil.WriteFile(configFilePath, yamlData, 0644)
	if err != nil {
		return err
	}

	return nil
}
