package handle

import (
	"context"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"log"
	"os"
	"tabManager/internal/define"
	"time"
)

func CreateToPDFTask(tab define.Tab) {

	now := time.Now()
	nowDate := now.Format("2006-01-02")
	timestamp := time.Now().Unix()

	var interfaces []interface{}
	task := define.Task{}
	task.TargetId = tab.Id
	task.CreateTime = nowDate
	task.TimeStamp = timestamp
	task.Name = "快照创建"
	task.TargetType = "tab_to_pdf"
	interfaces = append(interfaces, task)

	db.BatchExec(define.INSERT_TASK, interfaces)
}
func ToPDF(tab define.Tab) {
	// 创建 context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	// 生成pdf
	var buf []byte
	if err := chromedp.Run(ctx, printToPDF(tab.Url, &buf)); err != nil {
		log.Fatal(err)
	}
	if err := os.WriteFile("aa.pdf", buf, 0644); err != nil {
		log.Fatal(err)
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
