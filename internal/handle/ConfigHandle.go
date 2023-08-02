package handle

import (
	"database/sql"
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

func InitConfig(){
	var config define.Config
	config.Key="path"
	config.Value="file"
	config.Describe="文件路径"
	path :=GetConfigByKey("path")
	if path.Id == 0 {
		SaveConfig([]define.Config{config})
	}
}

func SaveConfig(configs []define.Config){
	var interfaces []interface{}
	for i, n := 0, len(configs); i < n; i++ {
		configObj := configs[i]
		config:= define.Config{}
		config.Key = configObj.Key
		config.Value = configObj.Value
		config.Describe = configObj.Describe
		interfaces = append(interfaces, config)
	}
	db.BatchExec(define.INSERT_CONFIG, interfaces)
}

func UpdateConfig(configs []define.Config){
	var interfaces []interface{}
	for i, n := 0, len(configs); i < n; i++ {
		configObj := configs[i]
		config:= define.Config{}
		config.Id = configObj.Id
		config.Key = configObj.Key
		config.Value = configObj.Value
		config.Describe = configObj.Describe
		interfaces = append(interfaces, config)
	}
	db.BatchExec(define.UPDATE_CONFIG, interfaces)
}

func GetConfigByKey(key string) define.Config{
	row := db.QueryRow(define.SELECT_CONFIG_BY_KEY, key)
	var config define.Config
	err := row.Scan(&config.Id, &config.Key, &config.Value, &config.Describe)
	if err == sql.ErrNoRows {
		return config
	}
	if err != nil {
		panic(err)
	}
	return config
}

func GetConfigValueByKey(key string) string{
	row := db.QueryRow(define.SELECT_CONFIG_BY_KEY, key)
	var config define.Config
	err := row.Scan(&config.Id, &config.Key, &config.Value, &config.Describe)
	if err == sql.ErrNoRows {
		return ""
	}
	if err != nil {
		panic(err)
	}
	return config.Value
}