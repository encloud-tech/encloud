package api

import (
	"errors"
	"log"
	"os"
	"path/filepath"

	"github.com/encloud-tech/encloud/config"
	"github.com/encloud-tech/encloud/pkg/service"
	"github.com/encloud-tech/encloud/pkg/types"
	thirdparty "github.com/encloud-tech/encloud/third_party"
)

func RetrieveByUUID(uuid string, kek string, privateKey string, retrievalFileStoragePath string) (types.FileMetadata, error) {
	if retrievalFileStoragePath == "" {
		if _, err := os.Stat(config.Download); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(config.Download, os.ModePerm)
			if err != nil {
				log.Println(err)
			}
		}
		retrievalFileStoragePath = config.Download
	}

	cfg, err := Fetch()
	if err != nil {
		return types.FileMetadata{}, err
	}
	estuaryService := service.New(cfg)
	dbService := service.NewDB(cfg)

	fileMetaData := dbService.FetchByCid(thirdparty.DigestString(kek) + ":" + uuid)
	var decryptedDek []byte
	if fileMetaData.KekType == "rsa" {
		rsaKey, err := thirdparty.DecryptWithRSA(fileMetaData.Dek, thirdparty.GetIdRsaFromStr(privateKey))
		if err != nil {
			return types.FileMetadata{}, err
		}
		decryptedDek = rsaKey
	} else if fileMetaData.KekType == "ecies" {
		rsaKey, err := thirdparty.DecryptWithEcies(thirdparty.NewPrivateKeyFromHex(privateKey), fileMetaData.Dek)
		if err != nil {
			return types.FileMetadata{}, err
		}
		decryptedDek = rsaKey
	} else {
		return types.FileMetadata{}, err
	}

	baseApiErr := estuaryService.DownloadContent(config.Assets+"/"+thirdparty.GenerateFileName(fileMetaData.Timestamp, "retrieve", filepath.Ext(fileMetaData.Name)), fileMetaData.Cid[0], cfg.Estuary.BaseApiUrl)
	if baseApiErr != nil {
		gatewayApiErr := estuaryService.DownloadContent(config.Assets+"/"+thirdparty.GenerateFileName(fileMetaData.Timestamp, "retrieve", filepath.Ext(fileMetaData.Name)), fileMetaData.Cid[0], cfg.Estuary.GatewayApiUrl)
		if gatewayApiErr != nil {
			cdnApiErr := estuaryService.DownloadContent(config.Assets+"/"+thirdparty.GenerateFileName(fileMetaData.Timestamp, "retrieve", filepath.Ext(fileMetaData.Name)), fileMetaData.Cid[0], cfg.Estuary.CdnApiUrl)
			return types.FileMetadata{}, cdnApiErr
		}
	}

	filePath := config.Assets + "/" + thirdparty.GenerateFileName(fileMetaData.Timestamp, "retrieve", filepath.Ext(fileMetaData.Name))

	if fileMetaData.DekType == "aes" {
		err := thirdparty.DecryptWithAES(decryptedDek, filePath, retrievalFileStoragePath+"/"+fileMetaData.Name)
		if err != nil {
			return types.FileMetadata{}, err
		}
	} else {
		err := thirdparty.DecryptWithChacha20poly1305(decryptedDek, filePath, retrievalFileStoragePath+"/"+fileMetaData.Name)
		if err != nil {
			return types.FileMetadata{}, err
		}
	}

	os.Remove(config.Assets + "/" + thirdparty.GenerateFileName(fileMetaData.Timestamp, "retrieve", filepath.Ext(fileMetaData.Name)))
	return fileMetaData, nil
}
