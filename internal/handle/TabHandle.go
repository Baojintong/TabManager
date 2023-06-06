package handle

import (
	"C"
	"database/sql"
	"encoding/json"
	"github.com/labstack/gommon/log"
	_ "github.com/mattn/go-sqlite3"
	"io"
	"net/http"
	"tabManager/internal/define"
	"time"
)

var dbPath = "./db/tabs.db"
var db, _ = sql.Open("sqlite3", define.DATA_SOURCE_NAME)

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
	var _, err = db.Exec("create table if not exists tabs " +
		"(id integer not null constraint tabs_pk primary key autoincrement,title TEXT,icon_url TEXT,url TEXT,describe TEXT,save_time TEXT not null,status integer default 0 not null)")
	if err != nil {
		log.Error("createTabTable Error:", err)
	}
}

func GetTabList() ([]define.Tab, error) {
	rows, err := db.Query("SELECT * FROM tabs order by time_stamp desc")

	if err != nil {
		log.Error("queryAllTabs Query Error:", err)
	}
	var tabList []define.Tab
	for rows.Next() {
		var tab define.Tab
		err := rows.Scan(&tab.Id, &tab.Title, &tab.IconUrl, &tab.Url, &tab.Describe, &tab.SaveTime, &tab.Status, &tab.TimeStamp)
		if err != nil {
			log.Error("queryAllTabs Scan Error:", err)
			return nil, err
		}
		tabList = append(tabList, tab)
	}
	return tabList, err
}

func batchInsert(Tabs []define.Tab) {
	tx, err := db.Begin()
	if err != nil {
		log.Error("batchInsert Begin Error:", err)
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("INSERT INTO tabs(title,icon_url,url,describe,save_time,time_stamp) VALUES (?,?,?,?,?,?)")
	if err != nil {
		log.Error("batchInsert Prepare Error:", err)
	}

	timestamp := time.Now().Unix()
	for _, d := range Tabs {
		_, err = stmt.Exec(d.Title, d.IconUrl, d.Url, d.Describe, d.SaveTime, timestamp)
		if err != nil {
			log.Error("batchInsert Exec Error:", err)
		}
	}

	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Error("batchInsert Close Error:", err)
		}
	}(stmt)

	err = tx.Commit()
	if err != nil {
		log.Error("batchInsert Commit Error:", err)
	}

}

func UpdateTab(tab define.Tab) {
	tx, err := db.Begin()
	if err != nil {
		log.Error("UpdateTab Begin Error:", err)
	}
	defer tx.Rollback()

	stmt, err := db.Prepare("UPDATE tabs SET title=?,`describe`=? WHERE id=?")
	_, err = stmt.Exec(tab.Title, tab.Describe, tab.Id)
	if err != nil {
		log.Error("UpdateTab Exec error:", err)
	}

	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Error("UpdateTab Close error:", err)
		}
	}(stmt)

	err = tx.Commit()
	if err != nil {
		log.Error("UpdateTab Commit Error:", err)
	}
}

func DeleteTab(tab define.Tab) {
	tx, err := db.Begin()
	if err != nil {
		log.Error("DeleteTab Begin Error:", err)
	}
	defer tx.Rollback()

	stmt, err := db.Prepare("DELETE FROM tabs WHERE id=?")
	_, err = stmt.Exec(tab.Id)
	if err != nil {
		log.Error("DeleteTab Exec error:", err)
	}

	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Error("DeleteTab Close error:", err)
		}
	}(stmt)

	err = tx.Commit()
	if err != nil {
		log.Error("DeleteTab Commit Error:", err)
	}
}
