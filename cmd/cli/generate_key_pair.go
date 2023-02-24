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

func GenerateKeyPairCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generate-key-pair",
		Short: "Generate your key pair",
		Long:  `Generate your public key and private key which helps to encrypt and decrypt your data`,
		Run: func(cmd *cobra.Command, args []string) {
			cfg, err := api.Fetch()
			if err != nil {
				fmt.Fprintf(cmd.OutOrStderr(), err.Error())
				os.Exit(-1)
			}
			keys, err := api.GenerateKeyPair(cfg.Stat.KekType)
			if err != nil {
				fmt.Fprintf(cmd.OutOrStderr(), err.Error())
				os.Exit(-1)
			}
			response := types.GenerateKeyPairResponse{
				Status:     "success",
				StatusCode: http.StatusCreated,
				Message:    "Keys generated successfully.",
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
	RootCmd.AddCommand(GenerateKeyPairCmd())
}
