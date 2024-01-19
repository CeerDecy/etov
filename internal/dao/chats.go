package dao

import (
	"gorm.io/gorm"

	"etov/internal/model"
)

type ChatsDao struct {
	db *gorm.DB
}

func NewChatsDao(db *gorm.DB) *ChatsDao {
	return &ChatsDao{db: db}
}

func (c *ChatsDao) GetChats() ([]*model.Chat, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ChatsDao) GetChat(id int64) (*model.Chat, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ChatsDao) GetChatByUid(uid int64) ([]*model.Chat, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ChatsDao) CreateChat(chat *model.Chat) (*model.Chat, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ChatsDao) UpdateChat(chat *model.Chat) (*model.Chat, error) {
	//TODO implement me
	panic("implement me")
}

func (c *ChatsDao) DeleteChat(id string) error {
	//TODO implement me
	panic("implement me")
}
