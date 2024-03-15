package etovs

import (
	"etov/internal/gpt/engine"
	"etov/internal/gpt/message"
	"etov/internal/gpt/session"
)

type Interface interface {
	Execute(engine engine.ChatEngine, params ...any) (*session.Session, error)
	AppendContext(msg message.Messages) Interface
}
