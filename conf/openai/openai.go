package openai

import (
	_ "embed"
)

type OpenAI struct {
	AuthToken string `json:"authToken"`
	BaseUrl   string `json:"baseUrl"`
}
