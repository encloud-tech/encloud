package cmd

import (
	"encloud/service"
	thirdparty "encloud/third_party"
	"encloud/types"
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
			fileMetaData := service.Fetch(kek)
			os.Remove("assets/downloaded.bin")
			response := types.ListContentResponse{
				Status:     "success",
				StatusCode: http.StatusFound,
				Message:    "Content fetched successfully.",
				Data:       fileMetaData,
			}
			encoded, err := json.MarshalIndent(response, "", "    ")
			if err != nil {
				fmt.Println(err)
				return
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
