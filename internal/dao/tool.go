package dao

import (
	"gorm.io/gorm"

	"etov/internal/model"
	"etov/internal/orm"
)

type ToolDao struct {
	db *gorm.DB
}

func NewToolDao(db *gorm.DB) *ToolDao {
	return &ToolDao{db: db}
}

func (t *ToolDao) GetAll() ([]model.Tool, error) {
	//TODO implement me
	panic("implement me")
}

func (t *ToolDao) GetAllPublic() ([]model.Tool, error) {
	var tools []model.Tool
	if err := t.db.Model(&model.Tool{}).Where("is_public = 'Y'").Order("created_at ASC").Find(&tools).Error; err != nil {
		return nil, err
	}
	return tools, nil
}

func (t *ToolDao) GetByFields(wheres ...orm.Wheres) ([]model.Tool, error) {
	//TODO implement me
	panic("implement me")
}
