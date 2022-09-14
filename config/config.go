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

var defaultConf = []byte(`
estuary:
  shuttle_api_url: "https://shuttle-4.estuary.tech"
  download_api_url: "https://dweb.link/ipfs"
  base_api_url: "https://api.estuary.tech",
  token: "ESTb2e5e305-1af1-4c72-89ab-c85404439fcdARY"	
stat:
  badgerdb:
    path: "badger.db"
`)

// ConfYaml is config structure.
type ConfYaml struct {
	Estuary SectionEstuary `yaml:"estuary"`
	Stat    SectionStat    `yaml:"stat"`
}

// SectionEstuary is sub section of config.
type SectionEstuary struct {
	ShuttleApiUrl  string `yaml:"shuttle_api_url"`
	DownloadApiUrl string `yaml:"download_api_url"`
	BaseApiUrl     string `yaml:"base_api_url"`
	Token          string `yaml:"token"`
}

// SectionStat is sub section of config.
type SectionStat struct {
	BadgerDB SectionBadgerDB `yaml:"badgerdb"`
}

// SectionBadgerDB is sub section of config.
type SectionBadgerDB struct {
	Path string `yaml:"path"`
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

	// Stat Engine
	conf.Stat.BadgerDB.Path = viper.GetString("stat.badgerdb.path")

	return conf, nil
}
