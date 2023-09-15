package api

import (
	"log"

	"github.com/encloud-tech/encloud/pkg/service"
	"github.com/encloud-tech/encloud/pkg/types"
)

func ListKeys() (types.ListKeys, error) {
	cfg, err := Fetch()
	if err != nil {
		log.Println("Error load config data", err.Error())
		return nil, err
	}
	dbService := service.NewDB(cfg)

	if cfg.Stat.StorageType == "couchbase" {
		keys := dbService.FetchKeys()
		return keys, nil
	} else {
		fileMetadata := dbService.Fetch("")

		collections := make(map[string]types.FetchKeysResponse)
		for _, b := range fileMetadata {
			if collections[b.PublicKey].PublicKey != b.PublicKey {
				collections[b.PublicKey] = types.FetchKeysResponse{
					PublicKey: b.PublicKey,
					Files:     1,
				}
			} else {
				collections[b.PublicKey] = types.FetchKeysResponse{
					PublicKey: b.PublicKey,
					Files:     collections[b.PublicKey].Files + 1,
				}
			}
		}

		var keys types.ListKeys

		for key := range collections {
			keys = append(keys, collections[key])
		}

		return keys, nil
	}
}
