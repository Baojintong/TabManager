package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"io"
	"net/http"
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

type TabsData struct {
	Title   string `json:"title"`
	Url     string `json:"url"`
	IconUrl string `json:"iconUrl"`
}

func startServer() {
	http.HandleFunc("/tabs", tabsHandler)
	err := http.ListenAndServe("127.0.0.1:12315", nil)
	if err != nil {
		println(" startServer Error:", err.Error())
	}
}

func tabsHandler(w http.ResponseWriter, r *http.Request) {
	var body, err = io.ReadAll(r.Body)
	if err != nil {
		println(" tabsHandler Error:", err.Error())
	}
	var tabsData []TabsData
	var jsonStr = string(body)
	err = json.Unmarshal([]byte(jsonStr), &tabsData)
	if err != nil {
		println(" tabsHandler Error:", err.Error())
		return
	}

	for i, n := 0, len(tabsData); i < n; i++ { // 常见的 for 循环，支持初始化语句。
		fmt.Printf("IconUrl value %s\n", tabsData[i].IconUrl)
		fmt.Printf("Title value %s\n", tabsData[i].Title)
		fmt.Printf("Url value %s\n", tabsData[i].Url)
	}
	// 回复
	_, err = w.Write([]byte("success"))
	if err != nil {
		println(" tabsHandler resp Error:", err.Error())
	}
}
