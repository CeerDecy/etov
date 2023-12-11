package router

import (
	"github.com/gin-gonic/gin"

	"etov/internal/handle"
)

func DefaultRouter(router *gin.Engine) {
	router.GET("/ping", handle.Ping)
	router.GET("/chat", handle.Chat)
}
