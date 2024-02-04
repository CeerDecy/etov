package repo

import "etov/internal/model"

type APIKeyRepo interface {
	GetEngineByUid(uid int64) ([]model.APIKey, error)
}
