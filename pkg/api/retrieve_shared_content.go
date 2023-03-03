package api

import (
	"encloud/config"
	"encloud/pkg/service"
	thirdparty "encloud/third_party"
	"fmt"
	"os"
	"time"
)

func RetrieveSharedContent(decryptedDekPath string, dekType string, cid string, fileName string, retrievalFileStoragePath string) error {
	cfg, err := Fetch()
	if err != nil {
		return err
	}
	estuaryService := service.New(cfg)

	dek := thirdparty.ReadFile(decryptedDekPath)

	timestamp := time.Now().Unix()

	filepath, err := estuaryService.DownloadContent(config.Assets+"/"+fileName+"_"+fmt.Sprint(timestamp), cid)
	if err != nil {
		return err
	}
	if dekType == "aes" {
		err := thirdparty.DecryptWithAES(dek, filepath, retrievalFileStoragePath+"/"+fileName)
		if err != nil {
			return err
		}
	} else {
		err := thirdparty.DecryptWithChacha20poly1305(dek, filepath, retrievalFileStoragePath+"/"+fileName)
		if err != nil {
			return err
		}
	}

	os.Remove(config.Assets + "/" + fileName + "_" + fmt.Sprint(timestamp))
	return nil
}
