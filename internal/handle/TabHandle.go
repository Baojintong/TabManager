package handle

import (
	"C"
	"database/sql"
	"encoding/json"
	"github.com/labstack/gommon/log"
	_ "github.com/mattn/go-sqlite3"
	"github.com/wailsapp/wails/v2/pkg/runtime"
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
	runtime.EventsEmit(*context_, "flushTabs")
}

func saveTab(datas []interface{}) {
	createTabTable()
	batchInsert(datas)
}

func createTabTable() {
	db.Exec(define.CREATE_TAB_TABLE)
	db.Exec(define.CREATE_TAB_LABEL_TABLE)
}

func GetTabList(labelId uint32) []define.Tab {
	log.Info("labelId:", labelId)
	var rows *sql.Rows = nil
	if labelId == 0 {
		rows = db.Query(define.SELECT_TAB_LIST)
	} else {
		rows = db.Query(define.SELECT_TAB_JOIN_LABEL, labelId)
	}
	var tabList []define.Tab
	if rows != nil {
		for rows.Next() {
			var tab define.Tab
			err := rows.Scan(&tab.Id, &tab.Title, &tab.IconUrl, &tab.Url, &tab.Describe, &tab.SaveTime, &tab.Status, &tab.TimeStamp)
			if err != nil {
				log.Error("queryAllTabs Scan Error:", err)
				return nil
			}
			tabList = append(tabList, tab)
		}
	}
	return tabList
}

func GetTab(tabId uint32) define.Tab {
	row := db.QueryRow(define.SELECT_TAB, tabId)
	var tab define.Tab
	err := row.Scan(&tab.Id, &tab.Title, &tab.IconUrl, &tab.Url, &tab.Describe, &tab.SaveTime, &tab.Status, &tab.TimeStamp)
	if err != nil {
		panic(err)
	}
	return tab
}

func batchInsert(datas []interface{}) {
	db.BatchExec(define.INSERT_TAB, datas)
}

func batchInsertTabLabel(datas []interface{}) {
	db.BatchExec(define.INSERT_TAB_LABEL, datas)
}

func UpdateTab(tab define.Tab) {
	db.Exec(define.UPDATE_TAB, tab.Title, tab.Describe, tab.Id)
	tabId := tab.Id
	labelIds := tab.LabelIds
	var interfaces []interface{}
	for i, n := 0, len(labelIds); i < n; i++ {
		labelId := labelIds[i]
		tagLabel := define.TagLabel{}
		tagLabel.TabId = tabId
		tagLabel.LabelId = labelId
		interfaces = append(interfaces, tagLabel)
	}
	cleanTabLabel(tabId)
	batchInsertTabLabel(interfaces)
}

func DeleteTab(tab define.Tab) {
	db.Exec(define.DELETE_TAB, tab.Id)
}

func cleanTabLabel(tabId uint32) {
	db.Exec(define.DELETE_TAB_LABEL, tabId)
}

func QueryTabLabel(tabId uint32) []define.TagLabel {
	return getTabLabelData(tabId, define.SELECT_TAB_LABEL)
}

func QueryLabelTab(labelId uint32) []define.TagLabel {
	return getTabLabelData(labelId, define.SELECT_LABEL_TAB)
}

func getTabLabelData(id uint32, sql string) []define.TagLabel {
	rows := db.Query(sql, id)
	var tagLabelList []define.TagLabel
	for rows.Next() {
		var tagLabel define.TagLabel
		err := rows.Scan(&tagLabel.Id, &tagLabel.TabId, &tagLabel.LabelId)
		if err != nil {
			log.Error("queryTabLabel Scan Error:", err)
			return nil
		}
		tagLabelList = append(tagLabelList, tagLabel)
	}
	return tagLabelList
}
