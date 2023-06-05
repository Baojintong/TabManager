package handle

import (
	"database/sql"
	"github.com/labstack/gommon/log"
	"tabManager/internal/define"
)

func SaveLabel(label define.Label) {
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
