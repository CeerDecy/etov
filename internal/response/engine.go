package response

type Engine struct {
	Name string `json:"name"`
	ID   int64  `json:"id"`
}

type GetSupportEngineResponse struct {
	Platform []Engine `json:"platform"`
	Custom   []Engine `json:"custom"`
}
