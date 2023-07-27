package handle

import (
	"context"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/labstack/gommon/log"
	"os"
	"path/filepath"
	"strconv"
	"tabManager/internal/define"
	"tabManager/internal/utils"
)

var channel = make(chan define.Tab, 64)

func CreateToPDFTask(tab define.Tab) {

	timestamp,nowDate := utils.GetCurrentTime()

	var interfaces []interface{}
	task := define.Task{}
	task.TargetId = tab.Id
	task.CreateTime = nowDate
	task.TimeStamp = timestamp
	task.Name = "快照创建"
	task.TargetType = "tab_to_pdf"
	interfaces = append(interfaces, task)

	db.BatchExec(define.INSERT_TASK, interfaces)

	channel <- tab
}
func ToPDFConsumer() {
	log.Info("启动ToPDFConsumer.....")
	for tab := range channel {
		log.Info("触发ToPDFConsumer.....")
		// 创建 context
		ctx, cancel := chromedp.NewContext(context.Background())
		defer cancel()
		// 生成pdf
		var buf []byte
		if err := chromedp.Run(ctx, printToPDF(tab.Url, &buf)); err != nil {
			log.Fatal(err)
		}
		name := strconv.FormatUint(uint64(tab.Id), 10) + ".pdf"
		dirPath := GetConfigValueByKey("path")+"/pdfs/"
		filePath := filepath.Join(dirPath, name)

		if _, err := os.Stat(dirPath); os.IsNotExist(err) {
			err := os.MkdirAll(dirPath, 0777)
			if err != nil {
				log.Fatal(err)
			}
		}

		if err := os.WriteFile(filePath, buf, 0644); err != nil {
			log.Fatal(err)
		}
	}
}

func printToPDF(urlstr string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.ActionFunc(func(ctx context.Context) error {
			buf, _, err := page.PrintToPDF().WithPrintBackground(false).Do(ctx)
			if err != nil {
				return err
			}
			*res = buf
			return nil
		}),
	}
}
