package cmd

import (
	thirdparty "encloud/third_party"
	"encloud/types"
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
			thirdparty.InitCrypto()
			response := types.GenerateKeyPairResponse{
				Status:     "success",
				StatusCode: http.StatusCreated,
				Message:    "Keys generated successfully.",
				Data:       types.Keys{PublicKey: thirdparty.GetIdRsaPubStr(), PrivateKey: thirdparty.GetIdRsaStr()},
			}
			os.Remove(".keys/.idRsaPub")
			os.Remove(".keys/.idRsa")
			encoded, err := json.Marshal(response)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Fprintf(cmd.OutOrStdout(), string(encoded))
		},
	}

	return cmd
}

func init() {
	RootCmd.AddCommand(GenerateKeyPairCmd())
}
