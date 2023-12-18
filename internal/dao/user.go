package dao

import (
	"gorm.io/gorm"

	"etov/internal/model"
)

type UserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{db: db}
}

func (u *UserDao) Create(user *model.User) error {
	//TODO implement me
	panic("implement me")
}

func (u *UserDao) GetByID(id int64) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserDao) GetByEmail(email string) (*model.User, error) {
	var user model.User
	if err := u.db.Model(&model.User{}).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserDao) GetAll() ([]*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserDao) Update(user *model.User) error {
	//TODO implement me
	panic("implement me")
}
