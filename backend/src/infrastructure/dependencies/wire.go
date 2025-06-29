//go:build wireinject
// +build wireinject

package dependencies

import (
	"item-detail-api/config"
	itemProvider "item-detail-api/src/core/providers/items"
	itemUsecase "item-detail-api/src/core/usecases/items"
	itemContainer "item-detail-api/src/entrypoints"
	itemHandler "item-detail-api/src/entrypoints/rest/items"
	itemRepository "item-detail-api/src/repositories/items"

	pingContainer "item-detail-api/src/entrypoints"
	pingHandler "item-detail-api/src/entrypoints/rest/ping"

	"github.com/google/wire"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func BuildWeb(conf *config.Config, logger *zap.Logger, db *gorm.DB) *WebHandlerContainer {
	wire.Build(
		handlersSet,
		containersSet,
		usecasesSet,
		repositoriesSet,
		wire.Struct(new(WebHandlerContainer), "*"),
	)
	return &WebHandlerContainer{}
}

var containersSet = wire.NewSet(
	itemContainer.NewItemContainer,
	pingContainer.NewPingContainer,
)

var handlersSet = wire.NewSet(
	itemHandler.NewWireGetItemHandler,
	pingHandler.NewWirePingHandler,
)

var usecasesSet = wire.NewSet(
	itemUsecase.NewGetItemImpl,
)

var repositoriesSet = wire.NewSet(
	itemRepository.NewItemRepository,
	wire.Bind(new(itemProvider.ItemsPersistor), new(*itemRepository.ItemRepository)),
)
