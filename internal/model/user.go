package model

import "time"

type User struct {
	Id        int64      `gorm:"column:id;primaryKey;autoIncrement:true"`
	NickName  string     `gorm:"column:nickname"`
	Email     string     `gorm:"column:email"`
	Phone     string     `gorm:"column:phone"`
	Password  string     `gorm:"column:password"`
	Salt      string     `gorm:"column:salt"`
	Avatar    string     `gorm:"column:avatar"`
	ApiKey    string     `gorm:"column:api_key"`
	Validate  string     `gorm:"column:validate"`
	CreatedAt time.Time  `gorm:"column:created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at"`
}

func (u *User) TableName() string {
	return "users"
}
