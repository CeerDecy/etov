package request

type TranslatorRequest struct {
	TargetLang string `json:"target_lang"`
	Content    string `json:"content"`
	EngineId   string `json:"engineId"`
}
