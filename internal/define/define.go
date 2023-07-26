package define

type Tab struct {
	Id        uint32   `json:"id" auto:"id"`
	Title     string   `json:"title" input:"title"`
	Url       string   `json:"url" input:"url"`
	IconUrl   string   `json:"iconUrl" input:"iconUrl"`
	Describe  string   `json:"describe" input:"describe"`
	SaveTime  string   `json:"saveTime" input:"saveTime"`
	Status    uint8    `json:"status" auto:"status"`
	TimeStamp int64    `json:"timeStamp" input:"timeStamp"`
	LabelIds  []uint32 `json:"labelIds"`
}

type Label struct {
	Id    uint32 `json:"id" auto:"id"`
	Name  string `json:"name" input:"title"`
	Color string `json:"color" input:"color"`
}

type TagLabel struct {
	Id      uint32 `json:"id" auto:"id"`
	TabId   uint32 `json:"tabId" input:"tabId"`
	LabelId uint32 `json:"labelId" input:"labelId"`
}

type Task struct {
	Id         uint32 `json:"id" auto:"id"`
	Name       string `json:"name" input:"name"`
	CreateTime string `json:"createTime" input:"createTime"`
	Status     uint8  `json:"status" auto:"status"`
	TimeStamp  int64  `json:"timeStamp" input:"timeStamp"`
	TargetId   uint32 `json:"targetId" input:"targetId"`
	TargetType string `json:"targetType" input:"targetType"`
}

type Config struct {
	Id       uint32 `json:"id" auto:"id"`
	Key      string `json:"key" input:"key"`
	Value    string `json:"value" input:"value"`
	Describe string `json:"describe" input:"describe"`
}
