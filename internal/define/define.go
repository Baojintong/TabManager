package define

type TabsData struct {
	Id        uint32 `json:"id"`
	Title     string `json:"title"`
	Url       string `json:"url"`
	IconUrl   string `json:"iconUrl"`
	Describe  string `json:"describe"`
	SaveTime  string `json:"saveTime"`
	Status    uint8  `json:"status"`
	TimeStamp int64  `json:"timeStamp"`
}

type Label struct {
	Id    uint32 `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
}
