package cmd

import (
	"filecoin-encrypted-data-storage/config"
	"filecoin-encrypted-data-storage/service"
	thirdparty "filecoin-encrypted-data-storage/third_party"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func RetrieveSharedContentCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "retrieve-shared-content",
		Short: "Retrieve specific uploaded content using your cid",
		Long:  `Retrieve specific uploaded content using your cid and decrypt it using your private key`,
		Run: func(cmd *cobra.Command, args []string) {
			cfg, _ := config.LoadConf("./config.yml")
			estuaryService := service.New(cfg)

			decryptedDekPath, _ := cmd.Flags().GetString("dek")
			cid, _ := cmd.Flags().GetString("cid")

			dek := thirdparty.ReadFile(decryptedDekPath)
			log.Println(dek)

			filepath := estuaryService.DownloadContent(cid)
			err := thirdparty.DecryptFile(dek, filepath)
			if err != nil {
				fmt.Println(err)
			}
			os.Remove("assets/downloaded.bin")
			fmt.Fprintf(cmd.OutOrStdout(), string("content downloaded successfully."))
		},
	}

	cmd.Flags().StringP("dek", "d", "", "Enter your dek")
	cmd.Flags().StringP("cid", "c", "", "Enter your cid")
	cmd.MarkFlagRequired("dek")
	cmd.MarkFlagRequired("cid")
	return cmd
}

func init() {
	RootCmd.AddCommand(RetrieveSharedContentCmd())
}
