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

			response := types.RetrieveByUUIDContentResponse{
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

	cmd.Flags().StringP("pubkey", "p", "", "KEK public key")
	cmd.Flags().StringP("privkey", "k", "", "KEK private key")
	cmd.Flags().StringP("uuid", "u", "", "UUID of file to retrieve")
	cmd.Flags().StringP("email", "e", "", "Email to share file with")
	cmd.Flags().BoolP("read_pub_from_path", "r", false, "Allows to read KEK public key from path")
	cmd.Flags().BoolP("read_priv_from_path", "o", false, "Allows to read KEK private key from path")
	cmd.MarkFlagRequired("pubkey")
	cmd.MarkFlagRequired("privkey")
	cmd.MarkFlagRequired("uuid")
	cmd.MarkFlagRequired("email")
	return cmd
}

func init() {
	RootCmd.AddCommand(ShareCmd())
}
