package handle

import (
	"C"
	"encoding/json"
	"github.com/labstack/gommon/log"
	_ "github.com/mattn/go-sqlite3"
	"io"
	"net/http"
	"tabManager/internal/define"
	"time"
)

func TabHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("TabHandler.........")
	var body, err = io.ReadAll(r.Body)
	if err != nil {
		log.Error(" tabsHandler Error:", err)
	}
	var Tab []define.Tab
	var jsonStr = string(body)
	err = json.Unmarshal([]byte(jsonStr), &Tab)
	if err != nil {
		log.Error(" tabsHandler Error:", err)
		return
	}
	now := time.Now()
	nowDate := now.Format("2006-01-02")
	timestamp := time.Now().Unix()
	var interfaces []interface{}
	for i, n := 0, len(Tab); i < n; i++ {
		Tab[i].Describe = Tab[i].Title
		Tab[i].SaveTime = nowDate
		Tab[i].TimeStamp = timestamp
		interfaces = append(interfaces, Tab[i])
	}
	saveTab(interfaces)

	_, err = w.Write([]byte("success"))
	if err != nil {
		log.Error(" tabsHandler resp Error:", err)
	}
}

func saveTab(datas []interface{}) {
	createTabTable()
	batchInsert(datas)
}

func createTabTable() {
	db.Exec("create table if not exists tabs " +
		"(id integer not null constraint tabs_pk primary key autoincrement,title TEXT,icon_url TEXT,url TEXT,describe TEXT,save_time TEXT not null,status integer default 0 not null)")
}

func GetTabList() []define.Tab {
	rows := db.Query("SELECT * FROM tabs order by time_stamp desc")
	var tabList []define.Tab
	for rows.Next() {
		var tab define.Tab
		err := rows.Scan(&tab.Id, &tab.Title, &tab.IconUrl, &tab.Url, &tab.Describe, &tab.SaveTime, &tab.Status, &tab.TimeStamp)
		if err != nil {
			log.Error("queryAllTabs Scan Error:", err)
			return nil
		}
		tabList = append(tabList, tab)
	}
	return tabList
}

func batchInsert(datas []interface{}) {
	db.BatchExec("INSERT INTO tabs(title,icon_url,url,describe,save_time,time_stamp) VALUES (:title,:iconUrl,:url,:describe,:saveTime,:timeStamp)", datas)
}

func UpdateTab(tab define.Tab) {
	db.Exec("UPDATE tabs SET title=?,`describe`=? WHERE id=?", tab.Title, tab.Describe, tab.Id)
}

func DeleteTab(tab define.Tab) {
	db.Exec("DELETE FROM tabs WHERE id=?", tab.Id)
}
