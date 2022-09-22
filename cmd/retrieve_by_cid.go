package cmd

import (
	"encoding/json"
	"filecoin-encrypted-data-storage/config"
	"filecoin-encrypted-data-storage/service"
	thirdparty "filecoin-encrypted-data-storage/third_party"
	"filecoin-encrypted-data-storage/types"
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var retrieveByCidCmd = &cobra.Command{
	Use:   "retrieve-by-cid",
	Short: "Retrieve specific uploaded content using your cid",
	Long:  `Retrieve specific uploaded content using your cid and decrypt it using your private key`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg, _ := config.LoadConf("config.yml")
		estuaryService := service.New(cfg)

		kek, _ := cmd.Flags().GetString("publicKey")
		privateKey, _ := cmd.Flags().GetString("privateKey")
		cid, _ := cmd.Flags().GetString("cid")
		fileMetaData := service.FetchByCid(kek + "-" + cid)
		decryptedDek, err := thirdparty.DecryptWithRSA(fileMetaData.Dek, thirdparty.GetIdRsaFromStr(privateKey))
		if err != nil {
			fmt.Println(err)
		}
		filepath := estuaryService.DownloadContent(fileMetaData.Cid)
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
		fmt.Println(string(encoded))
	},
}

func init() {
	rootCmd.AddCommand(retrieveByCidCmd)
	retrieveByCidCmd.Flags().StringP("publicKey", "p", "", "Enter your public key")
	retrieveByCidCmd.Flags().StringP("privateKey", "k", "", "Enter your private key")
	retrieveByCidCmd.Flags().StringP("cid", "c", "", "Enter your cid")
	retrieveByCidCmd.MarkFlagRequired("publicKey")
	retrieveByCidCmd.MarkFlagRequired("privateKey")
	retrieveByCidCmd.MarkFlagRequired("cid")
}
