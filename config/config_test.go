package config

import (
	"github.com/encloud-tech/encloud/pkg/types"

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
