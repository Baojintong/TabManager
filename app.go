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

func (a *App) GetTabList(labelId uint32) H {
	tabList := handle.GetTabList(labelId)
	return M{
		"code": 200,
		"data": tabList,
	}
}

func (a *App) GetTab(tabId uint32) H {
	tab := handle.GetTab(tabId)
	return M{
		"code": 200,
		"data": tab,
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

func (a *App) CreateToPDFTask(item string) H {
	log.Info("CreateToPDFTask start.......")
	var tab define.Tab
	err := json.Unmarshal([]byte(item), &tab)
	if err != nil {
		return M{
			"code": -1,
			"msg":  "ERROR : " + err.Error(),
		}
	}
	handle.CreateToPDFTask(tab)
	return M{
		"code": 200,
		"data": "success",
	}
}

func (a *App) GetConfigList() H {
	configList := handle.GetConfigList()
	return M{
		"code": 200,
		"data": configList,
	}
}

func (a *App) SaveConfig(configs string) H {
	var configList []define.Config
	err := json.Unmarshal([]byte(configs), &configList)
	if err != nil {
		return M{
			"code": -1,
			"msg":  "ERROR : " + err.Error(),
		}
	}
	handle.UpdateConfig(configList)
	return M{
		"code": 200,
		"data": "",
	}
}