package handle

import (
	"C"
	"database/sql"
	"encoding/json"
	_ "github.com/mattn/go-sqlite3"
	"io"
	"net/http"
	"time"
)

type TabsData struct {
	Title    string `json:"title"`
	Url      string `json:"url"`
	IconUrl  string `json:"iconUrl"`
	Describe string `json:"describe"`
}

func TabHandler(w http.ResponseWriter, r *http.Request) {

	var body, err = io.ReadAll(r.Body)
	if err != nil {
		println(" tabsHandler Error:", err.Error())
	}
	var tabsData []TabsData
	var jsonStr = string(body)
	err = json.Unmarshal([]byte(jsonStr), &tabsData)
	if err != nil {
		println(" tabsHandler Error:", err.Error())
		return
	}
	now := time.Now()
	nowDate := now.Format("2006-01-02")
	for i, n := 0, len(tabsData); i < n; i++ {
		tabsData[i].Describe = tabsData[i].Title
	}
	saveTabsData(nowDate, tabsData)

	_, err = w.Write([]byte("success"))
	if err != nil {
		println(" tabsHandler resp Error:", err.Error())
	}
}

func saveTabsData(fold string, tabsDatas []TabsData) {
	db, err := sql.Open("sqlite3", "file:./db/tabs.db?mode=rwc")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("CREATE TABLE log(id INT, content VARCHAR(1024))")
	if err != nil {
		panic(err)
	}
}
