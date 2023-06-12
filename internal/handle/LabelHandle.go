package handle

import (
	"github.com/labstack/gommon/log"
	"tabManager/internal/define"
)


func SaveLabel(label define.Label) {
	createLabelTable()
	saveLabel(label)
}

func saveLabel(label define.Label) {
	db.Exec("INSERT INTO label(name,color) VALUES (?,?)", label.Name, label.Color)
}

func createLabelTable() {
	db.Exec("create table if not exists label" +
		"(id integer not null constraint label_pk primary key autoincrement, name TEXT default '自定义标签' not null, color TEXT not null );")
}

func GetLabelList() []define.Label {
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
	return labelList
}
