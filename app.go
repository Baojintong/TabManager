package main

import (
	"context"
	"encoding/json"
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

// startup is called when the app starts. The context is saved,
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	handle.SetContext(ctx)
}

func (a *App) GetTabList() H {
	tabList := handle.GetTabList()
	return M{
		"code": 200,
		"data": tabList,
	}
}

func (a *App) UpdateTab(item string) H {
	log.Info("UpdateTab start.......:", item)
	var tab define.Tab
	err := json.Unmarshal([]byte(item), &tab)
	if err != nil {
		return M{
			"code": -1,
			"msg":  "ERROR : " + err.Error(),
		}
	}
	handle.UpdateTab(tab)
	return M{
		"code": 200,
		"data": "",
	}
}

func (a *App) DeleteTab(item string) H {
	log.Info("DeleteTab start.......")
	var tab define.Tab
	err := json.Unmarshal([]byte(item), &tab)
	if err != nil {
		return M{
			"code": -1,
			"msg":  "ERROR : " + err.Error(),
		}
	}
	handle.DeleteTab(tab)
	return M{
		"code": 200,
		"data": "",
	}
}

func (a *App) SaveLabel(labels string) H {
	log.Info("SaveLabel start.......")
	var label define.Label
	err := json.Unmarshal([]byte(labels), &label)
	if err != nil {
		return M{
			"code": -1,
			"msg":  "ERROR : " + err.Error(),
		}
	}
	handle.SaveLabel(label)
	return M{
		"code": 200,
		"data": "",
	}
}

func (a *App) GetLabelList() H {
	labelList := handle.GetLabelList()
	return M{
		"code": 200,
		"data": labelList,
	}
}

func (a *App) GetTabLabelList(tagId uint32) H {
	list := handle.QueryTabLabel(tagId)
	return M{
		"code": 200,
		"data": list,
	}
}
