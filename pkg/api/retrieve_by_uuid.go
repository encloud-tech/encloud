package api

import (
	"encloud/config"
	"encloud/pkg/service"
	"encloud/pkg/types"
	thirdparty "encloud/third_party"
	"os"
	"path/filepath"
)

func RetrieveByUUID(uuid string, kek string, privateKey string, retrievalFileStoragePath string) (types.FileMetadata, error) {
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

	filePath := estuaryService.DownloadContent(config.Assets+"/"+thirdparty.GenerateFileName(fileMetaData.Timestamp, "retrieve", filepath.Ext(fileMetaData.Name)), fileMetaData.Cid[0])
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
