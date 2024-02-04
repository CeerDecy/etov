package request

type ReduceDuplication struct {
	Mode     string `json:"mode"`
	EngineId string `json:"engineId"`
	Content  string `json:"content"`
}
