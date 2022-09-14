package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// Test file is missing
func TestMissingFile(t *testing.T) {
	filename := "test"
	_, err := LoadConf(filename)

	assert.NotNil(t, err)
}

type ConfigTestSuite struct {
	suite.Suite
	ConfDefault *ConfYaml
	Conf        *ConfYaml
}

func (suite *ConfigTestSuite) SetupTest() {
	var err error
	suite.ConfDefault, err = LoadConf()
	if err != nil {
		panic("failed to load default config.yml")
	}
	suite.Conf, err = LoadConf("testdata/config.yml")
	if err != nil {
		panic("failed to load config.yml from file")
	}
}

func (suite *ConfigTestSuite) TestValidateConfDefault() {
	// estuary
	assert.Equal(suite.T(), "https://dweb.link/ipfs", suite.ConfDefault.Estuary.DownloadApiUrl)
	assert.Equal(suite.T(), "https://api.estuary.tech", suite.ConfDefault.Estuary.BaseApiUrl)
	assert.Equal(suite.T(), "https://shuttle-4.estuary.tech", suite.ConfDefault.Estuary.ShuttleApiUrl)
	assert.Equal(suite.T(), "ESTb2e5e305-1af1-4c72-89ab-c85404439fcdARY", suite.ConfDefault.Estuary.Token)

	assert.Equal(suite.T(), "badger.db", suite.ConfDefault.Stat.BadgerDB.Path)
}

func (suite *ConfigTestSuite) TestValidateConf() {
	// estuary
	assert.Equal(suite.T(), "https://dweb.link/ipfs", suite.Conf.Estuary.DownloadApiUrl)
	assert.Equal(suite.T(), "https://api.estuary.tech", suite.Conf.Estuary.BaseApiUrl)
	assert.Equal(suite.T(), "https://shuttle-4.estuary.tech", suite.Conf.Estuary.ShuttleApiUrl)
	assert.Equal(suite.T(), "ESTb2e5e305-1af1-4c72-89ab-c85404439fcdARY", suite.Conf.Estuary.Token)

	assert.Equal(suite.T(), "badger.db", suite.Conf.Stat.BadgerDB.Path)
}

func TestConfigTestSuite(t *testing.T) {
	suite.Run(t, new(ConfigTestSuite))
}

func TestLoadWrongDefaultYAMLConfig(t *testing.T) {
	defaultConf = []byte(`a`)
	_, err := LoadConf()
	assert.Error(t, err)
}
