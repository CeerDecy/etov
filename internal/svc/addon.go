package svc

import (
	"github.com/go-redis/redis"
	"gorm.io/gorm"

	"etov/client"
	"etov/conf"
	"etov/internal/gpt/chat"
	"etov/internal/gpt/gptclient"
)

type Addons struct {
	DB          *gorm.DB
	RedisClient *redis.Client
	Cache       *chat.Cache
	GPT         *gptclient.GptClient
}

func NewAddons(config *conf.EtovConfig) *Addons {
	db := client.ConnectDB(config.Mysql)
	return &Addons{
		DB:          db,
		RedisClient: client.ConnectRedis(config.Redis),
		Cache:       chat.NewCache(config.Cache.TTL, config.Cache.Size, db),
		GPT:         gptclient.DefaultClient(config.OpenAI),
	}
}
