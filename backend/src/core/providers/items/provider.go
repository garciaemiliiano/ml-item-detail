package items

import (
	"context"
	entity "item-detail-api/src/core/entities/items"

	"github.com/google/uuid"
)

type ListConfig struct {
	ID           uuid.UUID
	Limit        int
	Offset       int
	Count        bool
	WithDetails  bool
	CategoryName string
	ProductName  string
}

//go:generate mockery --all
type ItemsPersistor interface {
	List(ctx context.Context, config ListConfig) ([]entity.Item, int, error)
}
