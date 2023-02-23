package main

import (
	"encloud/pkg/api"
	"encloud/pkg/types"
	thirdparty "encloud/third_party"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

func ListContentsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List your uploaded contents",
		Long:  `List your uploaded contents data which contains file meta informations`,
		Run: func(cmd *cobra.Command, args []string) {
			kek := ""
			publicKey, _ := cmd.Flags().GetString("publicKey")
			readPublicKeyFromPath, _ := cmd.Flags().GetBool("readPublicKeyFromPath")
			if readPublicKeyFromPath {
				kek = thirdparty.ReadKeyFile(publicKey)
			} else {
				kek = publicKey
			}

			fileMetaData := api.List(kek)

			response := types.ListContentResponse{
				Status:     "success",
				StatusCode: http.StatusFound,
				Message:    "Content fetched successfully.",
				Data:       fileMetaData,
			}
			encoded, err := json.MarshalIndent(response, "", "    ")
			if err != nil {
				fmt.Fprintf(cmd.OutOrStderr(), err.Error())
				os.Exit(-1)
			}
			fmt.Fprintf(cmd.OutOrStdout(), string(encoded))
		},
	}

	cmd.Flags().StringP("publicKey", "p", "", "Enter your public key")
	cmd.Flags().BoolP("readPublicKeyFromPath", "r", false, "Do you want public key read from path you have entered?")
	cmd.MarkFlagRequired("publicKey")
	return cmd
}

func init() {
	RootCmd.AddCommand(ListContentsCmd())
}