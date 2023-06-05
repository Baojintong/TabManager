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
	var tabsData []define.TabsData
	var jsonStr = string(body)
	err = json.Unmarshal([]byte(jsonStr), &tabsData)
	if err != nil {
		log.Error(" tabsHandler Error:", err)
		return
	}
	now := time.Now()
	nowDate := now.Format("2006-01-02")
	for i, n := 0, len(tabsData); i < n; i++ {
		tabsData[i].Describe = tabsData[i].Title
		tabsData[i].SaveTime = nowDate
	}
	saveTab(tabsData)

	_, err = w.Write([]byte("success"))
	if err != nil {
		log.Error(" tabsHandler resp Error:", err)
	}
}

func saveTab(tabsDatas []define.TabsData) {
	createTable()
	batchInsert(tabsDatas)
}

func createTable() {
	var _, err = db.Exec("CREATE TABLE IF NOT EXISTS tabs " +
		"(id integer not null constraint tabs_pk primary key autoincrement,title TEXT,icon_url TEXT,url TEXT,describe TEXT,save_time TEXT not null,status integer default 0 not null)")
	if err != nil {
		log.Error("createTable Error:", err)
	}
}

func batchInsert(tabsDatas []define.TabsData) {
	tx, err := db.Begin()
	if err != nil {
		log.Error("batchInsert Error:", err)
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("INSERT INTO tabs(title,icon_url,url,describe,save_time,time_stamp) VALUES (?,?,?,?,?,?)")
	if err != nil {
		log.Error("batchInsert Error:", err)
	}
	defer stmt.Close()
	timestamp := time.Now().Unix()
	for _, d := range tabsDatas {
		_, err = stmt.Exec(d.Title, d.IconUrl, d.Url, d.Describe, d.SaveTime, timestamp)
		if err != nil {
			log.Error("batchInsert Error:", err)
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Error("batchInsert Error:", err)
	}
}

func QueryAllTabs() ([]define.TabsData, error) {
	rows, err := db.Query("SELECT * FROM tabs order by time_stamp desc")

	if err != nil {
		log.Error("queryAllTabs Error:", err)
	}
	tabsData := []define.TabsData{}
	for rows.Next() {
		var tabs define.TabsData
		err := rows.Scan(&tabs.Id, &tabs.Title, &tabs.IconUrl, &tabs.Url, &tabs.Describe, &tabs.SaveTime, &tabs.Status, &tabs.TimeStamp)
		if err != nil {
			log.Error("queryAllTabs Error:", err)
			return nil, err
		}
		tabsData = append(tabsData, tabs)
	}
	return tabsData, err
}

func UpdateTab(tab define.TabsData) {
	tx, err := db.Begin()
	stmt, err := db.Prepare("UPDATE tabs SET title=?,`describe`=? WHERE id=?")
	defer stmt.Close()

	result, err := stmt.Exec(tab.Title, tab.Describe, tab.Id)
	_, err = result.RowsAffected()
	if err != nil {
		log.Error("UpdateTab error:", err)
	}

	err = tx.Commit()
	if err != nil {
		log.Error("UpdateTab Error:", err)
	}
}

func DeleteTab(tab define.TabsData) {
	stmt, err := db.Prepare("DELETE FROM tabs WHERE id=?")
	defer stmt.Close()

	result, err := stmt.Exec(tab.Id)
	_, err = result.RowsAffected()
	if err != nil {
		log.Error("UpdateTab error:", err)
	}
}
