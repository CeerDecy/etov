package model

import "time"

// User /**
/**
CREATE TABLE USERS
(
    `id`         INTEGER PRIMARY KEY AUTO_INCREMENT,
    `nickname`   varchar(255) DEFAULT '' COMMENT '昵称',
    `email`      varchar(255) DEFAULT '' COMMENT '邮箱',
    `phone`      char(11) NOT NULL COMMENT '手机号',
    `avatar`     char(11)     DEFAULT '' COMMENT '手机号',
    `validate`   char(1)      DEFAULT 'N' COMMENT '是否通过学生验证',
    `created_at` datetime     DEFAULT NULL COMMENT '创建时间',
    `updated_at` datetime     DEFAULT NULL COMMENT '更新时间',
    `deleted_at` datetime     DEFAULT NULL COMMENT '删除时间',
    INDEX `index_create_at` (`created_at`),
    INDEX `index_update_at` (`updated_at`),
    INDEX `index_delete_at` (`deleted_at`)
);
*/
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
