package items

import (
	"context"
	entity "item-detail-api/src/core/entities/items"
	provider "item-detail-api/src/core/providers/items"

	"github.com/google/uuid"
)

const (
	defaultLimit = 50
	maxLimit     = 100
)

func (r ItemRepository) List(ctx context.Context, config provider.ListConfig) ([]entity.Item, int, error) {
	var daos []ItemDAO
	count := int64(0)

	if config.Limit == 0 {
		config.Limit = defaultLimit
	}
	if config.Limit > maxLimit {
		config.Limit = maxLimit
	}

	q := r.db.WithContext(ctx).Model(&ItemDAO{})

	if config.ID != uuid.Nil {
		q = q.Where("id = ?", config.ID)
	}

	if config.Count {
		if err := q.Count(&count).Error; err != nil {
			return []entity.Item{}, 0, err
		}
	}

	if config.WithDetails {
		q = q.Preload("Product.Category").Preload("Provider")
	}

	q = q.Limit(config.Limit).Offset(config.Offset)
	if err := q.Debug().Find(&daos).Error; err != nil {
		return []entity.Item{}, 0, err
	}

	return toEntities(daos), int(count), nil
}
