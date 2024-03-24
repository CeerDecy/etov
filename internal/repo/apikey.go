package repo

import "etov/internal/model"

type APIKeyRepo interface {
	GetEngineByUid(uid int64) ([]model.APIKey, error)
	GetEngineByiId(id int64) (model.APIKey, error)
	SaveAPIKey(key model.APIKey) error
	DeleteAPIKey(id int64) error
}
