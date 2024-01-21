package handle

import "etov/internal/router"

func RegisterHandler(router *router.Router) {
	router.GET("/ping", Ping)

	router.GET("/api/chat", ChatGET)
	router.POST("/api/chat", ChatPOST)
	router.POST("/api/chat/get/chats", GetChats)
	router.POST("/api/chat/create/chatId", CreateChat)

	router.POST("/api/auth/hasRegistered", HasRegistered)
	router.POST("/api/auth/register", Register)
	router.POST("/api/auth/login", Login)
}
