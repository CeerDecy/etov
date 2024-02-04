package tools

import (
	"fmt"
	"strconv"

	"github.com/sirupsen/logrus"

	"etov/internal/dao"
	"etov/internal/gpt/engine"
	"etov/internal/gpt/message"
	"etov/internal/repo"
	"etov/internal/request"
	"etov/internal/svc"
)

func ReduceDuplication(ctx *svc.Context) {
	var (
		apiKeyRepo repo.APIKeyRepo = dao.NewAPIKeyDao(ctx.DB)
		e          engine.ChatEngine
		req        = request.ReduceDuplication{}
	)

	if err := ctx.ShouldBind(&req); err != nil {
		logrus.Error(err)
		ctx.Error(err)
		return
	}

	cc, err := ctx.ClientCache.Get(req.EngineId)
	if err != nil {
		apiId, _ := strconv.Atoi(req.EngineId)
		apikey, err := apiKeyRepo.GetEngineByiId(int64(apiId))
		if err != nil {
			logrus.Error(err)
			ctx.Error(err)
			return
		}
		e = engine.NewChatEngine(apikey.ModelTag, apikey.APIKey, apikey.Host)
		_ = ctx.ClientCache.Set(req.EngineId, e)
	} else {
		logrus.Info("use cache")
		e = cc.(engine.ChatEngine)
	}
	messages := message.NewMessages()
	messages.AddChatMessageGPTMsg(paresMode(req.Content, req.Mode))
	logrus.Println("messages:", messages)
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
		logrus.Println("mode 1")
		return fmt.Sprintf("`%s` 将这段话简明阐述", content)
	case "2":
		logrus.Println("mode 2")
		return fmt.Sprintf("`%s` 将这段话换一种表述方式", content)
	case "3":
		logrus.Println("mode 3")
		return fmt.Sprintf("`%s` 保证这段话意思不变，并补充点内容", content)
	}
	return fmt.Sprintf("`%s` 将这段话换一种表述方式", content)
}
