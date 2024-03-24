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
	if err := a.db.Model(&model.APIKey{}).Where("uid = ? and is_deleted = 0", uid).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func (a *APIKeyDao) GetEngineByiId(id int64) (model.APIKey, error) {
	var res model.APIKey
	if err := a.db.Model(&model.APIKey{}).Where("id = ?", id).First(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (a *APIKeyDao) SaveAPIKey(key model.APIKey) error {
	if err := a.db.Save(&key).Error; err != nil {
		return err
	}
	return nil
}

func (a *APIKeyDao) DeleteAPIKey(id int64) error {
	if err := a.db.Model(&model.APIKey{}).Where("id = ?", id).Update("is_deleted", 1).Error; err != nil {
		return err
	}
	return nil
}
