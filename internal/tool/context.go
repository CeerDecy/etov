package tool

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type EtovCtx struct {
	*gin.Context
}

func NewEtovCtx(ctx *gin.Context) *EtovCtx {
	return &EtovCtx{
		ctx,
	}
}

func (e *EtovCtx) Success() {
	e.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "success"})
}

func (e *EtovCtx) SuccessData(data any) {
	e.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "success", "data": data})
}

// Error 返回错误
func (e *EtovCtx) Error(msg string) {
	e.JSON(http.StatusOK, gin.H{"code": http.StatusInternalServerError, "msg": msg})
}

func (e *EtovCtx) Build() *gin.Context {
	return e.Context
}
