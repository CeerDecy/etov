package etovs

import (
	"fmt"

	"etov/internal/gpt/engine"
	"etov/internal/gpt/message"
	"etov/internal/gpt/session"
)

type Summary struct {
	msg *message.Messages
}

func (s *Summary) AppendContext(msg message.Messages) Interface {
	s.msg.Append(msg)
	return s
}

func (s *Summary) Execute(engine engine.ChatEngine, params ...any) (*session.Session, error) {
	if len(params) > 0 {
		s.msg.AddChatMessageRoleUserMsg(fmt.Sprintf("%s", params...))
	}
	return engine.Push(s.msg)
}

func NewSummary() Interface {
	msg := message.NewMessages()
	msg.AddChatMessageRoleUserMsg("你是一个总结的AI助手。请总结以下内容，并使用中文回答。")
	return &Summary{
		msg: msg,
	}
}
