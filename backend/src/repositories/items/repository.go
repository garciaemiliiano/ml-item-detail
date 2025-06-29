package items

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ItemRepository struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewItemRepository(db *gorm.DB, logger *zap.Logger) *ItemRepository {
	return &ItemRepository{
		db:     db,
		logger: logger,
	}
}
