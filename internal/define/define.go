package define

type Tab struct {
	Id        uint32 `json:"id"`
	Title     string `json:"title" db:"title"`
	Url       string `json:"url" db:"url"`
	IconUrl   string `json:"iconUrl" db:"iconUrl"`
	Describe  string `json:"describe" db:"describe"`
	SaveTime  string `json:"saveTime" db:"saveTime"`
	Status    uint8  `json:"status"`
	TimeStamp int64  `json:"timeStamp" db:"timeStamp"`
}

type Label struct {
	Id    uint32 `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}
