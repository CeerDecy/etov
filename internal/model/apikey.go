package model

import (
	"time"
)

type APIKey struct {
	ID        int64      `gorm:"column:id;primaryKey"`
	UID       int64      `gorm:"column:uid"`
	KeyName   string     `gorm:"column:key_name"`
	APIKey    string     `gorm:"column:apikey"`
	ModelTag  string     `gorm:"column:model_tag"`
	Host      string     `gorm:"column:host"`
	CreatedAt time.Time  `gorm:"column:created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at"`
}

func (a *APIKey) TableName() string {
	return "apikey"
}
