package tools

import (
	"etov/internal/gpt/etovs"
	"etov/internal/gpt/message"
	"etov/internal/request"
	"etov/internal/svc"
)

func Write(ctx *svc.Context) {
	var req request.WriteRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.Error(err)
		return
	}
	client, err := ctx.ClientCache.GetClient(req.EngineId, ctx.DB)
	if err != nil {
		ctx.Error(err)
		return
	}

	writer := etovs.NewWriter()
	msg := message.Messages{}
	msg.AddChatMessageRoleUserMsg(req.Content)
	writer.AppendContext(msg)
	sess, err := writer.Execute(client, req.Types)
	ctx.Stream(sess.HandleStream(nil))
}
