package items

import (
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
}
