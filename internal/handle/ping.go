package handle

import (
	"github.com/gin-gonic/gin"

	"etov/internal/svc"
)

func Ping(ctx *svc.Context) {
	ctx.JSON(200, gin.H{
		"msg": "pong!",
	})
}
