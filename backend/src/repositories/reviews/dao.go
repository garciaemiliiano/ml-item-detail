package reviews

import (
	"item-detail-api/src/core/entities/reviews"
	"item-detail-api/src/repositories/basemodel"

	"github.com/google/uuid"
)

type ReviewDAO struct {
	basemodel.BaseModel
	Description string    `gorm:"column:description"`
	Rating      int       `gorm:"column:rating"`
	ProductID   uuid.UUID `gorm:"column:product_id"`
	UserID      uuid.UUID `gorm:"column:user_id"`
}

func (ReviewDAO) TableName() string {
	return "reviews"
}

func (m ReviewDAO) ToEntity() reviews.Review {
	return reviews.Review{
		ID:          m.ID,
		Description: m.Description,
		Rating:      m.Rating,
		ProductID:   m.ProductID,
		UserID:      m.UserID,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}
}

func ToEntities(daos []ReviewDAO) []reviews.Review {
	entities := make([]reviews.Review, 0)
	for _, dao := range daos {
		entities = append(entities, dao.ToEntity())
	}
	return entities
}
