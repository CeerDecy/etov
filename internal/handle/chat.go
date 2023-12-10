package handle

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"

	"etov/internal/gpt/gptclient"
	"etov/internal/model"
	"etov/internal/session"
	"etov/internal/tool"
)

func Chat(ctx *gin.Context) {
	var chat model.ChatRequest
	value, _ := ctx.GetQuery("content")
	chat.Content = value
	client := gptclient.DefaultClient()
	stream, err := client.GetStreamResponse(chat.Content)
	if err != nil {
		tool.NewEtovCtx(ctx).Error(err.Error())
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
