package web

import (
	"item-detail-api/config"
	"item-detail-api/src/infrastructure/dependencies"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func configureMappings(router *gin.Engine, handlers *dependencies.WebHandlerContainer, conf *config.Config, logger *zap.Logger, db *gorm.DB) {
	setRoutes(router, handlers, conf)

}

func setRoutes(router *gin.Engine, handlers *dependencies.WebHandlerContainer, conf *config.Config) {
	items := router.Group("/items")
	items.GET("", handlers.ItemContainer.ListItems.Handle())
	items.GET("/:id", handlers.ItemContainer.GetProduct.Handle())

	ping := router.Group("/ping")
	ping.GET("", handlers.PingContainer.Ping.Handle())
}
