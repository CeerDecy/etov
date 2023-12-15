package chat

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	cmap "github.com/orcaman/concurrent-map"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"etov/internal/gpt/message"
)

const TempChatIdPrefix = "temp-"

func GenerateTempChatId() string {
	u := uuid.New()
	return TempChatIdPrefix + strings.ReplaceAll(u.String(), "-", "")
}

type Cache struct {
	db    *gorm.DB
	cache *cmap.ConcurrentMap
}

func NewCache(DB *gorm.DB) *Cache {
	cm := cmap.New()
	return &Cache{
		db:    DB,
		cache: &cm,
	}
}

func (c *Cache) StoreMessages(chatId string, msg *message.Messages) {
	c.cache.Set(chatId, msg)
}

func (c *Cache) GetMessages(chatId string) (*message.Messages, error) {
	val, exists := c.cache.Get(chatId)
	if !exists {
		if strings.HasPrefix(chatId, TempChatIdPrefix) {
			c.StoreMessages(chatId, message.NewMessages())
		} else {
			logrus.Info("find in database")
			//c.findAndSave(chatId)
		}
	}
	messages, ok := val.(*message.Messages)
	if !ok {
		err := fmt.Errorf("unable convert to *message.Messages")
		return nil, err
	}
	return messages, nil
}

func (c *Cache) findAndSave(chatId string) {

}
