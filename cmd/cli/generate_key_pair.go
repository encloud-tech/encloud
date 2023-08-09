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
		Use:   "keygen",
		Short: "Generate new key pair",
		Long:  `Generate ECIES secp256k1 OR RSA 2048 key pair to encrypt & decrypt the AES-256 keys`,
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
