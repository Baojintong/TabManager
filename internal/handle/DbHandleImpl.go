package handle

import (
	"database/sql"
	"encoding/json"
	"github.com/labstack/gommon/log"
	"os"
	"reflect"
	"regexp"
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
	placeholderNames := GetPlaceholderNames(sql_)
	transaction(db.db, func(tx *sql.Tx) {
		stmt, err := tx.Prepare(sql_)
		if err != nil {
			panic(err)
		}
		for _, data := range datas {
			nameList := BuildNameListByNames(placeholderNames, data)
			var args []interface{}
			for _, namedArg := range nameList {
				args = append(args, namedArg)
			}
			execStmt(stmt, args...)
		}
	})
}

func GetPlaceholderNames(sql string) []string {
	re := regexp.MustCompile(`:(\w+)`)
	names := re.FindAllStringSubmatch(sql, -1)
	var result []string
	for _, name := range names {
		result = append(result, name[1])
	}
	return result
}

func BuildNameListByNames(names []string, data interface{}) []sql.NamedArg {
	namedArgs := make([]sql.NamedArg, 0)

	// 使用 reflect 获取数据结构
	params := getDbFields(data)
	for _, name := range names {
		field := params[name]
		log.Info("name:", name, " value:", field)
		namedArgs = append(namedArgs, sql.Named(name, field))
	}

	return namedArgs
}

func getDbFields(data interface{}) map[string]interface{} {
	params := make(map[string]interface{})
	v := reflect.ValueOf(data)
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("db")

		if tag != "" {
			fieldValue := v.Field(i).Interface()
			params[tag] = fieldValue
		}
	}

	return params
}

func (db *DbHandleImpl) Exec(sql_ string, args ...any) {
	transaction(db.db, func(tx *sql.Tx) {
		stmt, err := tx.Prepare(sql_)
		if err != nil {
			panic(err)
		}
		execStmt(stmt, args...)
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
	log.Info("execStmt:", string(jsonBytes))
	result, err := stmt.Exec(args...)
	if err != nil {
		panic(err)
	}
	count, _ := result.RowsAffected()
	log.Info("args num:", len(args))
	log.Info("exec count:", count)
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
