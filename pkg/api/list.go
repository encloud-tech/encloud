package api

import (
	"encloud/pkg/service"
	"encloud/pkg/types"
	thirdparty "encloud/third_party"
	"log"
)

func List(kek string) types.FileData {
	cfg, err := Fetch()
	if err != nil {
		log.Println("Error load config data", err.Error())
	}
	dbService := service.NewDB(cfg)

	return dbService.Fetch(thirdparty.DigestString(kek))
}
