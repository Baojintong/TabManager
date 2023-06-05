package handle

import (
	"github.com/labstack/gommon/log"
	"tabManager/internal/define"
)

func SaveLabel(label define.Label) {
	tx, err := db.Begin()
	if err != nil {
		log.Error("SaveLabel Error:", err)
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("INSERT INTO label(name,color) VALUES (?,?)")
	if err != nil {
		log.Error("SaveLabel Error:", err)
	}
	defer stmt.Close()
	stmt.Exec(label.Name, label.Color)

	err = tx.Commit()
	if err != nil {
		log.Error("SaveLabel Error:", err)
	}
}
