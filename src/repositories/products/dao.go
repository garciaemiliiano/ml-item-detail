package products

import (
	"item-detail-api/src/repositories/basemodel"
	"item-detail-api/src/repositories/categories"

	"github.com/google/uuid"
)

type ProductDAO struct {
	basemodel.BaseModel
	Name        string                 `gorm:"column:name"`
	Price       float64                `gorm:"column:price"`
	Description string                 `gorm:"column:description"`
	CategoryID  uuid.UUID              `gorm:"column:category_id"`
	Category    categories.CategoryDAO `gorm:"foreignKey:CategoryID;references:ID"`
}
