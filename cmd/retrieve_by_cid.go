package cmd

import (
	"encoding/json"
	"encloud/config"
	"encloud/service"
	thirdparty "encloud/third_party"
	"encloud/types"
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
			cfg, _ := config.LoadConf("./config.yml")
			estuaryService := service.New(cfg)

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

			fileMetaData := service.FetchByCid(kek + ":" + uuid)
			decryptedDek, err := thirdparty.DecryptWithRSA(fileMetaData.Dek, thirdparty.GetIdRsaFromStr(privateKey))
			if err != nil {
				fmt.Println(err)
			}

			filepath := estuaryService.DownloadContent(fileMetaData.Cid[0])
			err = thirdparty.DecryptFile(decryptedDek, filepath)
			if err != nil {
				fmt.Println(err)
			}
			os.Remove("assets/downloaded.bin")
			response := types.RetrieveByCIDContentResponse{
				Status:     "success",
				StatusCode: http.StatusFound,
				Message:    "Content fetched successfully.",
				Data:       fileMetaData,
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
