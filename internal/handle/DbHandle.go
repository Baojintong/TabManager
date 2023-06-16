package handle

import (
	"database/sql"
)

type DbHandle interface {
	Connect()
	Query(q string, args ...any) *sql.Rows
	QueryRow(q string, args ...any) *sql.Row
	Exec(sql_ string, args ...any)
	BatchExec(sql_ string, datas []interface{})
	Close()
}
