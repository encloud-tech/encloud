package config

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/spf13/viper"
)

var DotKeys = ".keys"
var IdRsa = ".keys/.idRsa"
var IdRsaPub = ".keys/.idRsaPub"
var KeySize = 3072

var SaltSize = 32                  // in bytes
var NonceSize = 24                 // in bytes. taken from aead.NonceSize()
var EncryptionKeySize = uint32(32) // KeySize is 32 bytes (256 bits).
var KeyTime = uint32(5)
var KeyMemory = uint32(1024 * 64) // KeyMemory in KiB. here, 64 MiB.
var KeyThreads = uint8(4)
var ChunkSize = 1024 * 32 // chunkSize in bytes. here, 32 KiB.

var defaultConf = []byte(`
estuary:
  base_api_url: 'https://api.estuary.tech'
  download_api_url: 'https://dweb.link/ipfs'
  shuttle_api_url: 'https://shuttle-4.estuary.tech'
  token: EST6315eb22-5c76-4d47-9b75-1acb4a954070ARY
email:
  server: smtp.mailtrap.io
  port: 2525
  username: ac984e52bfd35d
  password: 861b495c076713
  from: noreply@bond180.com
stat:
  kekType: ecies
  storageType: badgerdb
  badgerdb:
    path: badger.db
  couchbase:
    host: localhost
    username: Administrator
    password: Encloud@2022
    bucket:
      name: encloud
      scope: file
      collection: metadata
`)

// ConfYaml is config structure.
type ConfYaml struct {
	Estuary SectionEstuary `yaml:"estuary"`
	Email   EmailStat      `yaml:"email"`
	Stat    SectionStat    `yaml:"stat"`
}

// SectionEstuary is sub section of config.
type SectionEstuary struct {
	ShuttleApiUrl  string `yaml:"shuttle_api_url"`
	DownloadApiUrl string `yaml:"download_api_url"`
	BaseApiUrl     string `yaml:"base_api_url"`
	Token          string `yaml:"token"`
}

// EmailStat is sub section of config.
type EmailStat struct {
	Server   string `yaml:"server"`
	Port     int64  `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	From     string `yaml:"from"`
}

// SectionStat is sub section of config.
type SectionStat struct {
	BadgerDB    SectionBadgerDB  `yaml:"badgerdb"`
	Couchbase   SectionCouchbase `yaml:"couchbase"`
	StorageType string           `yaml:"storageType"`
	KekType     string           `yaml:"kekType"`
}

// SectionBadgerDB is sub section of config.
type SectionBadgerDB struct {
	Path string `yaml:"path"`
}

// SectionCouchbae is sub section of config.
type SectionCouchbase struct {
	Host     string        `yaml:"host"`
	Username string        `yaml:"username"`
	Password string        `yaml:"password"`
	Bucket   SectionBucket `yaml:"bucket"`
}

type SectionBucket struct {
	Name       string `yaml:"name"`
	Scope      string `yaml:"scope"`
	Collection string `yaml:"collection"`
}

// LoadConf load config from file and read in environment variables that match
func LoadConf(confPath ...string) (*ConfYaml, error) {
	conf := &ConfYaml{}

	viper.SetConfigType("yaml")
	viper.AutomaticEnv() // read in environment variables that match
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if len(confPath) > 0 && confPath[0] != "" {
		content, err := ioutil.ReadFile(confPath[0])
		if err != nil {
			return conf, err
		}

		if err := viper.ReadConfig(bytes.NewBuffer(content)); err != nil {
			return conf, err
		}
	} else {
		viper.AddConfigPath("./testdata/")
		viper.SetConfigName("config")

		// If a config file is found, read it in.
		if err := viper.ReadInConfig(); err == nil {
			fmt.Println("Using config file:", viper.ConfigFileUsed())
		} else if err := viper.ReadConfig(bytes.NewBuffer(defaultConf)); err != nil {
			// load default config
			return conf, err
		}
	}

	// estuary
	conf.Estuary.BaseApiUrl = viper.GetString("estuary.base_api_url")
	conf.Estuary.DownloadApiUrl = viper.GetString("estuary.download_api_url")
	conf.Estuary.ShuttleApiUrl = viper.GetString("estuary.shuttle_api_url")
	conf.Estuary.Token = viper.GetString("estuary.token")

	// email
	conf.Email.Server = viper.GetString("email.server")
	conf.Email.Username = viper.GetString("email.username")
	conf.Email.Password = viper.GetString("email.password")
	conf.Email.From = viper.GetString("email.from")
	conf.Email.Port = viper.GetInt64("email.port")

	// Stat Engine
	conf.Stat.KekType = viper.GetString("stat.kekType")
	conf.Stat.StorageType = viper.GetString("stat.storageType")
	conf.Stat.BadgerDB.Path = viper.GetString("stat.badgerdb.path")
	conf.Stat.Couchbase.Host = viper.GetString("stat.couchbase.host")
	conf.Stat.Couchbase.Username = viper.GetString("stat.couchbase.username")
	conf.Stat.Couchbase.Password = viper.GetString("stat.couchbase.password")
	conf.Stat.Couchbase.Bucket.Name = viper.GetString("stat.couchbase.bucket.name")
	conf.Stat.Couchbase.Bucket.Scope = viper.GetString("stat.couchbase.bucket.scope")
	conf.Stat.Couchbase.Bucket.Collection = viper.GetString("stat.couchbase.bucket.collection")

	return conf, nil
}
