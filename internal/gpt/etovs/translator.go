package etovs

import (
	"fmt"

	"etov/internal/gpt/engine"
	"etov/internal/gpt/message"
	"etov/internal/gpt/session"
)

type Translator struct {
	msg *message.Messages
}

func (t *Translator) AppendContext(msg message.Messages) Interface {
	t.msg.Append(msg)
	return t
}

func NewTranslator() Interface {
	msg := message.NewMessages()
	msg.AddChatMessageRoleUserMsg("你是一个AI翻译器，请按照要求将文本翻译成目标语言，不要输出多余的话，只需要翻译结果")
	return &Translator{
		msg: msg,
	}
}

func (t *Translator) Execute(engine engine.ChatEngine, params ...any) (*session.Session, error) {
	t.msg.AddChatMessageRoleUserMsg(fmt.Sprintf("翻译`%s`这段文字为%s", params...))
	return engine.Push(t.msg)
}
