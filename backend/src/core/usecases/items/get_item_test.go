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

func TestGetItemImpl_Execute_OK(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	mockLogger := zaptest.NewLogger(t)

	itemPersistor := mocks.NewItemsPersistor(t)
	itemPersistor.On("List", ctx, mock.Anything).Return([]items.Item{{
		ID:           uuid.New(),
		ProductID:    uuid.New(),
		ProviderID:   uuid.New(),
		SerialNumber: "SN123456",
		Stock:        10,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		DeletedAt:    nil,
	}}, 1, nil)

	useCase := NewGetItemImpl(mockLogger, itemPersistor)

	_, err := useCase.Execute(ctx, GetItemParams{ID: uuid.New()})
	assert.Nil(err)
}

func TestGetItemImpl_Execute_Error(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()
	mockLogger := zaptest.NewLogger(t)

	itemPersistor := mocks.NewItemsPersistor(t)

	useCase := NewGetItemImpl(mockLogger, itemPersistor)

	_, err := useCase.Execute(ctx, GetItemParams{ID: uuid.Nil})
	assert.Error(err)

	itemPersistor.On("List", ctx, mock.Anything).Return([]items.Item{}, 0, errors.New("repo error")).Once()

	_, err = useCase.Execute(ctx, GetItemParams{ID: uuid.New()})
	assert.Error(err)

	itemPersistor.On("List", ctx, mock.Anything).Return([]items.Item{}, 0, nil).Once()
	_, err = useCase.Execute(ctx, GetItemParams{ID: uuid.New()})
	assert.Error(err)
}
