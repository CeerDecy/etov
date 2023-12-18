package repo

import (
	"etov/internal/model"
)

type UserRepo interface {
	Create(user *model.User) error
	GetByID(id int64) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	GetAll() ([]*model.User, error)
	Update(user *model.User) error
}
