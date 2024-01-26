package repo

import (
	"etov/internal/model"
	"etov/internal/orm"
)

type ToolRepo interface {
	GetAll() ([]model.Tool, error)
	GetAllPublic() ([]model.Tool, error)
	GetByFields(wheres ...orm.Wheres) ([]model.Tool, error)
}
