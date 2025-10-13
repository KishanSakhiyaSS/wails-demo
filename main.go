package main

import (
	"embed"

	"github.com/kishansakhiya/wails-demo/backend/app"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	a := app.NewApp()
	err := wails.Run(&options.App{
		Title:     "Wails Demo",
		Width:     1024,
		Height:    768,
		OnStartup: a.Startup,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Bind: []any{
			a.UserHandler,
			a.SystemHandler,
		},
	})
	if err != nil {
		println("Error:", err.Error())
	}
}
