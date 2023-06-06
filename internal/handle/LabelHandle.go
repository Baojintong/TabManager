package handle

import (
	"database/sql"
	"github.com/labstack/gommon/log"
	"tabManager/internal/define"
)

func SaveLabel(label define.Label) {
	createLabelTable()
	saveLabel(label)
}

func saveLabel(label define.Label) {
	tx, err := db.Begin()
	if err != nil {
		log.Error("SaveLabel Begin Error:", err)
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("INSERT INTO label(name,color) VALUES (?,?)")
	_, err = stmt.Exec(label.Name, label.Color)
	if err != nil {
		log.Error("SaveLabel Exec Error:", err)
	}

	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Error("SaveLabel Close error:", err)
		}
	}(stmt)

	err = tx.Commit()
	if err != nil {
		log.Error("SaveLabel Error:", err)
	}
}

func createLabelTable() {
	var _, err = db.Exec("create table if not exists label" +
		"(id integer not null constraint label_pk primary key autoincrement, name TEXT default '自定义标签' not null, color TEXT not null );")
	if err != nil {
		log.Error("createLabelTable Error:", err)
	}
}

func GetLabelList() ([]define.Label, error) {
	rows, err := db.Query("select * from label")

	if err != nil {
		log.Error("GetLabelList Query Error:", err)
	}
	var labelList []define.Label
	for rows.Next() {
		var label define.Label
		err := rows.Scan(&label.Id, &label.Name, &label.Color)
		if err != nil {
			log.Error("GetLabelList Scan Error:", err)
			return nil, err
		}
		labelList = append(labelList, label)
	}
	return labelList, err
}
