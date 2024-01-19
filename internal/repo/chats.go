package repo

import "etov/internal/model"

type ChatsRepo interface {
	GetChats() ([]*model.Chat, error)
	GetChat(id string) (*model.Chat, error)
	GetChatByUid(uid string) (*model.Chat, error)
	CreateChat(chat *model.Chat) (*model.Chat, error)
	UpdateChat(chat *model.Chat) (*model.Chat, error)
	DeleteChat(id string) error
}
