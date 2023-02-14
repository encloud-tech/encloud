package main

import (
	"encloud/config"
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
	cfg, err := config.LoadConf("./config.yaml")
	if err != nil {
		// Load default configuration from config.go file if config.yaml file not found
		cfg, _ = config.LoadConf()
	}
	cmd := &cobra.Command{
		Use:   "upload",
		Short: "Upload your content to filecoin storage",
		Long:  `Upload your content to filecoin storage which is encrypted using your public key`,
		Run: func(cmd *cobra.Command, args []string) {
			kek := ""
			publicKey, _ := cmd.Flags().GetString("publicKey")
			path, _ := cmd.Flags().GetString("filePath")
			dekType, _ := cmd.Flags().GetString("dekType")
			readPublicKeyFromPath, _ := cmd.Flags().GetBool("readPublicKeyFromPath")
			if readPublicKeyFromPath {
				kek = thirdparty.ReadKeyFile(publicKey)
			} else {
				kek = publicKey
			}
			uuid, err := api.Upload(path, cfg.Stat.KekType, dekType, kek)
			if err != nil {
				fmt.Fprintf(cmd.OutOrStderr(), err.Error())
				os.Exit(-1)
			}
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
		},
	}

	cmd.Flags().StringP("publicKey", "p", "", "Enter your public key")
	cmd.Flags().BoolP("readPublicKeyFromPath", "r", false, "Do you want public key read from path you have entered?")
	cmd.Flags().StringP("filePath", "f", "", "Enter your file path")
	cmd.Flags().StringP("dekType", "e", "chacha20", "Enter which type of encryption do you want?")
	cmd.MarkFlagRequired("publicKey")
	cmd.MarkFlagRequired("filepath")
	return cmd
}

func init() {
	RootCmd.AddCommand(UploadContentCmd())
}
