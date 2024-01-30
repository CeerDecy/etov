package engine

import (
	"context"

	"github.com/sashabaranov/go-openai"
	"github.com/sirupsen/logrus"

	"etov/internal/gpt/message"
	"etov/internal/gpt/session"
	"etov/internal/gpt/types"
)

type ChatGPT struct {
	client *openai.Client
	model  string
}

func newChatGPT(config types.ChatEngineConfig) *ChatGPT {
	clientConfig := openai.DefaultConfig(config.AuthToken)
	if config.BaseUrl != "" {
		clientConfig.BaseURL = config.BaseUrl
	}
	client := openai.NewClientWithConfig(clientConfig)
	return &ChatGPT{
		client: client,
		model:  config.Model,
	}
}

func (c *ChatGPT) Push(msg *message.Messages) (*session.Session, error) {
	stream, err := c.client.CreateChatCompletionStream(context.Background(),
		openai.ChatCompletionRequest{
			Model:    string(c.model),
			Messages: msg.GetMessages(),
			Stream:   true,
		})
	if err != nil {
		logrus.Error("Failed to create chat completion stream:", err)
		return nil, err
	}
	sess := session.NewSession(stream)
	go sess.ReadStream()
	return sess, nil
}
