package request

type ChatRequest struct {
	ChatId   string `json:"chatId"`
	Content  string `json:"content"`
	EngineId string `json:"engineId"`
}
