package model

import (
	"time"
)

type Tool struct {
	ID          int        `gorm:"primaryKey" json:"id"`
	Name        string     `gorm:"column:name" json:"name"`
	Logo        string     `gorm:"column:logo" json:"logo"`
	URL         string     `gorm:"column:url" json:"url"`
	Description string     `gorm:"column:description" json:"description"`
	IsPublic    string     `gorm:"column:is_public" json:"is_public"`
	Disable     string     `gorm:"column:disabled" json:"disable"`
	AuthorID    int        `gorm:"column:author_id" json:"author_id"`
	CreatedAt   time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt   *time.Time `gorm:"column:deleted_at;index" json:"deleted_at"`
}

func (t *Tool) TableName() string {
	return "tool"
}
