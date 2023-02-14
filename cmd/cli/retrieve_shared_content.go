package main

import (
	"encloud/pkg/api"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func RetrieveSharedContentCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "retrieve-shared-content",
		Short: "Retrieve specific uploaded content using your cid",
		Long:  `Retrieve specific uploaded content using your cid and decrypt it using your private key`,
		Run: func(cmd *cobra.Command, args []string) {
			decryptedDekPath, _ := cmd.Flags().GetString("dek")
			cid, _ := cmd.Flags().GetString("cid")
			dekType, _ := cmd.Flags().GetString("dekType")
			err := api.RetrieveSharedContent(decryptedDekPath, dekType, cid)
			if err != nil {
				fmt.Fprintf(cmd.OutOrStderr(), err.Error())
				os.Exit(-1)
			}
			fmt.Fprintf(cmd.OutOrStdout(), string("content downloaded successfully."))
		},
	}

	cmd.Flags().StringP("dek", "d", "", "Enter your dek path")
	cmd.Flags().StringP("cid", "c", "", "Enter your cid")
	cmd.Flags().StringP("dekType", "e", "", "Enter DEK type")
	cmd.MarkFlagRequired("dek")
	cmd.MarkFlagRequired("cid")
	cmd.MarkFlagRequired("dekType")
	return cmd
}

func init() {
	RootCmd.AddCommand(RetrieveSharedContentCmd())
}
