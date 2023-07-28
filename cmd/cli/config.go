package main

import (
	"encloud/pkg/api"
	"encloud/pkg/types"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

func ConfigCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Update app configurations",
		Long:  `Update configurations for the application using a compatible yaml file`,
		Run: func(cmd *cobra.Command, args []string) {
			configFilePath, _ := cmd.Flags().GetString("path")
			data, err := ioutil.ReadFile(configFilePath)
			if err != nil {
				fmt.Fprintf(cmd.OutOrStderr(), err.Error())
				os.Exit(-1)
			}

			var conf types.ConfYaml
			if err := yaml.Unmarshal(data, &conf); err != nil {
				fmt.Fprintf(cmd.OutOrStderr(), err.Error())
				os.Exit(-1)
			}

			err = api.Store(conf)
			if err != nil {
				fmt.Fprintf(cmd.OutOrStderr(), err.Error())
				os.Exit(-1)
			}

			response := types.ConfigResponse{
				Status:     "success",
				StatusCode: http.StatusFound,
				Message:    "Configuration updated successfully.",
				Data:       conf,
			}
			encoded, err := json.MarshalIndent(response, "", "    ")
			if err != nil {
				fmt.Fprintf(cmd.OutOrStderr(), err.Error())
				os.Exit(-1)
			}
			fmt.Fprintf(cmd.OutOrStdout(), string(encoded))
		},
	}

	cmd.Flags().StringP("path", "p", "", "Path to config file")
	cmd.MarkFlagRequired("path")
	return cmd
}

func init() {
	RootCmd.AddCommand(ConfigCmd())
}
