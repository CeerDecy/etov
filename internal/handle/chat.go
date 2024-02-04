package handle

import (
	"fmt"
	"io"
	"strconv"

	"github.com/sirupsen/logrus"

	"etov/internal/dao"
	"etov/internal/gpt/chat"
	"etov/internal/gpt/message"
	"etov/internal/gpt/session"
	"etov/internal/request"
	"etov/internal/response"
	"etov/internal/svc"
)

func GetChats(ctx *svc.Context) {
	var resp = response.GetChatsResponse{Chats: make([]*response.ChatItem, 0)}
	uid, exists := ctx.Get("uid")
	if !exists {
		ctx.Success(resp)
		return
	}
	chatsDao := dao.NewChatsDao(ctx.DB)
	chats, err := chatsDao.GetChatByUid(uid.(int64))
	if err != nil {
		ctx.Error(err)
		return
	}
	for _, v := range chats {
		resp.Chats = append(resp.Chats, &response.ChatItem{Id: strconv.FormatInt(v.Id, 10), Title: v.Title})
	}
	ctx.Success(resp)
}

func CreateChat(ctx *svc.Context) {
	var resp = response.CreateChatResponse{}
	uid, exists := ctx.Get("uid")
	var chatId string
	if exists {
		_ = uid
		// db ...
	} else {
		chatId = chat.GenerateTempChatId()
		logrus.Info("store chatId: ", chatId)
		err := ctx.Cache.Set(chatId, message.NewMessages())
		if err != nil {
			ctx.Error(err)
			return
		}
		resp.Chat = &response.ChatItem{Id: chatId, Title: chatId}
	}
	ctx.Success(resp)
}

func ChatGET(ctx *svc.Context) {
	var req request.ChatRequest
	ctx.Writer.Header().Set("Content-Type", "text/event-stream;charset=utf-8")
	value, _ := ctx.GetQuery("content")
	req.Content = value
	messages := message.NewMessages()
	messages.AddChatMessageRoleUserMsg(req.Content)
	stream, err := ctx.GPT.GetStreamResponse(messages)
	if err != nil {
		ctx.Stream(func(w io.Writer) bool {
			_, err := w.Write([]byte(err.Error()))
			return err == nil
		})
		return
	}
	sess := session.NewSession(stream)
	go sess.ReadStream()

	ctx.Stream(sess.HandleStream(messages))
}

func ChatPOST(ctx *svc.Context) {
	var req request.ChatRequest
	ctx.Writer.Header().Set("Content-Type", "text/event-stream;charset=utf-8")
	if err := ctx.ShouldBind(&req); err != nil {
		logrus.Error(err.Error())
		ctx.Stream(func(w io.Writer) bool {
			_, _ = w.Write([]byte(err.Error()))
			return false
		})
		return
	}
	logrus.Info(req.Content)
	logrus.Info("get cache chatId ", req.ChatId)
	ca, err := ctx.Cache.Get(req.ChatId)
	if err != nil {
		logrus.Error(err.Error())
		ctx.Stream(func(w io.Writer) bool {
			_, _ = w.Write([]byte(err.Error()))
			return false
		})
		return
	}
	msg, ok := ca.(*message.Messages)
	if !ok {
		err = fmt.Errorf("cannot convert to *message.Messages")
		logrus.Error(err)
		ctx.Error(fmt.Errorf("您的临时会话已过期，请重新刷新开始对话；或者登录可以自动保存对话记录，下次可继续对话"))
		return
	}
	msg.AddChatMessageRoleUserMsg(req.Content)
	chatGPT, err := ctx.ClientCache.GetClient(req.EngineId, ctx.DB)
	if err != nil {
		logrus.Error(err)
		ctx.Error(err)
		return
	}
	sess, err := chatGPT.Push(msg)
	if err != nil {
		logrus.Error("ChatGPT.Push error: ", err)
		ctx.Error(err)
		return
	}
	ctx.Stream(sess.HandleStream(msg))

}
