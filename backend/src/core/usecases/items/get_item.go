package items

import (
	"context"

	itemEntity "item-detail-api/src/core/entities/items"
	itemProvider "item-detail-api/src/core/providers/items"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type GetItemParams struct {
	ID uuid.UUID
}

//go:generate mockery --all
type GetItem interface {
	Execute(ctx context.Context, config GetItemParams) (itemEntity.Item, error)
}
type GetItemImpl struct {
	logger        *zap.Logger
	itemPersistor itemProvider.ItemsPersistor
}

func NewGetItemImpl(
	logger *zap.Logger,
	itemPersistor itemProvider.ItemsPersistor) GetItemImpl {
	return GetItemImpl{
		logger:        logger,
		itemPersistor: itemPersistor,
	}
}

func (uc GetItemImpl) Execute(ctx context.Context, params GetItemParams) (itemEntity.Item, error) {
	if err := params.validate(); err != nil {
		return itemEntity.Item{}, err
	}
	item, _, err := uc.itemPersistor.List(ctx, itemProvider.ListConfig{
		ID:          params.ID,
		WithDetails: true,
	})
	if err != nil {
		return itemEntity.Item{}, err
	}
	if len(item) == 0 {
		return itemEntity.Item{}, ErrItemNotFound
	}

	return item[0], nil
}

func (params GetItemParams) validate() error {
	if params.ID == uuid.Nil {
		return ErrItemIDRequired
	}
	return nil
}
