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
	Cache       *chat.Cache
	GPT         *gptclient.GptClient
	*gin.Context
}

func NewContext(conf *conf.EtovConfig) *Context {
	db := client.ConnectDB(conf.Mysql)
	return &Context{
		DB:          db,
		RedisClient: client.ConnectRedis(conf.Redis),
		Cache:       chat.NewCache(conf.Cache.TTL, conf.Cache.Size, db),
		GPT:         gptclient.DefaultClient(conf.OpenAI),
	}
}

func NewContextFromAddon(middle *Addons) *Context {
	return &Context{
		DB:          middle.DB,
		RedisClient: middle.RedisClient,
		Cache:       middle.Cache,
		GPT:         middle.GPT,
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
