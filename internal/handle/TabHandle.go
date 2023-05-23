package handle

import (
	"C"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"io"
	"net/http"
	"tabManager/internal/define"
	"time"
)

var db, _ = sql.Open("sqlite3", "file:./db/tabs.db?mode=rwc")

func TabHandler(w http.ResponseWriter, r *http.Request) {

	var body, err = io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(" tabsHandler Error:", err.Error())
	}
	var tabsData []define.TabsData
	var jsonStr = string(body)
	err = json.Unmarshal([]byte(jsonStr), &tabsData)
	if err != nil {
		fmt.Println(" tabsHandler Error:", err.Error())
		return
	}
	now := time.Now()
	nowDate := now.Format("2006-01-02")
	for i, n := 0, len(tabsData); i < n; i++ {
		tabsData[i].Describe = tabsData[i].Title
		tabsData[i].SaveTime = nowDate
	}
	saveTabsData(nowDate, tabsData)

	_, err = w.Write([]byte("success"))
	if err != nil {
		fmt.Println(" tabsHandler resp Error:", err.Error())
	}
}

func saveTabsData(fold string, tabsDatas []define.TabsData) {
	createTable()
	batchInsert(tabsDatas)
}

func createTable() {
	var _, err = db.Exec("CREATE TABLE IF NOT EXISTS tabs " +
		"(id integer not null constraint tabs_pk primary key autoincrement,title TEXT,icon_url TEXT,url TEXT,describe TEXT,save_time TEXT not null,status integer default 0 not null)")
	if err != nil {
		fmt.Println("createTable Error:", err)
	}
}

func batchInsert(tabsDatas []define.TabsData) {
	tx, err := db.Begin()
	if err != nil {
		fmt.Println("batchInsert Error:", err)
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("INSERT INTO tabs(title,icon_url,url,describe,save_time) VALUES (?,?,?,?,?)")
	if err != nil {
		fmt.Println("batchInsert Error:", err)
	}
	defer stmt.Close()

	for _, d := range tabsDatas {
		_, err = stmt.Exec(d.Title, d.IconUrl, d.Url, d.Describe, d.SaveTime)
		if err != nil {
			fmt.Println("batchInsert Error:", err)
		}
	}

	err = tx.Commit()
	if err != nil {
		fmt.Println("batchInsert Error:", err)
	}
}

func QueryAllTabs() ([]define.TabsData, error) {
	rows, err := db.Query("SELECT * FROM tabs")

	if err != nil {
		fmt.Println("queryAllTabs Error:", err)
	}
	tabsData := []define.TabsData{}
	for rows.Next() {
		var tabs define.TabsData
		err := rows.Scan(&tabs.Id, &tabs.Title, &tabs.IconUrl, &tabs.Url, &tabs.Describe, &tabs.SaveTime, &tabs.Status)
		if err != nil {
			fmt.Println("queryAllTabs Error:", err)
			return nil, err
		}
		tabsData = append(tabsData, tabs)
	}
	return tabsData, err
}
