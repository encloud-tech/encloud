package config

import (
	"github.com/spf13/viper"
)

var DotKeys = ".keys"
var IdRsa = ".keys/.idRsa"
var IdRsaPub = ".keys/.idRsaPub"
var KeySize = 3072

var defaultConf = []byte(`
stat:
  badgerdb:
    path: "badger.db"
`)

// ConfYaml is config structure.
type ConfYaml struct {
	Stat SectionStat `yaml:"stat"`
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

	// Stat Engine
	conf.Stat.BadgerDB.Path = viper.GetString("stat.badgerdb.path")

	return conf, nil
}
