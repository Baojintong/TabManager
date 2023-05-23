package define

type TabsData struct {
	Id       int32  `json:"id"`
	Title    string `json:"title"`
	Url      string `json:"url"`
	IconUrl  string `json:"iconUrl"`
	Describe string `json:"describe"`
	SaveTime string `json:"saveTime"`
	Status   int8   `json:"status"`
}
