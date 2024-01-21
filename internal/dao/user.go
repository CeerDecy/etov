package dao

import (
	"gorm.io/gorm"

	"etov/internal/model"
	"etov/internal/orm"
)

type UserDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{db: db}
}

func (u *UserDao) Create(user *model.User) error {
	if err := u.db.Create(user).Error; err != nil {
		return err
	}
	return nil
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

func (u *UserDao) GetByPhone(phone string) (*model.User, error) {
	var user model.User
	if err := u.db.Model(&model.User{}).Where("phone = ?", phone).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserDao) GetByFields(wheres ...orm.Wheres) ([]model.User, error) {
	var user []model.User
	err := orm.MountTransaction(u.db.Model(&model.User{}), wheres...).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserDao) GetAll() ([]*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserDao) Update(user *model.User) error {
	//TODO implement me
	panic("implement me")
}
