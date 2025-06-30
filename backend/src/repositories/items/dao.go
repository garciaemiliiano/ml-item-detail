package items

import (
	entity "item-detail-api/src/core/entities/items"
	"item-detail-api/src/core/utils"
	"item-detail-api/src/repositories/basemodel"
	"item-detail-api/src/repositories/products"
	"item-detail-api/src/repositories/providers"

	"github.com/google/uuid"
)

type ItemDAO struct {
	basemodel.BaseModel
	ProductID    uuid.UUID `gorm:"column:product_id"`
	ProviderID   uuid.UUID `gorm:"column:provider_id"`
	SerialNumber string    `gorm:"column:serial_number"`
	Stock        int       `gorm:"column:stock"`

	Product  products.ProductDAO   `gorm:"foreignKey:ProductID;references:ID"`
	Provider providers.ProviderDAO `gorm:"foreignKey:ProviderID;references:ID"`
}

func (ItemDAO) TableName() string {
	return "items"
}

func (m ItemDAO) ToEntity() entity.Item {
	return entity.Item{
		ID:           m.ID,
		ProductID:    m.ProductID,
		ProviderID:   m.ProviderID,
		SerialNumber: m.SerialNumber,
		Stock:        m.Stock,
		CreatedAt:    utils.ParseDateStr(m.CreatedAt),
		UpdatedAt:    utils.ParseDateStr(m.UpdatedAt),
		Product:      m.Product.ToEntity(),
		Provider:     m.Provider.ToEntity(),
	}
}

func toEntities(daos []ItemDAO) []entity.Item {
	entities := make([]entity.Item, 0)
	for _, dao := range daos {
		entities = append(entities, dao.ToEntity())
	}
	return entities
}
