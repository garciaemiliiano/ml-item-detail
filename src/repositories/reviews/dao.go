package reviews

import (
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
