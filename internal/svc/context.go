package svc

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"gorm.io/gorm"

	"etov/client"
	"etov/conf"
	"etov/internal/cache"
	"etov/internal/gpt/chat"
	"etov/internal/gpt/gptclient"
	"etov/internal/response"
)

type HandlerFunc func(ctx *Context)

type MiddleFunc func(next HandlerFunc) HandlerFunc

type Context struct {
	DB          *gorm.DB
	RedisClient *redis.Client
	Cache       *chat.Cache
	ClientCache *cache.GptClientCache
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
		Cache:       middle.ChatCache,
		ClientCache: middle.ClientCache,
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

// ErrorMsg Error 返回错误
func (e *Context) ErrorMsg(msg string) {
	e.JSON(http.StatusOK, response.ErrorMsgResp(msg))
}

func (e *Context) GetUserId() (int64, bool) {
	value, exists := e.Get("userID")
	if exists {
		return value.(int64), true
	}
	return 0, false
}
