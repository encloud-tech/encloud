package api

import (
	"encloud/config"
	"encloud/pkg/service"
	"encloud/pkg/types"
	thirdparty "encloud/third_party"
)

func List(kek string) types.FileData {
	cfg, err := config.LoadConf("./config.yaml")
	if err != nil {
		// Load default configuration from config.go file if config.yaml file not found
		cfg, _ = config.LoadConf()
	}
	dbService := service.NewDB(cfg)

	return dbService.Fetch(thirdparty.DigestString(kek))
}
