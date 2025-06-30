package items

import (
	"context"

	itemEntity "item-detail-api/src/core/entities/items"
	itemProvider "item-detail-api/src/core/providers/items"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type ListItemsConfig struct {
	Limit        int
	Offset       int
	Count        bool
	WithDetails  bool
	ID           uuid.UUID
	CategoryName string
	ProductName  string
}

//go:generate mockery --all
type ListItems interface {
	Execute(ctx context.Context, config ListItemsConfig) ([]itemEntity.Item, int, error)
}
type ListItemsImpl struct {
	logger        *zap.Logger
	itemPersistor itemProvider.ItemsPersistor
}

func NewListItemsImpl(
	logger *zap.Logger,
	itemPersistor itemProvider.ItemsPersistor) ListItemsImpl {
	return ListItemsImpl{
		logger:        logger,
		itemPersistor: itemPersistor,
	}
}

func (uc ListItemsImpl) Execute(ctx context.Context, params ListItemsConfig) ([]itemEntity.Item, int, error) {
	items, count, err := uc.itemPersistor.List(ctx, params.toListConfig())
	if err != nil {
		return []itemEntity.Item{}, 0, err
	}

	return items, count, nil
}

func (params ListItemsConfig) toListConfig() itemProvider.ListConfig {
	return itemProvider.ListConfig{
		Limit:        params.Limit,
		Offset:       params.Offset,
		Count:        params.Count,
		ID:           params.ID,
		CategoryName: params.CategoryName,
		ProductName:  params.ProductName,
		WithDetails:  params.WithDetails,
	}
}
