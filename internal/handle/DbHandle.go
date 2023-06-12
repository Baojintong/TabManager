package handle

import (
	"database/sql"
)

type DbHandle interface {
	Connect()
	Query(q string, args ...any) *sql.Rows
	Exec(sql_ string, args ...any)
	BatchExec(sql_ string, datas []interface{})
	Close()
}
