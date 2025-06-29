package items

import (
	"item-detail-api/src/core/entities/products"
	"item-detail-api/src/core/entities/providers"
	"time"

	"github.com/google/uuid"
)

type Item struct {
	ID           uuid.UUID
	ProductID    uuid.UUID
	ProviderID   uuid.UUID
	SerialNumber string
	Stock        int
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
	Product      products.Product
	Provider     providers.Provider
}
