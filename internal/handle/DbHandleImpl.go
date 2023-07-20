package handle

import (
	"database/sql"
	"encoding/json"
	"github.com/labstack/gommon/log"
	"os"
	"reflect"
	"tabManager/internal/define"
)

type DbHandleImpl struct {
	db *sql.DB
}
type Param struct {
	Value any
}

func (db *DbHandleImpl) Init() {
	if _, err := os.Stat(define.DB_ALL); os.IsNotExist(err) {
		if err := os.MkdirAll(define.DB_PATH, 0777); err != nil {
			panic(err)
		}
		if _, err := os.Create(define.DB_ALL); err != nil {
			panic(err)
		}
	}
}
func (db *DbHandleImpl) Connect() {
	if _, err := os.Stat(define.DB_PATH); os.IsNotExist(err) {
		log.Error("db未初始化")
		panic(err)
	}
	var err error
	db.db, err = sql.Open("sqlite3", define.DATA_SOURCE_NAME)
	if err != nil {
		panic(err)
	}
}
func (db *DbHandleImpl) Query(q string, args ...any) *sql.Rows {
	rows, err := db.db.Query(q, args...)
	if err != nil {
		panic(err)
	}
	return rows
}

func (db *DbHandleImpl) QueryRow(q string, args ...any) *sql.Row {
	row := db.db.QueryRow(q, args...)
	return row
}
func (db *DbHandleImpl) BatchExec(sql_ string, datas []interface{}) {
	transaction(db.db, func(tx *sql.Tx) {
		stmt, err := tx.Prepare(sql_)
		if err != nil {
			panic(err)
		}
		log.Info("sql_:", sql_)
		for _, v := range datas {
			value := reflect.ValueOf(v)
			t := value.Type()
			var nameList []interface{}
			for i := 0; i < t.NumField(); i++ {
				field := t.Field(i)
				name := field.Name
				tag := field.Tag.Get("db")
				fieldValue := value.FieldByName(name)
				fieldValues := fieldValue.Interface()
				if len(tag) != 0 {
					var named sql.NamedArg = sql.Named(tag, fieldValues)
					nameList = append(nameList, named)
				}
			}
			execStmt(stmt, nameList...)
		}
	})
}

func (db *DbHandleImpl) Exec(sql_ string, args ...any) {
	transaction(db.db, func(tx *sql.Tx) {
		stmt, err := tx.Prepare(sql_)
		if err != nil {
			panic(err)
		}
		result := execStmt(stmt, args...)
		count, _ := result.RowsAffected()
		log.Info("sql_:", sql_)
		log.Info("exec count:", count)

		err = stmt.Close()
		if err != nil {
			panic(err)
		}
	})
}

func (db *DbHandleImpl) ExecNoTran(sql_ string) {
	var _, err = db.db.Exec(sql_)
	if err != nil {
		panic(err)
	}
}

func execStmt(stmt *sql.Stmt, args ...any) sql.Result {
	jsonBytes, _ := json.Marshal(args)
	log.Info("execStmt:",string(jsonBytes))
	result, err := stmt.Exec(args...)
	if err != nil {
		return nil
	}
	return result
}

func transaction(db *sql.DB, fn func(tx *sql.Tx)) {
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	defer tx.Rollback()
	fn(tx)
	err = tx.Commit()
	if err != nil {
		panic(err)
	}
}
func (db *DbHandleImpl) Close() {
	err := db.db.Close()
	if err != nil {
		panic(err)
	}
}
