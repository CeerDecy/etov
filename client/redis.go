package client

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"

	r "etov/conf/redis"
)

func ConnectRedis(cfg r.Redis) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
	})
	_, err := client.Ping().Result()
	if err != nil {
		logrus.Error(err.Error())
		return nil
	}
	return client
}
