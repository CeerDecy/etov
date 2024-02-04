package cache

import (
	"strconv"

	"github.com/bluele/gcache"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"etov/internal/dao"
	"etov/internal/gpt/engine"
	"etov/internal/repo"
)

type GptClientCache struct {
	gcache.Cache
}

func NewGptClientCache(size int) *GptClientCache {
	return &GptClientCache{
		Cache: gcache.New(size).LFU().Build(),
	}
}

func (g *GptClientCache) GetClient(engineId string, db *gorm.DB) (engine.ChatEngine, error) {
	var e engine.ChatEngine
	var apiKeyRepo repo.APIKeyRepo = dao.NewAPIKeyDao(db)
	cc, err := g.Get(engineId)
	if err != nil {
		apiId, _ := strconv.Atoi(engineId)
		apikey, err := apiKeyRepo.GetEngineByiId(int64(apiId))
		if err != nil {
			return nil, err
		}
		e = engine.NewChatEngine(apikey.ModelTag, apikey.APIKey, apikey.Host)
		_ = g.Set(engineId, e)
	} else {
		logrus.Info("use cache")
		e = cc.(engine.ChatEngine)
	}
	return e, nil
}
