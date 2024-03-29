package api

import (
	"crypto/rand"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/encloud-tech/encloud/config"
	"github.com/encloud-tech/encloud/pkg/service"
	"github.com/encloud-tech/encloud/pkg/types"
	thirdparty "github.com/encloud-tech/encloud/third_party"
)

func Upload(filePath string, kekType string, dekType string, kek string) (string, error) {
	cfg, err := Fetch()
	if err != nil {
		return "", err
	}

	estuaryService := service.New(cfg)
	dbService := service.NewDB(cfg)

	const (
		DDMMYYYYhhmmss = "2006-01-02 15:04:05"
	)
	uploadedAt := time.Now().UTC().Format(DDMMYYYYhhmmss)
	timestamp := time.Now().Unix()
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("File open error : ", err)
		os.Exit(-1)
	}

	defer file.Close()

	fileInfo, _ := file.Stat()

	//generate a random 32 byte key for AES-256
	dek := make([]byte, 32)
	if _, err := rand.Read(dek); err != nil {
		return "", err
	}

	if _, err := os.Stat(config.Assets); os.IsNotExist(err) {
		err := os.Mkdir(config.Assets, 0777)
		if err != nil {
			return "", err
		}
	}

	if dekType == "aes" {
		err = thirdparty.EncryptWithAES(dek, filePath, config.Assets+"/"+thirdparty.GenerateFileName(timestamp, "encrypt", filepath.Ext(fileInfo.Name())))
		if err != nil {
			return "", err
		}
	} else {
		err = thirdparty.EncryptWithChacha20poly1305(dek, filePath, config.Assets+"/"+thirdparty.GenerateFileName(timestamp, "encrypt", filepath.Ext(fileInfo.Name())))
		if err != nil {
			return "", err
		}
	}

	var cids []string
	var uuid = thirdparty.GenerateUuid()
	content, err := estuaryService.UploadContent(config.Assets + "/" + thirdparty.GenerateFileName(timestamp, "encrypt", filepath.Ext(fileInfo.Name())))
	if err != nil {
		return "", err
	}
	if content.CID != "" {
		cids = append(cids, content.CID)
	}

	if len(cids) > 0 {
		var encryptedDek []byte
		if kekType == "rsa" {
			encryptedDek, err = thirdparty.EncryptWithRSA(dek, thirdparty.GetIdRsaPubFromStr(kek))
			if err != nil {
				return "", err
			}
		} else if kekType == "ecies" {
			encryptedDek, err = thirdparty.EncryptWithEcies(thirdparty.NewPublicKeyFromHex(kek), dek)
			if err != nil {
				return "", err
			}
		} else {
			return "", err
		}
		hash := thirdparty.DigestString(kek)
		fileData := types.FileMetadata{
			Timestamp:  timestamp,
			Name:       fileInfo.Name(),
			Size:       int(fileInfo.Size()),
			FileType:   filepath.Ext(fileInfo.Name()),
			Dek:        encryptedDek,
			Cid:        cids,
			Uuid:       uuid,
			Md5Hash:    hash,
			PublicKey:  kek,
			DekType:    dekType,
			KekType:    kekType,
			UploadedAt: uploadedAt,
		}
		dbService.Store(hash+":"+uuid, fileData)
		os.Remove(config.Assets + "/" + thirdparty.GenerateFileName(timestamp, "encrypt", filepath.Ext(fileInfo.Name())))
		return uuid, nil
	} else {
		return "", errors.New("Something went wrong!")
	}
}
