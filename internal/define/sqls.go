package define

const (
	//tab
	CREATE_TAB_TABLE      = "create table if not exists tab (id integer not null constraint tabs_pk primary key autoincrement,title TEXT,icon_url TEXT,url TEXT,describe TEXT,save_time TEXT not null,status integer default 0 not null,time_stamp integer default 0 not null)"
	SELECT_TAB_LIST       = "SELECT * FROM tab order by time_stamp desc"
	SELECT_TAB            = "SELECT * FROM tab where id=? order by time_stamp desc"
	SELECT_TAB_JOIN_LABEL = "select tab.* from tab left join tab_label tl on tab.id = tl.tab_id where tl.label_id=? order by tab.time_stamp desc"
	INSERT_TAB            = "INSERT INTO tab(title,icon_url,url,describe,save_time,time_stamp) VALUES (:title,:iconUrl,:url,:describe,:saveTime,:timeStamp)"
	UPDATE_TAB            = "UPDATE tab SET title=?,`describe`=? WHERE id=?"
	DELETE_TAB            = "DELETE FROM tab WHERE id=?"

	//tab_label
	CREATE_TAB_LABEL_TABLE = "create table if not exists tab_label (id integer not null constraint tab_label_pk primary key autoincrement,tab_id integer not null, label_id integer not null)"
	INSERT_TAB_LABEL       = "INSERT INTO tab_label(tab_id, label_id) VALUES (:tabId,:labelId)"
	DELETE_TAB_LABEL       = "DELETE FROM tab_label WHERE tab_id=?"
	SELECT_TAB_LABEL       = "SELECT * FROM tab_label WHERE tab_id=?"
	SELECT_LABEL_TAB       = "SELECT * FROM tab_label WHERE label_id=?"

	//label
	CREATE_LABEL_TABLE = "create table if not exists label (id integer not null constraint label_pk primary key autoincrement, name TEXT default '自定义标签' not null, color TEXT not null )"
	INSERT_LABEL       = "INSERT INTO label(name,color) VALUES (?,?)"
	SELECT_LABEL       = "SELECT * FROM label"

	//task
	CREATE_TASK_TABLE = "create table if not exists task(id integer not null constraint task_pk primary key autoincrement, name TEXT, create_time TEXT not null, status integer default 0 not null, time_stamp integer default 0 not null ,target_id INTEGER DEFAULT (0), target_type TEXT)"
	INSERT_TASK       = "INSERT INTO task (name, create_time, time_stamp, target_id,target_type) VALUES(:name,:createTime,:timeStamp,:targetId,:targetType)"

	//config
	CREATE_CONFIG_TABLE = "create table if not exists config(id integer not null constraint config_pk primary key autoincrement, key TEXT default '' not null,value TEXT default '' not null,describe TEXT)"

	SELECT_CONFIG = "SELECT * FROM config"

	UPDATE_CONFIG = "UPDATE config SET value=:value,describe=:describe,key=:key WHERE id=:id"

	INSERT_CONFIG = "INSERT INTO config(key,value,describe) VALUES (:key,:value,:describe)"
)
