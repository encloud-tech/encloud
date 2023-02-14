package api

import (
	"encloud/config"
	"encloud/pkg/service"
	"encloud/pkg/types"
	thirdparty "encloud/third_party"
	"os"
)

func RetrieveByUUID(uuid string, kek string, privateKey string) (types.FileMetadata, error) {
	cfg, err := config.LoadConf("./config.yaml")
	if err != nil {
		// Load default configuration from config.go file if config.yaml file not found
		cfg, _ = config.LoadConf()
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

	filepath := estuaryService.DownloadContent(fileMetaData.Cid[0])
	if fileMetaData.DekType == "aes" {
		err := thirdparty.DecryptWithAES(decryptedDek, filepath, "assets/decrypted.csv")
		if err != nil {
			return types.FileMetadata{}, err
		}
	} else {
		err := thirdparty.DecryptWithChacha20poly1305(decryptedDek, filepath, "assets/decrypted.csv")
		if err != nil {
			return types.FileMetadata{}, err
		}
	}

	os.Remove("assets/downloaded.bin")
	return fileMetaData, nil
}
