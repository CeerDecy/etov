package handle

import (
	"io"
	"net/http"

	"github.com/sirupsen/logrus"

	"etov/internal/gpt/gptclient"
	"etov/internal/request"
	"etov/internal/session"
	"etov/internal/svc"
)

func ChatGET(ctx *svc.Context) {
	var chat request.ChatRequest
	value, _ := ctx.GetQuery("content")
	chat.Content = value
	client := gptclient.DefaultClient()
	stream, err := client.GetStreamResponse(chat.Content)
	if err != nil {
		ctx.Error(err.Error())
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
	chat := request.ChatRequest{}
	err := ctx.ShouldBind(&chat)
	if err != nil {
		logrus.Error(err.Error())
		ctx.Error(err.Error())
		return
	}
	client := gptclient.DefaultClient()
	stream, err := client.GetStreamResponse(chat.Content)
	if err != nil {
		ctx.Error(err.Error())
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
