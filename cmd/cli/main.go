package main

import (
	"log"

	"github.com/encloud-tech/encloud/config"

	"github.com/adrg/xdg"
	"github.com/spf13/cobra/doc"
)

func main() {
	_, err := xdg.SearchConfigFile("encloud/config.yaml")
	if err != nil {
		confErr := config.LoadDefaultConf()
		if confErr != nil {
			log.Println("Load default config error:", confErr.Error())
		}
	}

	Execute()

	err = doc.GenMarkdownTree(RootCmd, "./../../docs")
	if err != nil {
		log.Fatal(err)
	}
}
