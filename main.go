package main

import (
	"embed"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"net/http"
	"tabManager/handle"
)

var assets embed.FS

func main() {
	go startServer()
	// Create an instance of the app structure
	app := NewApp()
	// Create application with options
	err := wails.Run(&options.App{
		Title:         "TabManager",
		Width:         1024,
		Height:        768,
		DisableResize: true,
		Fullscreen:    false,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 30, G: 30, B: 38, A: 100},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})
	if err != nil {
		println("start app Error:", err.Error())
	}
}

func startServer() {
	http.HandleFunc("/tabs", handle.TabHandler)
	err := http.ListenAndServe("localhost:12315", nil)
	if err != nil {
		println(" startServer Error:", err.Error())
	}
}
