package products

import (
	"item-detail-api/src/core/entities/categories"
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID          uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
	Name        string
	Price       float64
	Description string
	ImageURL    string
	CategoryID  uuid.UUID
	Category    categories.Category
}
