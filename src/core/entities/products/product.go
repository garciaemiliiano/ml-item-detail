package products

import (
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
	CategoryID  uuid.UUID
}
