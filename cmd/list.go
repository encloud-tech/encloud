package cmd

import (
	"encoding/json"
	"filecoin-encrypted-data-storage/service"
	"filecoin-encrypted-data-storage/types"
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
			kek, _ := cmd.Flags().GetString("publicKey")
			fileMetaData := service.Fetch(kek)
			os.Remove("assets/downloaded.bin")
			response := types.ListContentResponse{
				Status:     "success",
				StatusCode: http.StatusFound,
				Message:    "Content fetched successfully.",
				Data:       fileMetaData,
			}
			encoded, err := json.Marshal(response)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Fprintf(cmd.OutOrStdout(), string(encoded))
		},
	}

	cmd.Flags().StringP("publicKey", "p", "", "Enter your public key")
	cmd.MarkFlagRequired("publicKey")
	return cmd
}

func init() {
	RootCmd.AddCommand(ListContentsCmd())
}
