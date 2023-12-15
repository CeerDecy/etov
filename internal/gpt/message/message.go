package message

import "github.com/sashabaranov/go-openai"

type Messages []openai.ChatCompletionMessage

func NewMessages() *Messages {
	return &Messages{}
}

func (m *Messages) AddChatMessageRoleUserMsg(content string) {
	*m = append(*m, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: content,
	})
}

func (m *Messages) AddChatMessageGPTMsg(content string) {
	*m = append(*m, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleSystem,
		Content: content,
	})
}

func (m *Messages) GetMessages() []openai.ChatCompletionMessage {
	return *m
}
