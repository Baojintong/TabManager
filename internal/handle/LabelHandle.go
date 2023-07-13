package handle

import (
	"github.com/labstack/gommon/log"
	"tabManager/internal/define"
)

func SaveLabel(label define.Label) {
	//createLabelTable()
	saveLabel(label)
}

func saveLabel(label define.Label) {
	db.Exec(define.INSERT_LABEL, label.Name, label.Color)
}

func CreateLabelTable() {
	db.ExecNoTran(define.CREATE_LABEL_TABLE)
}

func GetLabelList() []define.Label {
	rows := db.Query(define.SELECT_LABEL)
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
