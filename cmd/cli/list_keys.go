package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/encloud-tech/encloud/pkg/api"
	"github.com/encloud-tech/encloud/pkg/types"

	"github.com/spf13/cobra"
)

func ListKeysCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "keys",
		Short: "List keys",
		Long:  `List generated keys with number of associated files`,
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
