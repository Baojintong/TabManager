package main

import (
	"context"
	"encoding/json"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/labstack/gommon/log"
	"os"
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

func (a *App) ToPdf() H {
	// 创建 context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	// 生成pdf
	var buf []byte
	if err := chromedp.Run(ctx, printToPDF(`https://colobu.com/2021/05/05/generate-pdf-for-a-web-page-by-using-chromedp/`, &buf)); err != nil {
		log.Fatal(err)
	}
	if err := os.WriteFile("colobu.pdf", buf, 0644); err != nil {
		log.Fatal(err)
	}
	return M{
		"code": 200,
		"data": "success",
	}
}
func printToPDF(urlstr string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr), // 浏览指定的页面
		chromedp.ActionFunc(func(ctx context.Context) error {
			buf, _, err := page.PrintToPDF().WithPrintBackground(false).Do(ctx) // 通过cdp执行PrintToPDF
			if err != nil {
				return err
			}
			*res = buf
			return nil
		}),
	}
}
