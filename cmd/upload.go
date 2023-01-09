package cmd

import (
	"crypto/rand"
	"encloud/config"
	"encloud/service"
	thirdparty "encloud/third_party"
	"encloud/types"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
)

func UploadContentCmd() *cobra.Command {
	cfg, _ := config.LoadConf("./config.yaml")
	cmd := &cobra.Command{
		Use:   "upload",
		Short: "Upload your content to filecoin storage",
		Long:  `Upload your content to filecoin storage which is encrypted using your public key`,
		Run: func(cmd *cobra.Command, args []string) {
			estuaryService := service.New(cfg)
			dbService := service.NewDB(cfg)

			kek := ""
			publicKey, _ := cmd.Flags().GetString("publicKey")
			path, _ := cmd.Flags().GetString("filePath")
			dekType, _ := cmd.Flags().GetString("dekType")
			readPublicKeyFromPath, _ := cmd.Flags().GetBool("readPublicKeyFromPath")
			if readPublicKeyFromPath {
				kek = thirdparty.ReadKeyFile(publicKey)
			} else {
				kek = publicKey
			}
			timestamp := time.Now().Unix()
			file, err := os.Open(path)
			if err != nil {
				fmt.Println("File open error : ", err)
				os.Exit(-1)
			}

			defer file.Close()

			fileInfo, _ := file.Stat()

			//generate a random 32 byte key for AES-256
			dek := make([]byte, 32)
			if _, err := rand.Read(dek); err != nil {
				fmt.Println(err)
			}

			if _, err := os.Stat("assets"); os.IsNotExist(err) {
				err := os.Mkdir("assets", 0777)
				if err != nil {
					fmt.Println(err)
				}
			}

			if dekType == "aes" {
				err = thirdparty.EncryptWithAES(dek, path, "assets/encrypted.bin")
				if err != nil {
					fmt.Println(err)
				}
			} else {
				err = thirdparty.EncryptWithChacha20poly1305(dek, path, "assets/encrypted.bin")
				if err != nil {
					fmt.Println(err)
				}
			}

			var cids []string
			var uuid = thirdparty.GenerateUuid()
			content, err := estuaryService.UploadContent("assets/encrypted.bin")
			if err != nil {
				fmt.Println(err)
			}
			cids = append(cids, content.CID)

			if cids != nil {
				var encryptedDek []byte
				if cfg.Stat.EncryptionAlgorithmType == "rsa" {
					encryptedDek, err = thirdparty.EncryptWithRSA(dek, thirdparty.GetIdRsaPubFromStr(kek))
					if err != nil {
						fmt.Println("err" + err.Error())
					}
				} else {
					encryptedDek, err = thirdparty.EncryptWithEcies(thirdparty.NewPublicKeyFromHex(kek), dek)
					if err != nil {
						fmt.Println("err" + err.Error())
					}
				}
				hash := thirdparty.DigestString(kek)
				fileData := types.FileMetadata{Timestamp: timestamp, Name: fileInfo.Name(), Size: int(fileInfo.Size()), FileType: filepath.Ext(fileInfo.Name()), Dek: encryptedDek, Cid: cids, Uuid: uuid, Md5Hash: hash, DekType: dekType}
				dbService.Store(hash+":"+uuid, fileData)
			}

			os.Remove("assets/encrypted.bin")
			response := types.UploadContentResponse{
				Status:     "success",
				StatusCode: http.StatusCreated,
				Message:    "Content uploaded successfully.",
				Data:       types.Uuid{Uuid: uuid},
			}
			encoded, err := json.MarshalIndent(response, "", "    ")
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Fprintf(cmd.OutOrStdout(), string(encoded))
		},
	}

	cmd.Flags().StringP("publicKey", "p", "", "Enter your public key")
	cmd.Flags().BoolP("readPublicKeyFromPath", "r", false, "Do you want public key read from path you have entered?")
	cmd.Flags().StringP("filePath", "f", "", "Enter your file path")
	cmd.Flags().StringP("dekType", "e", "chacha20", "Enter which type of encryption do you want?")
	cmd.MarkFlagRequired("publicKey")
	cmd.MarkFlagRequired("filepath")
	return cmd
}

func init() {
	RootCmd.AddCommand(UploadContentCmd())
}
