package model

import "time"

type Chat struct {
	Id        int64      `json:"id"`
	Uid       int64      `json:"uid"`
	Title     string     `json:"title"`
	CreatedAt time.Time  `gorm:"column:created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at"`
}
