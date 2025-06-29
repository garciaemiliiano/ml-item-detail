package dependencies

import (
	"item-detail-api/config"
	"item-detail-api/src/entrypoints"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type WebHandlerContainer struct {
	ItemContainer entrypoints.ItemContainer
	PingContainer entrypoints.PingContainer
}

func Build(conf *config.Config, logger *zap.Logger, db *gorm.DB) *WebHandlerContainer {
	return BuildWeb(conf, logger, db)
}
