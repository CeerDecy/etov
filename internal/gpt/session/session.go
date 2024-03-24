package session

import (
	"io"
	"sync"

	"github.com/sashabaranov/go-openai"
	_ "github.com/sashabaranov/go-openai"

	"etov/internal/gpt/message"
)

type Session struct {
	stream  *openai.ChatCompletionStream
	buf     []byte
	content []byte
	done    bool
	sign    chan struct{}
	sync.Mutex
}

func NewSession(stream *openai.ChatCompletionStream) *Session {
	return &Session{
		stream:  stream,
		buf:     make([]byte, 0),
		content: make([]byte, 0),
		done:    false,
		sign:    make(chan struct{}, 1),
		Mutex:   sync.Mutex{},
	}
}

func (s *Session) readBuf(msg *message.Messages) []byte {
	s.Lock()
	s.Unlock()
	if s.done && msg != nil {
		msg.AddChatMessageGPTMsg(string(s.content))
	}
	res := make([]byte, len(s.buf))
	copy(res, s.buf)
	s.buf = make([]byte, 0)
	return res
}

func (s *Session) HandleStream(msg *message.Messages) func(w io.Writer) bool {
	return func(w io.Writer) bool {
		bytes := s.readBuf(msg)
		var err error
		if len(bytes) > 0 {
			_, err = w.Write(bytes)
		}
		select {
		case <-s.sign:
			return false
		default:
			return err == nil
		}
	}
}

func (s *Session) ReadStream() {
	go func() {
		for {
			recv, err := s.stream.Recv()
			if err != nil {
				s.stream.Close()
				s.Lock()
				s.done = true
				s.Unlock()
				s.sign <- struct{}{}
				return
			}
			for _, v := range recv.Choices {
				s.Lock()
				s.buf = append(s.buf, v.Delta.Content...)
				s.content = append(s.content, v.Delta.Content...)
				s.Unlock()
			}
		}
	}()
}
