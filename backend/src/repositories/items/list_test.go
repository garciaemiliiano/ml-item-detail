package items

import (
	"context"
	itemProvider "item-detail-api/src/core/providers/items"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestItemRepository_List_OK(t *testing.T) {
	defer suite.Clear()
	assert := assert.New(t)
	ctx := context.Background()

	itemID := uuid.New()
	itemID2 := uuid.New()

	populate(t, itemID, itemID2)
	r := NewItemRepository(suite.DB, nil)

	_, count, err := r.List(ctx, itemProvider.ListConfig{
		Limit: 10,
		Count: true,
	})

	assert.NoError(err)
	assert.Equal(int(2), count)

	_, count, err = r.List(ctx, itemProvider.ListConfig{
		Limit: 101,
		Count: true,
	})

	assert.NoError(err)
	assert.Equal(int(2), count)

	_, count, err = r.List(ctx, itemProvider.ListConfig{
		Count:        true,
		WithDetails:  true,
		CategoryName: "Hogar",
	})

	assert.NoError(err)
}

func TestObjectAcquisitionRepository_Error(t *testing.T) {
	defer suite.Clear()

	assert := assert.New(t)
	ctx := context.Background()
	r := NewItemRepository(suite.DB.Delete(ItemDAO{}), nil)
	_, _, err := r.List(ctx, itemProvider.ListConfig{
		Limit:  1,
		Offset: 0,
		Count:  true,
	})
	assert.Error(err)

	_, _, err = r.List(ctx, itemProvider.ListConfig{
		Limit:  1,
		Offset: 0,
	})
	assert.Error(err)
}
