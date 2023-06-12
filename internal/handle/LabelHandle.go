package handle

import (
	"github.com/labstack/gommon/log"
	"tabManager/internal/define"
)

var db DbHandle = new(DbHandleImpl)

func SaveLabel(label define.Label) {
	createLabelTable()
	saveLabel(label)
}

func saveLabel(label define.Label) {
	db.Connect()
	db.Exec("INSERT INTO label(name,color) VALUES (?,?)", label.Name, label.Color)
	db.Close()
}

func createLabelTable() {
	db.Connect()
	db.Exec("create table if not exists label" +
		"(id integer not null constraint label_pk primary key autoincrement, name TEXT default '自定义标签' not null, color TEXT not null );")
	db.Close()
}

func GetLabelList() []define.Label {
	db.Connect()
	rows := db.Query("select * from label")
	var labelList []define.Label
	for rows.Next() {
		var label define.Label
		err := rows.Scan(&label.Id, &label.Name, &label.Color)
		if err != nil {
			log.Error("GetLabelList Scan Error:", err)
			return nil
		}
		labelList = append(labelList, label)
	}
	db.Close()
	return labelList
}
