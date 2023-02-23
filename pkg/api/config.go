package api

import (
	"encloud/pkg/types"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/adrg/xdg"
	"gopkg.in/yaml.v2"
)

func Store(conf types.ConfYaml) error {
	configFilePath, err := xdg.ConfigFile("encloud/config.yaml")
	if err != nil {
		return err
	}
	log.Println("Save the config file at:", configFilePath)

	yamlData, err := yaml.Marshal(&conf)

	err = ioutil.WriteFile(configFilePath, yamlData, 0644)
	if err != nil {
		return err
	}

	return err
}

func Fetch() (types.ConfYaml, error) {
	configFilePath, err := xdg.SearchConfigFile("encloud/config.yaml")
	if err != nil {
		log.Println("Error :", err.Error())
		os.Exit(-1)
		return types.ConfYaml{}, err
	}
	log.Println("Serach config file at:", configFilePath)

	data, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		log.Println("Error :", err.Error())
		os.Exit(-1)
		return types.ConfYaml{}, err
	}

	var conf types.ConfYaml
	if err := yaml.Unmarshal(data, &conf); err != nil {
		log.Println("Error :", err.Error())
		os.Exit(-1)
		return types.ConfYaml{}, err
	}

	fmt.Println(conf)
	return conf, nil
}
