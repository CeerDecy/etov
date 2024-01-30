package engine

import (
	"etov/internal/gpt/message"
	"etov/internal/gpt/session"
	"etov/internal/gpt/types"
)

type ChatEngine interface {
	Push(msg *message.Messages) (*session.Session, error)
}

func NewChatEngine(model, authToken, apiUrl string) ChatEngine {
	config := types.ChatEngineConfig{
		Model:     model,
		AuthToken: authToken,
		BaseUrl:   apiUrl,
	}
	return newChatGPT(config)
}
