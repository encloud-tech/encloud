package api

import (
	"log"

	"github.com/encloud-tech/encloud/pkg/service"
	"github.com/encloud-tech/encloud/pkg/types"
	thirdparty "github.com/encloud-tech/encloud/third_party"
)

func List(kek string) types.FileData {
	cfg, err := Fetch()
	if err != nil {
		log.Println("Error load config data", err.Error())
	}
	dbService := service.NewDB(cfg)

	return dbService.Fetch(thirdparty.DigestString(kek))
}
