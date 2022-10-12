package cmd

import (
	"crypto/rand"
	"encoding/json"
	"filecoin-encrypted-data-storage/config"
	"filecoin-encrypted-data-storage/service"
	thirdparty "filecoin-encrypted-data-storage/third_party"
	"filecoin-encrypted-data-storage/types"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
)

func UploadContentCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upload",
		Short: "Upload your content to filecoin storage",
		Long:  `Upload your content to filecoin storage which is encrypted using your public key`,
		Run: func(cmd *cobra.Command, args []string) {
			cfg, _ := config.LoadConf("./config.yml")
			estuaryService := service.New(cfg)

			kek, _ := cmd.Flags().GetString("publicKey")
			path, _ := cmd.Flags().GetString("filePath")
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

			err = thirdparty.EncryptFile(dek, path, "assets/encrypted.bin")
			if err != nil {
				fmt.Println(err)
			}
			content := estuaryService.UploadContent("assets/encrypted.bin")
			if content.CID != "" {
				log.Println(kek)
				encryptedDek, err := thirdparty.EncryptWithRSA(dek, thirdparty.GetIdRsaPubFromStr(kek))
				if err != nil {
					fmt.Println("err" + err.Error())
				}
				fileData := types.FileMetadata{Timestamp: timestamp, Name: fileInfo.Name(), Size: int(fileInfo.Size()), FileType: filepath.Ext(fileInfo.Name()), Dek: encryptedDek, Cid: content.CID}
				service.Store(kek+"-"+content.CID, fileData)
			}

			os.Remove("assets/encrypted.bin")
			response := types.UploadContentResponse{
				Status:     "success",
				StatusCode: http.StatusCreated,
				Message:    "Content uploaded successfully.",
				Data:       content,
			}
			encoded, err := json.Marshal(response)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Fprintf(cmd.OutOrStdout(), string(encoded))
		},
	}

	cmd.Flags().StringP("publicKey", "p", "", "Enter your public key")
	cmd.Flags().StringP("filePath", "f", "", "Enter your file path")
	cmd.MarkFlagRequired("publicKey")
	cmd.MarkFlagRequired("filepath")
	return cmd
}

func init() {
	RootCmd.AddCommand(UploadContentCmd())
}
