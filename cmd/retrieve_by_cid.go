package cmd

import (
	"encloud/config"
	"encloud/service"
	thirdparty "encloud/third_party"
	"encloud/types"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

func RetrieveByCidCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "retrieve-by-cid",
		Short: "Retrieve specific uploaded content using your cid",
		Long:  `Retrieve specific uploaded content using your cid and decrypt it using your private key`,
		Run: func(cmd *cobra.Command, args []string) {
			cfg, _ := config.LoadConf("./config.yaml")
			estuaryService := service.New(cfg)
			dbService := service.NewDB(cfg)

			kek := ""
			privateKey := ""
			publicKey, _ := cmd.Flags().GetString("publicKey")
			pk, _ := cmd.Flags().GetString("privateKey")
			uuid, _ := cmd.Flags().GetString("uuid")
			readPublicKeyFromPath, _ := cmd.Flags().GetBool("readPublicKeyFromPath")
			readPrivateKeyFromPath, _ := cmd.Flags().GetBool("readPrivateKeyFromPath")
			if readPublicKeyFromPath {
				kek = thirdparty.ReadKeyFile(publicKey)
			} else {
				kek = publicKey
			}

			if readPrivateKeyFromPath {
				privateKey = thirdparty.ReadKeyFile(pk)
			} else {
				privateKey = pk
			}

			fileMetaData := dbService.FetchByCid(thirdparty.DigestString(kek) + ":" + uuid)
			var decryptedDek []byte
			if fileMetaData.KekType == "rsa" {
				rsaKey, err := thirdparty.DecryptWithRSA(fileMetaData.Dek, thirdparty.GetIdRsaFromStr(privateKey))
				if err != nil {
					fmt.Println(err)
				}
				decryptedDek = rsaKey
			} else if fileMetaData.KekType == "ecies" {
				rsaKey, err := thirdparty.DecryptWithEcies(thirdparty.NewPrivateKeyFromHex(privateKey), fileMetaData.Dek)
				if err != nil {
					fmt.Println("err" + err.Error())
				}
				decryptedDek = rsaKey
			} else {
				fmt.Fprintf(cmd.OutOrStderr(), "Invalid argument")
				return
			}

			filepath := estuaryService.DownloadContent(fileMetaData.Cid[0])
			if fileMetaData.DekType == "aes" {
				err := thirdparty.DecryptWithAES(decryptedDek, filepath, "assets/decrypted.csv")
				if err != nil {
					fmt.Println(err)
				}
			} else {
				err := thirdparty.DecryptWithChacha20poly1305(decryptedDek, filepath, "assets/decrypted.csv")
				if err != nil {
					fmt.Println(err)
				}
			}

			os.Remove("assets/downloaded.bin")
			response := types.RetrieveByCIDContentResponse{
				Status:     "success",
				StatusCode: http.StatusFound,
				Message:    "Content fetched successfully.",
				Data:       fileMetaData,
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
	cmd.Flags().StringP("privateKey", "k", "", "Enter your private key")
	cmd.Flags().StringP("uuid", "u", "", "Enter your uuid")
	cmd.Flags().BoolP("readPublicKeyFromPath", "r", false, "Do you want public key read from path you have entered?")
	cmd.Flags().BoolP("readPrivateKeyFromPath", "o", false, "Do you want private key read from path you have entered?")
	cmd.MarkFlagRequired("publicKey")
	cmd.MarkFlagRequired("privateKey")
	cmd.MarkFlagRequired("uuid")
	return cmd
}

func init() {
	RootCmd.AddCommand(RetrieveByCidCmd())
}
