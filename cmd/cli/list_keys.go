package main

import (
	"encloud/pkg/api"
	"encloud/pkg/types"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

func ListKeysCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-keys",
		Short: "List your all keys",
		Long:  `List your all keys with files count`,
		Run: func(cmd *cobra.Command, args []string) {
			keys, err := api.ListKeys()
			if err != nil {
				fmt.Fprintf(cmd.OutOrStderr(), err.Error())
				os.Exit(-1)
			}
			response := types.ListKeysResponse{
				Status:     "success",
				StatusCode: http.StatusCreated,
				Message:    "Keys fetched successfully.",
				Data:       keys,
			}
			encoded, err := json.MarshalIndent(response, "", "    ")
			if err != nil {
				fmt.Fprintf(cmd.OutOrStderr(), err.Error())
				os.Exit(-1)
			}
			fmt.Fprintf(cmd.OutOrStdout(), string(encoded))
		},
	}

	return cmd
}

func init() {
	RootCmd.AddCommand(ListKeysCmd())
}
