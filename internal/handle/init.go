package handle

import "context"

var db DbHandle = new(DbHandleImpl)
var context_ *context.Context

func init() {
	db.Connect()
}

func SetContext(ctx context.Context) {
	context_ = &ctx
}
