package api

import (
	"encloud/config"
	"encloud/pkg/service"
	"encloud/pkg/types"
	thirdparty "encloud/third_party"
	"fmt"
	"io/ioutil"
)

func Share(uuid string, kek string, privateKey string, email string) (types.FileMetadata, error) {
	cfg, err := Fetch()
	if err != nil {
		return types.FileMetadata{}, err
	}
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

	// Writing decryption dek
	err = ioutil.WriteFile(config.Assets+"/"+fmt.Sprint(fileMetaData.Timestamp)+"_dek.txt", decryptedDek, 0777)
	if err != nil {
		return types.FileMetadata{}, err
	}

	subject := "Share content"
	r := service.NewRequest([]string{email}, subject, cfg)
	r.Send("../../templates/share.html", map[string]string{"cid": fileMetaData.Cid[0], "dekType": fileMetaData.DekType}, fileMetaData.Timestamp)

	return fileMetaData, nil
}
