package items

import (
	"time"

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

func parseDateStr(dateStr string) time.Time {
	t, err := time.Parse("2006-01-02 15:04:05", dateStr)
	if err != nil {
		return time.Time{}
	}
	return t
}
