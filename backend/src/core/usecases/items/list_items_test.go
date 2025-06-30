package items

import (
	"context"
	"errors"
	"item-detail-api/src/core/entities/items"
	"item-detail-api/src/core/providers/items/mocks"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap/zaptest"
)

func TestListItemsImpl_Execute_OK(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	mockLogger := zaptest.NewLogger(t)

	itemPersistor := mocks.NewItemsPersistor(t)
	id := uuid.New()
	itemPersistor.On("List", ctx, mock.Anything).Return([]items.Item{{
		ID:           id,
		ProductID:    uuid.New(),
		ProviderID:   uuid.New(),
		SerialNumber: "SN123456",
		Stock:        10,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		DeletedAt:    nil,
	}}, 1, nil)

	useCase := NewListItemsImpl(mockLogger, itemPersistor)

	items, _, err := useCase.Execute(ctx, ListItemsConfig{
		ID:           uuid.New(),
		Limit:        10,
		Offset:       0,
		Count:        true,
		WithDetails:  true,
		CategoryName: "Electronics",
		ProductName:  "Smartphone",
	})

	assert.Nil(err)
	assert.Equal(1, len(items))
	assert.Equal(id, items[0].ID)
}

func TestListItemsImpl_Execute_Error(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	mockLogger := zaptest.NewLogger(t)

	itemPersistor := mocks.NewItemsPersistor(t)
	itemPersistor.On("List", ctx, mock.Anything).Return([]items.Item{}, 1, errors.New("error listing items"))

	useCase := NewListItemsImpl(mockLogger, itemPersistor)

	_, _, err := useCase.Execute(ctx, ListItemsConfig{ID: uuid.Nil})
	assert.Error(err)

}
