package categories

import (
	"item-detail-api/src/core/entities/categories"
	"item-detail-api/src/core/utils"
	"item-detail-api/src/repositories/basemodel"
)

type CategoryDAO struct {
	basemodel.BaseModel
	Name string `gorm:"column:name"`
}

func (CategoryDAO) TableName() string {
	return "categories"
}

func (m CategoryDAO) ToEntity() categories.Category {
	return categories.Category{
		ID:        m.ID,
		Name:      m.Name,
		CreatedAt: utils.ParseDateStr(m.CreatedAt),
		UpdatedAt: utils.ParseDateStr(m.UpdatedAt),
	}
}
