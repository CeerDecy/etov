package tools

import (
	"etov/internal/gpt/etovs"
	"etov/internal/request"
	"etov/internal/svc"
)

func Translator(ctx *svc.Context) {
	var req request.TranslatorRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.Error(err)
		return
	}

	client, err := ctx.ClientCache.GetClient(req.EngineId, ctx.DB)
	if err != nil {
		ctx.Error(err)
		return
	}

	translator := etovs.NewTranslator()
	sess, err := translator.Execute(client, req.Content, req.TargetLang)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.Stream(sess.HandleStream(nil))
}
