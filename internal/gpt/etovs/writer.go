package etovs

import (
	"fmt"

	"etov/internal/gpt/engine"
	"etov/internal/gpt/message"
	"etov/internal/gpt/session"
)

type Writer struct {
	msg *message.Messages
}

func NewWriter() *Writer {
	msg := message.NewMessages()
	return &Writer{msg: msg}
}

func (w *Writer) Execute(engine engine.ChatEngine, params ...any) (*session.Session, error) {
	w.msg.AddChatMessageRoleUserMsg(fmt.Sprintf("你是一个写作助手，请根据以下内容写出一个%s", params...))
	return engine.Push(w.msg)
}

func (w *Writer) AppendContext(msg message.Messages) Interface {
	w.msg.Append(msg)
	return w
}
