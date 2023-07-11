package main

import (
	"embed"
	"github.com/labstack/gommon/log"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"net/http"
	"tabManager/internal/define"
	"tabManager/internal/handle"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	go startServer()
	go handle.ToPDFConsumer()
	//go toPdf()
	//createDataBase()
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
		log.Error("start app Error:", err)
	}
}

func startServer() {
	log.Info("startServer.......")
	http.HandleFunc("/tabs", handle.TabHandler)
	err := http.ListenAndServe(define.LISTEN_URL, nil)
	if err != nil {
		log.Error(" startServer Error:")
	}
}
