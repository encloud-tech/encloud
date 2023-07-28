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

func ShareCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "share",
		Short: "Share content",
		Long:  `Share your files with other users using the UUID and DEK`,
		Run: func(cmd *cobra.Command, args []string) {
			kek := ""
			privateKey := ""
			publicKey, _ := cmd.Flags().GetString("pubkey")
			pk, _ := cmd.Flags().GetString("privkey")
			uuid, _ := cmd.Flags().GetString("uuid")
			email, _ := cmd.Flags().GetString("email")
			readPublicKeyFromPath, _ := cmd.Flags().GetBool("read_pub_from_path")
			readPrivateKeyFromPath, _ := cmd.Flags().GetBool("read_priv_from_path")
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

			fileMetaData, err := api.Share(uuid, kek, privateKey, email)
			if err != nil {
				fmt.Fprintf(cmd.OutOrStderr(), err.Error())
				os.Exit(-1)
			}

			response := types.RetrieveByCIDContentResponse{
				Status:     "success",
				StatusCode: http.StatusFound,
				Message:    "Content shared successfully.",
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
	cmd.Flags().StringP("privateKey", "k", "", "Enter your private key")
	cmd.Flags().StringP("uuid", "u", "", "Enter your uuid")
	cmd.Flags().StringP("email", "e", "", "Enter email which you want to share")
	cmd.Flags().BoolP("readPublicKeyFromPath", "r", false, "Do you want public key read from path you have entered?")
	cmd.Flags().BoolP("readPrivateKeyFromPath", "o", false, "Do you want private key read from path you have entered?")
	cmd.MarkFlagRequired("publicKey")
	cmd.MarkFlagRequired("privateKey")
	cmd.MarkFlagRequired("uuid")
	cmd.MarkFlagRequired("email")
	return cmd
}

func init() {
	RootCmd.AddCommand(ShareCmd())
}
