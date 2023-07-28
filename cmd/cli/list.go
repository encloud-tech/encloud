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
		Use:   "contents",
		Short: "List contents",
		Long:  `List uploaded files and associated metadata`,
		Run: func(cmd *cobra.Command, args []string) {
			kek := ""
			publicKey, _ := cmd.Flags().GetString("pubkey")
			readPublicKeyFromPath, _ := cmd.Flags().GetBool("read_pub_from_path")
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

	cmd.Flags().StringP("pubkey", "p", "", "KEK public key")
	cmd.Flags().BoolP("read_pub_from_path", "r", false, "Allows to read KEK public key from path")
	cmd.MarkFlagRequired("pubkey")
	return cmd
}

func init() {
	RootCmd.AddCommand(ListContentsCmd())
}
