package svc

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"etov/conf"
	"etov/db"
)

type HandlerFunc func(etov *Context)

type Context struct {
	DB *gorm.DB
	*gin.Context
}

func NewContext(conf *conf.EtovConfig) *Context {
	DB := db.ConnectDB(conf.Mysql)
	return &Context{
		DB: DB,
	}
}

func (e *Context) Default() {
	e.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "success"})
}

func (e *Context) Success(code int, msg string) {
	e.JSON(http.StatusOK, gin.H{"code": code, "msg": msg})
}

func (e *Context) SuccessData(data any) {
	e.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "success", "data": data})
}

// Error 返回错误
func (e *Context) Error(msg string) {
	e.JSON(http.StatusOK, gin.H{"code": http.StatusInternalServerError, "msg": msg})
}
