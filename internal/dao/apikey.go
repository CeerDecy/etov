package dao

import (
	"gorm.io/gorm"

	"etov/internal/model"
)

type APIKeyDao struct {
	db *gorm.DB
}

func NewAPIKeyDao(db *gorm.DB) *APIKeyDao {
	return &APIKeyDao{db: db}
}

func (a *APIKeyDao) GetEngineByUid(uid int64) ([]model.APIKey, error) {
	var res []model.APIKey
	if err := a.db.Model(&model.APIKey{}).Where("uid = ?", uid).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
