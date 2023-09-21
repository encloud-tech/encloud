package api

import (
	"github.com/encloud-tech/encloud/pkg/service"
	"github.com/encloud-tech/encloud/pkg/types"
	thirdparty "github.com/encloud-tech/encloud/third_party"
)

func RetrieveDecryptedDEKsWithFileMetadataByUUID(uuid []string, kek string, privateKey string) (types.FileData, error) {
	cfg, err := Fetch()
	if err != nil {
		return types.FileData{}, err
	}

	dbService := service.NewDB(cfg)
	var fileList types.FileData

	for _, link := range uuid {
		fileMetaData := dbService.FetchByCid(thirdparty.DigestString(kek) + ":" + link)

		var decryptedDek []byte
		if fileMetaData.KekType == "rsa" {
			rsaKey, err := thirdparty.DecryptWithRSA(fileMetaData.Dek, thirdparty.GetIdRsaFromStr(privateKey))
			if err != nil {
				return types.FileData{}, err
			}
			decryptedDek = rsaKey
		} else if fileMetaData.KekType == "ecies" {
			rsaKey, err := thirdparty.DecryptWithEcies(thirdparty.NewPrivateKeyFromHex(privateKey), fileMetaData.Dek)
			if err != nil {
				return types.FileData{}, err
			}
			decryptedDek = rsaKey
		} else {
			return types.FileData{}, err
		}

		fileMetaData.Dek = decryptedDek
		fileList = append(fileList, fileMetaData)
	}

	return fileList, nil
}
