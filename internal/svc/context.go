package svc

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"gorm.io/gorm"

	"etov/client"
	"etov/conf"
	"etov/internal/gpt/chat"
	"etov/internal/gpt/gptclient"
	"etov/internal/response"
)

type HandlerFunc func(ctx *Context)

type Context struct {
	DB          *gorm.DB
	RedisClient *redis.Client
	ChatCache   *chat.Cache
	GPT         *gptclient.GptClient
	*gin.Context
}

func NewContext(conf *conf.EtovConfig) *Context {
	db := client.ConnectDB(conf.Mysql)
	return &Context{
		DB:          db,
		RedisClient: client.ConnectRedis(conf.Redis),
		ChatCache:   chat.NewCache(db),
		GPT:         gptclient.DefaultClient(conf.OpenAI),
	}
}

func (e *Context) Default() {
	e.JSON(http.StatusOK, response.SuccessResp(nil))
}

func (e *Context) Success(data any) {
	e.JSON(http.StatusOK, response.SuccessResp(data))
}

// Error 返回错误
func (e *Context) Error(err error) {
	e.JSON(http.StatusOK, response.ErrorResp(err))
}
