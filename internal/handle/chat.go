package handle

import (
	"io"
	"net/http"

	"github.com/sirupsen/logrus"

	"etov/internal/gpt/chat"
	"etov/internal/gpt/message"
	"etov/internal/gpt/session"
	"etov/internal/request"
	"etov/internal/response"
	"etov/internal/svc"
)

func CreateChat(ctx *svc.Context) {
	uid, exists := ctx.Get("uid")
	var chatId string
	if exists {
		_ = uid
		// db ...
	} else {
		chatId = chat.GenerateTempChatId()
		ctx.ChatCache.StoreMessages(chatId, message.NewMessages())
	}
	ctx.Success(response.NewCreateChatResponse(chatId))
}

func ChatGET(ctx *svc.Context) {
	var chat request.ChatRequest
	value, _ := ctx.GetQuery("content")
	chat.Content = value
	messages := message.NewMessages()
	messages.AddChatMessageRoleUserMsg(chat.Content)
	stream, err := ctx.GPT.GetStreamResponse(messages)
	if err != nil {
		ctx.Error(err)
		return
	}
	sess := session.NewSession(stream)
	go sess.ReadStream()

	if sess.Done {
		ctx.String(http.StatusOK, string(sess.Content))
	} else {
		ctx.Writer.Header().Set("Content-Type", "text/event-stream;charset=utf-8")
		ctx.Stream(func(w io.Writer) bool {
			res := sess.ReadResp()
			_, err := w.Write(res)
			select {
			case <-sess.Sign:
				return false
			default:
				return err == nil
			}
		})
	}
}

func ChatPOST(ctx *svc.Context) {
	var chatReq request.ChatRequest
	err := ctx.ShouldBind(&chatReq)
	if err != nil {
		logrus.Error(err.Error())
		ctx.Error(err)
		return
	}
	msg, err := ctx.ChatCache.GetMessages(chatReq.ChatId)
	if err != nil {
		logrus.Error(err.Error())
		ctx.Error(err)
		return
	}
	msg.AddChatMessageRoleUserMsg(chatReq.Content)
	stream, err := ctx.GPT.GetStreamResponse(msg)
	if err != nil {
		ctx.Error(err)
		return
	}
	sess := session.NewSession(stream)
	go sess.ReadStream()

	if sess.Done {
		ctx.String(http.StatusOK, string(sess.Content))
	} else {
		ctx.Writer.Header().Set("Content-Type", "text/event-stream;charset=utf-8")
		ctx.Stream(func(w io.Writer) bool {
			res := sess.ReadResp()
			_, err := w.Write(res)
			select {
			case <-sess.Sign:
				msg.AddChatMessageGPTMsg(string(sess.Content))
				return false
			default:
				return err == nil
			}
		})
	}
}
