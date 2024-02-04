package tools

import (
	"fmt"

	"github.com/sirupsen/logrus"

	"etov/internal/gpt/engine"
	"etov/internal/gpt/message"
	"etov/internal/request"
	"etov/internal/svc"
)

func ReduceDuplication(ctx *svc.Context) {
	var (
		e   engine.ChatEngine
		req = request.ReduceDuplication{}
	)

	if err := ctx.ShouldBind(&req); err != nil {
		logrus.Error(err)
		ctx.Error(err)
		return
	}

	e, err := ctx.ClientCache.GetClient(req.EngineId, ctx.DB)
	if err != nil {
		logrus.Error(err)
		ctx.Error(err)
		return
	}
	messages := message.NewMessages()
	messages.AddChatMessageGPTMsg(paresMode(req.Content, req.Mode))
	logrus.Infoln("论文降重：", req.Content)
	sess, err := e.Push(messages)
	if err != nil {
		logrus.Error(err)
		ctx.Error(err)
		return
	}
	ctx.Stream(sess.HandleStream(messages))
}

func paresMode(content, mode string) string {
	switch mode {
	case "1":
		return fmt.Sprintf("`%s` 将这段话简明阐述", content)
	case "2":
		return fmt.Sprintf("`%s` 将这段话换一种表述方式", content)
	case "3":
		return fmt.Sprintf("`%s` 保证这段话意思不变，并补充点内容", content)
	}
	return fmt.Sprintf("`%s` 将这段话换一种表述方式", content)
}
