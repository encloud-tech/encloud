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

func UploadContentCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upload",
		Short: "Upload content",
		Long: `Upload encrypted data to Filecoin.This command encrypts the specified file using a newly generated DEK. 
		The DEK is encrypted using the KEK and the metadata is stored on the local KV store`,
		Run: func(cmd *cobra.Command, args []string) {
			cfg, err := api.Fetch()
			if err != nil {
				fmt.Fprintf(cmd.OutOrStderr(), err.Error())
				os.Exit(-1)
			}

			kek := ""
			publicKey, _ := cmd.Flags().GetString("pubkey")
			path, _ := cmd.Flags().GetString("file")
			dekType, _ := cmd.Flags().GetString("type")
			readPublicKeyFromPath, _ := cmd.Flags().GetBool("read_pub_from_path")
			if readPublicKeyFromPath {
				kek = thirdparty.ReadKeyFile(publicKey)
			} else {
				kek = publicKey
			}
			uuid, err := api.Upload(path, cfg.Stat.KekType, dekType, kek)
			if err != nil {
				fmt.Fprintf(cmd.OutOrStderr(), err.Error())
				os.Exit(-1)
			} else {
				response := types.UploadContentResponse{
					Status:     "success",
					StatusCode: http.StatusCreated,
					Message:    "Content uploaded successfully.",
					Data:       types.Uuid{Uuid: uuid},
				}
				encoded, err := json.MarshalIndent(response, "", "    ")
				if err != nil {
					fmt.Fprintf(cmd.OutOrStderr(), err.Error())
					os.Exit(-1)
				}
				fmt.Fprintf(cmd.OutOrStdout(), string(encoded))
			}
		},
	}

	cmd.Flags().StringP("pubkey", "p", "", "KEK public key")
	cmd.Flags().BoolP("read_pub_from_path", "r", false, "Allows to read KEK public key from path")
	cmd.Flags().StringP("file", "f", "", "Upload file path")
	cmd.Flags().StringP("type", "t", "chacha20", "DEK type (chacha20/Aes)")
	cmd.MarkFlagRequired("pubkey")
	cmd.MarkFlagRequired("file")
	return cmd
}

func init() {
	RootCmd.AddCommand(UploadContentCmd())
}
