package chat

import (
	"time"

	"github.com/bluele/gcache"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Cache struct {
	gcache.Cache
	db *gorm.DB
}

func NewCache(ttl time.Duration, size int, DB *gorm.DB) *Cache {
	cache := &Cache{
		db: DB,
	}
	cache.Cache = gcache.New(size).LRU().Expiration(ttl).LoaderFunc(cache.loadChatCache).Build()
	return cache
}

func (c *Cache) loadChatCache(key any) (any, error) {
	logrus.Infof("load from db %v", key)
	return nil, nil
}
