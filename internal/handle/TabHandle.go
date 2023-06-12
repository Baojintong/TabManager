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
	for i, n := 0, len(Tab); i < n; i++ {
		Tab[i].Describe = Tab[i].Title
		Tab[i].SaveTime = nowDate
	}
	saveTab(Tab)

	_, err = w.Write([]byte("success"))
	if err != nil {
		log.Error(" tabsHandler resp Error:", err)
	}
}

func saveTab(Tabs []define.Tab) {
	createTabTable()
	batchInsert(Tabs)
}

func createTabTable() {
	db.Connect()
	db.Exec("create table if not exists tabs " +
		"(id integer not null constraint tabs_pk primary key autoincrement,title TEXT,icon_url TEXT,url TEXT,describe TEXT,save_time TEXT not null,status integer default 0 not null)")
	db.Close()
}

func GetTabList() []define.Tab {
	db.Connect()
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
	db.Close()
	return tabList
}

func batchInsert(Tabs []define.Tab) {
	db.Connect()
	var interfaces []interface{}
	timestamp := time.Now().Unix()
	for _, tab := range Tabs {
		tab.TimeStamp = timestamp
		interfaces = append(interfaces, tab)
	}
	db.BatchExec("INSERT INTO tabs(title,icon_url,url,describe,save_time,time_stamp) VALUES (:title,:iconUrl,:url,:describe,:saveTime,:timeStamp)", interfaces)
	db.Close()
}

func UpdateTab(tab define.Tab) {
	db.Connect()
	db.Exec("UPDATE tabs SET title=?,`describe`=? WHERE id=?", tab.Title, tab.Describe, tab.Id)
	db.Close()
}

func DeleteTab(tab define.Tab) {
	db.Connect()
	db.Exec("DELETE FROM tabs WHERE id=?", tab.Id)
	db.Close()
}
