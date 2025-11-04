package main

import (
	"embed"
	"log"
	"os"
	"strings"

	"github.com/kishansakhiya/wails-demo/backend/app"
	"github.com/kishansakhiya/wails-demo/backend/app/utils"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	appInstance := app.NewApp()

	// Check for URL scheme arguments (Windows/Linux)
	// On Windows/Linux, custom URLs are passed as command-line arguments
	if len(os.Args) > 1 {
		arg := os.Args[1]
		// Check if it's a custom URL scheme
		if strings.HasPrefix(arg, "wails-demo://") {
			appInstance.OnURL(arg)
		}
	}

	// Create application with options
	err := wails.Run(&options.App{
		Title:             utils.AppName,
		Width:             1280,
		Height:            800,
		MinWidth:          0,
		MinHeight:         0,
		HideWindowOnClose: false,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        appInstance.Startup,
		OnDomReady:       appInstance.DomReady,
		OnShutdown:       appInstance.Shutdown,
		Bind: []any{
			appInstance,
		},
		// macOS: Handle URL scheme via OnUrlOpen callback
		Mac: &mac.Options{
			OnUrlOpen: func(url string) {
				appInstance.OnURL(url)
			},
		},
		Windows: &windows.Options{
			WebviewIsTransparent: false,
			WindowIsTranslucent:  false,
			DisableWindowIcon:    false,
		},
		Debug: options.Debug{
			OpenInspectorOnStartup: false,
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}
