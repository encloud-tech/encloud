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
		Use:   "update-config",
		Short: "Update default configurations",
		Long:  `Modify configuration as per your needs`,
		Run: func(cmd *cobra.Command, args []string) {
			configFilePath, _ := cmd.Flags().GetString("configFilePath")
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

	cmd.Flags().StringP("configFilePath", "f", "", "Please enter path of config file.")
	cmd.MarkFlagRequired("configFilePath")
	return cmd
}

func init() {
	RootCmd.AddCommand(ConfigCmd())
}
