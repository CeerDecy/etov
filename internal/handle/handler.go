package handle

import (
	"etov/internal/handle/tools"
	"etov/internal/interceptor"
	"etov/internal/router"
)

func RegisterHandler(router *router.Router) {
	router.GET("/ping", Ping)
	router.Static("/api/static", "./static")
	router = router.Group("", interceptor.Recover)

	router.POST("/api/file/upload", tools.FileUpload)

	auth := router.Group("/api/auth")
	auth.POST("/hasRegistered", HasRegistered)
	auth.POST("/register", Register)
	auth.POST("/login", Login)

	chat := router.Group("/api/chat")
	chat.GET("", ChatGET)
	chat.POST("", ChatPOST)
	chat.POST("/get/chats", GetChats)
	chat.POST("/create/chatId", CreateChat)

	toolCommon := router.Group("/api/tool")
	toolCommon.GET("/get/public", GetPublicTools)
	toolCommon.POST("/reduce-duplication", tools.ReduceDuplication)
	toolCommon.POST("/translator", tools.Translator)
	toolCommon.POST("/summary", tools.Summary)
	toolCommon.POST("/write", tools.Write)

	engineRouter := router.Group("/api/engine", interceptor.AuthorizationNonMandatory)
	engineRouter.GET("/get/support", GetSupportEngine)
	engineRouter.GET("/get/apikeys", GetAPIKeys)
	engineRouter.POST("/create/apikey", SaveAPIKey)
	engineRouter.POST("/update/apikey", UpdateToken)
	engineRouter.POST("/delete/apikey", DeleteToken)

	common := router.Group("/api", interceptor.AuthorizationMandatory)

	user := common.Group("/user")
	user.GET("/info", UserInfo)
}
