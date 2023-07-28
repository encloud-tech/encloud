package main

import (
	"encloud/pkg/api"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func RetrieveSharedContentCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "shared",
		Short: "Retrieve shared content",
		Long:  `Retrieve shared content from other users using your CID, DEK type and DEK`,
		Run: func(cmd *cobra.Command, args []string) {
			decryptedDekPath, _ := cmd.Flags().GetString("dek")
			cid, _ := cmd.Flags().GetString("cid")
			fileName, _ := cmd.Flags().GetString("name")
			dekType, _ := cmd.Flags().GetString("type")
			retrievalFileStoragePath, _ := cmd.Flags().GetString("storage")
			err := api.RetrieveSharedContent(decryptedDekPath, dekType, cid, fileName, retrievalFileStoragePath)
			if err != nil {
				fmt.Fprintf(cmd.OutOrStderr(), err.Error())
				os.Exit(-1)
			}
			fmt.Fprintf(cmd.OutOrStdout(), string("content downloaded successfully."))
		},
	}

	cmd.Flags().StringP("dek", "d", "", "Path to DEK file")
	cmd.Flags().StringP("cid", "c", "", "CID of shared file to retrieve")
	cmd.Flags().StringP("type", "t", "chacha20", "DEK type")
	cmd.Flags().StringP("name", "n", "", "Name of retrieved file")
	cmd.Flags().StringP("storage", "s", "", "Path to store retrieved file under")
	cmd.MarkFlagRequired("dek")
	cmd.MarkFlagRequired("cid")
	cmd.MarkFlagRequired("type")
	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("storage")
	return cmd
}

func init() {
	RootCmd.AddCommand(RetrieveSharedContentCmd())
}
