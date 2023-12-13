package request

type ChatRequest struct {
	Uid     string `json:"uid"`
	Content string `json:"content"`
}
