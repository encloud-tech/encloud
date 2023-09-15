package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/encloud-tech/encloud/pkg/api"
	"github.com/encloud-tech/encloud/pkg/types"
	thirdparty "github.com/encloud-tech/encloud/third_party"

	"github.com/spf13/cobra"
)

func RetrieveByUUIDCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "retrieve",
		Short: "Retrieve content by UUID",
		Long: `Retrieve data from Filecoin with a specific UUID. This command decrypts encrypted data on Filecoin using the relevant DEK.
		The DEK is stored in encrypted form in the metadata and is itself decrypted first using the KEK Private Key.`,
		Run: func(cmd *cobra.Command, args []string) {
			kek := ""
			privateKey := ""
			publicKey, _ := cmd.Flags().GetString("pubkey")
			pk, _ := cmd.Flags().GetString("privkey")
			uuid, _ := cmd.Flags().GetString("uuid")
			readPublicKeyFromPath, _ := cmd.Flags().GetBool("read_pub_from_path")
			readPrivateKeyFromPath, _ := cmd.Flags().GetBool("read_priv_from_path")
			retrievalFileStoragePath, _ := cmd.Flags().GetString("storage")
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

			fileMetaData, err := api.RetrieveByUUID(uuid, kek, privateKey, retrievalFileStoragePath)
			if err != nil {
				fmt.Fprintf(cmd.OutOrStderr(), err.Error())
				os.Exit(-1)
			}
			response := types.RetrieveByUUIDContentResponse{
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
	cmd.Flags().StringP("privkey", "k", "", "KEK private key")
	cmd.Flags().StringP("uuid", "u", "", "UUID of file to retrieve")
	cmd.Flags().StringP("storage", "s", "", "Path to store retrieved file")
	cmd.Flags().BoolP("read_pub_from_path", "r", false, "Allows to read KEK public key from path")
	cmd.Flags().BoolP("read_priv_from_path", "o", false, "Allows to read KEK private key from path")
	cmd.MarkFlagRequired("pubkey")
	cmd.MarkFlagRequired("privkey")
	cmd.MarkFlagRequired("uuid")
	return cmd
}

func init() {
	RootCmd.AddCommand(RetrieveByUUIDCmd())
}
