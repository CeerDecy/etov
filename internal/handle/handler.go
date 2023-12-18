package handle

import "etov/internal/router"

func RegisterHandler(router *router.Router) {
	router.GET("/ping", Ping)

	router.GET("/chat", ChatGET)
	router.POST("/chat", ChatPOST)
	router.POST("/chat/create/chatId", CreateChat)

	router.POST("/auth/HasRegistered", HasRegistered)
}
