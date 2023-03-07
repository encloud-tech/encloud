package main

import (
	"embed"
	"encloud/config"
	"log"

	"github.com/adrg/xdg"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	_, confErr := xdg.SearchConfigFile("encloud/config.yaml")
	if confErr != nil {
		loadErr := config.LoadDefaultConf()
		if loadErr != nil {
			log.Println("Load default config error:", loadErr.Error())
		}
	}

	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "Encloud",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
