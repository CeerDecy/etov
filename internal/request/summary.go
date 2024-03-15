package request

type SummaryRequest struct {
	Content  string `json:"content"`
	FilePath string `json:"filepath"`
	EngineId string `json:"engineId"`
}
