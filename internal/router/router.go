package router

import (
	"github.com/gin-gonic/gin"

	"etov/internal/handle"
)

func Route(router *gin.Engine) {
	router.GET("/ping", handle.Ping)
}
