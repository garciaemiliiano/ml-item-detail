package products

import (
	"item-detail-api/src/core/entities/products"
	"item-detail-api/src/core/utils"
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
	ImageURL    string                 `gorm:"column:image_url"`
	Category    categories.CategoryDAO `gorm:"foreignKey:CategoryID;references:ID"`
}

func (ProductDAO) TableName() string {
	return "products"
}

func (m ProductDAO) ToEntity() products.Product {
	return products.Product{
		ID:          m.ID,
		Name:        m.Name,
		Price:       m.Price,
		Description: m.Description,
		CategoryID:  m.CategoryID,
		ImageURL:    m.ImageURL,
		CreatedAt:   utils.ParseDateStr(m.CreatedAt),
		UpdatedAt:   utils.ParseDateStr(m.UpdatedAt),
		Category:    m.Category.ToEntity(),
	}
}
