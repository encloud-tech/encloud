package api

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/encloud-tech/encloud/config"
	"github.com/encloud-tech/encloud/pkg/service"
	"github.com/encloud-tech/encloud/pkg/types"
	thirdparty "github.com/encloud-tech/encloud/third_party"
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
	if sent := r.Send(fileMetaData.Cid[0], fileMetaData.DekType, fileMetaData.Timestamp, cfg.Email.EmailType); sent {
		os.Remove(config.Assets + "/" + fmt.Sprint(fileMetaData.Timestamp) + "_dek.txt")
		return fileMetaData, nil
	} else {
		return types.FileMetadata{}, errors.New("Failed to send the email")
	}
}
