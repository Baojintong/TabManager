package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/labstack/gommon/log"
	"tabManager/internal/define"
	"tabManager/internal/handle"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) GetTabList() H {
	tabs, err := handle.QueryAllTabs()
	if err != nil {
		return M{
			"code": -1,
			"msg":  "ERROR : " + err.Error(),
		}
	}
	return M{
		"code": 200,
		"data": tabs,
	}
}

func (a *App) UpdateTab(item string) H {
	log.Info("UpdateTab start.......")
	var tab define.TabsData
	err := json.Unmarshal([]byte(item), &tab)
	if err != nil {
		log.Error(" UpdateTab Error:", err)
	}
	handle.UpdateTab(tab)
	return M{
		"code": 200,
		"data": "",
	}
}

func (a *App) DeleteTab(item string) H {
	log.Info("DeleteTab start.......")
	var tab define.TabsData
	err := json.Unmarshal([]byte(item), &tab)
	if err != nil {
		log.Error(" DeleteTab Error:", err)
	}
	handle.DeleteTab(tab)
	return M{
		"code": 200,
		"data": "",
	}
}
