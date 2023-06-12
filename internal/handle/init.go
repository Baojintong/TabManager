package handle

var db DbHandle = new(DbHandleImpl)
func init(){
	db.Connect()
}
