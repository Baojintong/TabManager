package handle

import (
	"context"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"log"
	"os"
	"strconv"
	"tabManager/internal/define"
	"time"
)

var channel = make(chan define.Tab, 64)

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

	channel <- tab
}
func ToPDFConsumer() {
	log.Println("启动ToPDFConsumer.....")
	for tab := range channel {
		log.Println("触发ToPDFConsumer.....")
		// 创建 context
		ctx, cancel := chromedp.NewContext(context.Background())
		defer cancel()
		// 生成pdf
		var buf []byte
		if err := chromedp.Run(ctx, printToPDF(tab.Url, &buf)); err != nil {
			log.Fatal(err)
		}
		name := strconv.FormatUint(uint64(tab.Id), 10) + ".pdf"
		if err := os.WriteFile(name, buf, 0644); err != nil {
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
