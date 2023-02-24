package config

import (
	"encloud/pkg/types"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ConfigTestSuite struct {
	suite.Suite
	ConfDefault *types.ConfYaml
}

func (suite *ConfigTestSuite) SetupTest() {
	err := LoadDefaultConf()
	if err != nil {
		panic("failed to load default config.yaml")
	}
}

func (suite *ConfigTestSuite) TestValidateConfDefault() {
	// estuary
	assert.Equal(suite.T(), "https://api.estuary.tech", suite.ConfDefault.Estuary.BaseApiUrl)
	assert.Equal(suite.T(), "ESTb2e5e305-1af1-4c72-89ab-c85404439fcdARY", suite.ConfDefault.Estuary.Token)

	assert.Equal(suite.T(), "badger.db", suite.ConfDefault.Stat.BadgerDB.Path)
}

func TestConfigTestSuite(t *testing.T) {
	suite.Run(t, new(ConfigTestSuite))
}
