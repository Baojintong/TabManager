package handle

import (
	"github.com/labstack/gommon/log"
	"tabManager/internal/define"
)

func CreateConfigTable() {
	db.ExecNoTran(define.CREATE_CONFIG_TABLE)
}


func GetConfigList() []define.Config{
	rows := db.Query(define.SELECT_CONFIG)
	var configList []define.Config
	for rows.Next() {
		var config define.Config
		err := rows.Scan(&config.Id, &config.Key, &config.Value)
		if err != nil {
			log.Error("GetLabelList Scan Error:", err)
			return nil
		}
		configList = append(configList, config)
	}
	return configList
}