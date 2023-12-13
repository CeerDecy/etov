package svc

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"gorm.io/gorm"

	"etov/client"
	"etov/conf"
)

type HandlerFunc func(ctx *Context)

type Context struct {
	DB          *gorm.DB
	RedisClient *redis.Client
	*gin.Context
}

func NewContext(conf *conf.EtovConfig) *Context {
	return &Context{
		DB:          client.ConnectDB(conf.Mysql),
		RedisClient: client.ConnectRedis(conf.Redis),
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
