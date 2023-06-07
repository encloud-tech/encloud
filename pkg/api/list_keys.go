package api

import (
	"encloud/pkg/service"
	"encloud/pkg/types"
	"log"
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
			if collections[b.Md5Hash].Md5Hash != b.Md5Hash {
				collections[b.Md5Hash] = types.FetchKeysResponse{
					Md5Hash: b.Md5Hash,
					Files:   1,
				}
			} else {
				collections[b.Md5Hash] = types.FetchKeysResponse{
					Md5Hash: b.Md5Hash,
					Files:   collections[b.Md5Hash].Files + 1,
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
