package turingapi

type Body struct {
	ReqType    int      `json:"reqType"`
	Perception *Query    `json:"perception"`
	UserInfo   *UserInfo `json:"userInfo"`
}

type Query struct {
	// 文本信息
	InputText struct{ Text string `json:"text"` } `json:"inputText"`
	// 图片信息
	InputImage struct{ Url string `json:"url"` } `json:"inputImage"`
	// 音频信息
	InputMedia struct{ Url string `json:"url"` } `json:"inputMedia"`
	SelfInfo struct {
		Location *Location `json:"location"`
	} `json:"selfInfo"`
}

type Location struct {
	City     string `json:"city"`
	Province string `json:"province"`
	Street   string `json:"street"`
}

type UserInfo struct {
	ApiKey string `json:"apiKey"`
	UserId string `json:"userId"`
}
