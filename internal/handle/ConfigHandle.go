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
		err := rows.Scan(&config.Id, &config.Key, &config.Value, &config.Describe)
		if err != nil {
			log.Error("GetLabelList Scan Error:", err)
			return nil
		}
		configList = append(configList, config)
	}
	return configList
}

func SaveConfig(configs []define.Config){
	var interfaces []interface{}
	for i, n := 0, len(configs); i < n; i++ {
		configObj := configs[i]
		config:= define.Config{}
		config.Id = configObj.Id
		config.Key = configObj.Key
		config.Value = "111111111"//configObj.Value
		config.Describe = configObj.Describe
		interfaces = append(interfaces, config)
	}
	db.BatchExec(define.UPDATE_CONFIG, interfaces)
}