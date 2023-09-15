package api

import (
	"fmt"
	"os"
	"time"

	"github.com/encloud-tech/encloud/config"
	"github.com/encloud-tech/encloud/pkg/service"
	thirdparty "github.com/encloud-tech/encloud/third_party"
)

func RetrieveSharedContent(decryptedDekPath string, dekType string, cid string, fileName string, retrievalFileStoragePath string) error {
	cfg, err := Fetch()
	if err != nil {
		return err
	}
	estuaryService := service.New(cfg)

	dek, err := thirdparty.ReadFile(decryptedDekPath)
	if err != nil {
		return err
	}

	timestamp := time.Now().Unix()

	baseApiErr := estuaryService.DownloadContent(config.Assets+"/"+fileName+"_"+fmt.Sprint(timestamp), cid, cfg.Estuary.BaseApiUrl)
	if baseApiErr != nil {
		gatewayApiErr := estuaryService.DownloadContent(config.Assets+"/"+fileName+"_"+fmt.Sprint(timestamp), cid, cfg.Estuary.GatewayApiUrl)
		if gatewayApiErr != nil {
			cdnApiErr := estuaryService.DownloadContent(config.Assets+"/"+fileName+"_"+fmt.Sprint(timestamp), cid, cfg.Estuary.CdnApiUrl)
			return cdnApiErr
		}
	}

	filePath := config.Assets + "/" + fileName + "_" + fmt.Sprint(timestamp)

	if dekType == "aes" {
		err := thirdparty.DecryptWithAES(dek, filePath, retrievalFileStoragePath+"/"+fileName)
		if err != nil {
			return err
		}
	} else {
		err := thirdparty.DecryptWithChacha20poly1305(dek, filePath, retrievalFileStoragePath+"/"+fileName)
		if err != nil {
			return err
		}
	}

	os.Remove(config.Assets + "/" + fileName + "_" + fmt.Sprint(timestamp))
	return nil
}
