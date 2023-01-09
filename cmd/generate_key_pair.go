package cmd

import (
	"encloud/config"
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
			cfg, _ := config.LoadConf("./config.yaml")
			var keys types.Keys
			if cfg.Stat.KekType == "rsa" {
				thirdparty.InitCrypto()
				keys = types.Keys{PublicKey: thirdparty.GetIdRsaPubStr(), PrivateKey: thirdparty.GetIdRsaStr()}
				os.Remove(".keys/.idRsaPub")
				os.Remove(".keys/.idRsa")
			} else if cfg.Stat.KekType == "ecies" {
				k := thirdparty.EciesGenerateKeyPair()
				keys = types.Keys{PublicKey: k.PublicKey.Hex(false), PrivateKey: k.Hex()}
			} else {
				fmt.Fprintf(cmd.OutOrStderr(), "Invalid argument")
				return
			}
			response := types.GenerateKeyPairResponse{
				Status:     "success",
				StatusCode: http.StatusCreated,
				Message:    "Keys generated successfully.",
				Data:       keys,
			}
			encoded, err := json.MarshalIndent(response, "", "    ")
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
