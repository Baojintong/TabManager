package handle

import (
	"context"
	"github.com/labstack/gommon/log"
)

var db DbHandle = new(DbHandleImpl)
var context_ *context.Context

func init() {
	db.Init()
	db.Connect()
	log.Info("table start init.........")
	CreateLabelTable()
	CreateTabTable()
	CreateTasktable()
	log.Info("table end init.........")
}

func SetContext(ctx context.Context) {
	context_ = &ctx
}
