package define

const (
	DATA_SOURCE_NAME = "file:" + DB_ALL + "?mode=rwc"

	DB_ALL = DB_PATH + DB_FILE

	DB_FILE    = "tabs.db"
	DB_PATH    = "./db/"
	LISTEN_URL = "localhost:12315"

	DATE_FORMAT = "2006-01-02 15:04:05"
)
