package handle

import (
	"database/sql"
)

type DbHandle interface {
	Init()
	Connect()
	Query(q string, args ...any) *sql.Rows
	QueryRow(q string, args ...any) *sql.Row
	Exec(sql_ string, args ...any)

	ExecNoTran(sql_ string)
	BatchExec(sql_ string, datas []interface{})
	Close()
}
