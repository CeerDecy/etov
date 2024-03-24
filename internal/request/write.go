package request

type WriteRequest struct {
	Content  string `json:"content"`
	Types    string `json:"types"`
	EngineId string `json:"engineId"`
}
