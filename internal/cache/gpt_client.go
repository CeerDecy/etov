package cache

import (
	"github.com/bluele/gcache"
)

type GptClientCache struct {
	gcache.Cache
}

func NewGptClientCache(size int) *GptClientCache {
	return &GptClientCache{
		Cache: gcache.New(size).LFU().Build(),
	}
}
