package interceptor

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"etov/internal/svc"
	"etov/internal/utils"
)

func Authorization(next svc.HandlerFunc) svc.HandlerFunc {
	return func(ctx *svc.Context) {
		token := ctx.GetHeader("Authorization")
		if token == "" {
			ctx.JSON(http.StatusOK,gin.H{
				"code":http.StatusUnauthorized,
				"msg":"用户未登录",
			})
			return
		}
		claims, err := utils.ParseTokenHs256(token)
		if err != nil {
			ctx.Error(err)
			return
		}
		ctx.Set("userID",claims.UserID)
		next(ctx)
	}
}
