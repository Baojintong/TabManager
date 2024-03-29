package define

type Tab struct {
	Id        uint32   `json:"id" db:"id"`
	Title     string   `json:"title" db:"title"`
	Url       string   `json:"url" db:"url"`
	IconUrl   string   `json:"iconUrl" db:"iconUrl"`
	Describe  string   `json:"describe" db:"describe"`
	SaveTime  string   `json:"saveTime" db:"saveTime"`
	Status    uint8    `json:"status" db:"status"`
	TimeStamp int64    `json:"timeStamp" db:"timeStamp"`
	LabelIds  []uint32 `json:"labelIds"`
}

type Label struct {
	Id    uint32 `json:"id" db:"id"`
	Name  string `json:"name" db:"title"`
	Color string `json:"color" db:"color"`
}

type TagLabel struct {
	Id      uint32 `json:"id" db:"id"`
	TabId   uint32 `json:"tabId" db:"tabId"`
	LabelId uint32 `json:"labelId" db:"labelId"`
}

type Task struct {
	Id         uint32 `json:"id" db:"id"`
	Name       string `json:"name" db:"name"`
	CreateTime string `json:"createTime" db:"createTime"`
	Status     uint8  `json:"status" db:"status"`
	TimeStamp  int64  `json:"timeStamp" db:"timeStamp"`
	TargetId   uint32 `json:"targetId" db:"targetId"`
	TargetType string `json:"targetType" db:"targetType"`
}

type Config struct {
	Id       uint32 `json:"id" db:"id"`
	Key      string `json:"key" db:"key"`
	Value    string `json:"value" db:"value"`
	Describe string `json:"describe" db:"describe"`
}
