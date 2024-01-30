package types

type ChatEngineConfig struct {
	AuthToken string `json:"authToken"`
	BaseUrl   string `json:"baseUrl"`
	Model     string `json:"model"`
}
