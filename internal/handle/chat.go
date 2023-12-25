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

	if sess.Done {
		ctx.Stream(func(w io.Writer) bool {
			_, err := w.Write(sess.Content)
			return err == nil
		})
	} else {
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
	ctx.Writer.Header().Set("Content-Type", "text/event-stream;charset=utf-8")
	err := ctx.ShouldBind(&chatReq)
	if err != nil {
		logrus.Error(err.Error())
		ctx.Stream(func(w io.Writer) bool {
			_, _ = w.Write([]byte(err.Error()))
			return false
		})
		return
	}
	logrus.Info(chatReq.Content)
	logrus.Info("get cache chatId ", chatReq.ChatId)
	ca, err := ctx.Cache.Get(chatReq.ChatId)
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
		ctx.Error(err)
		return
	}
	msg.AddChatMessageRoleUserMsg(chatReq.Content)
	stream, err := ctx.GPT.GetStreamResponse(msg)
	if err != nil {
		logrus.Info(chatReq.Content)
		logrus.Error(err)
		ctx.Stream(func(w io.Writer) bool {
			_, _ = w.Write([]byte(err.Error()))
			return false
		})
		return
	}
	sess := session.NewSession(stream)
	go sess.ReadStream()

	if sess.Done {
		ctx.Stream(func(w io.Writer) bool {
			_, _ = w.Write(sess.Content)
			return false
		})
	} else {
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
