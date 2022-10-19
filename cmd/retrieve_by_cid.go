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
	"sync"

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

			kek, _ := cmd.Flags().GetString("publicKey")
			privateKey, _ := cmd.Flags().GetString("privateKey")
			uuid, _ := cmd.Flags().GetString("uuid")
			fileMetaData := service.FetchByCid(kek + ":" + uuid)
			decryptedDek, err := thirdparty.DecryptWithRSA(fileMetaData.Dek, thirdparty.GetIdRsaFromStr(privateKey))
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(fileMetaData)
			filepath := "assets/downloaded.bin"
			var wg sync.WaitGroup
			// limit to four downloads at a time, this is called a semaphore
			limiter := make(chan struct{}, 4)
			for i, link := range fileMetaData.Cid {
				wg.Add(1)
				go estuaryService.DownloadContent(&wg, limiter, i, link, filepath)
			}
			wg.Wait()
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
	cmd.MarkFlagRequired("publicKey")
	cmd.MarkFlagRequired("privateKey")
	cmd.MarkFlagRequired("uuid")
	return cmd
}

func init() {
	RootCmd.AddCommand(RetrieveByCidCmd())
}
