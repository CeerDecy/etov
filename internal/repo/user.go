package repo

import (
	"etov/internal/model"
	"etov/internal/orm"
)

type UserRepo interface {
	Create(user *model.User) error
	GetByID(id int64) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	GetByPhone(phone string) (*model.User, error)
	GetByFields(wheres ...orm.Wheres) ([]model.User, error)
	GetAll() ([]*model.User, error)
	Update(user *model.User) error
}
