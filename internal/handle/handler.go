package handle

import (
	"etov/internal/interceptor"
	"etov/internal/router"
)

func RegisterHandler(router *router.Router) {
	router.GET("/ping", Ping)
	router.Static("/api/static", "./static")

	chat := router.Group("/api/chat")
	chat.GET("", ChatGET)
	chat.POST("", ChatPOST)
	chat.POST("/get/chats", GetChats)
	chat.POST("/create/chatId", CreateChat)

	auth := router.Group("/api/auth")
	auth.POST("/hasRegistered", HasRegistered)
	auth.POST("/register", Register)
	auth.POST("/login", Login)

	common := router.Group("/api", interceptor.Authorization, interceptor.Recover)

	user := common.Group("/user")
	user.GET("/info", UserInfo)
}
