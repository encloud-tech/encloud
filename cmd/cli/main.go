package main

import (
	"encloud/config"
	"log"

	"github.com/adrg/xdg"
)

func main() {
	configFilePath, err := xdg.SearchConfigFile("encloud/config.yaml")
	if err != nil {
		confErr := config.LoadDefaultConf()
		if confErr != nil {
			log.Println("Load default config error:", confErr.Error())
		}
	}

	log.Println("Config file was found at:", configFilePath)

	Execute()
}
