package items

import (
	contract "item-detail-api/src/core/contracts/items"
	usecase "item-detail-api/src/core/usecases/items"
	"item-detail-api/src/entrypoints/rest"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ListItemsHandler struct {
	rest.HandlerBase
	logger  *zap.Logger
	usecase usecase.ListItems
}

func NewListItemsHandler(logger *zap.Logger, usecase usecase.ListItems) ListItemsHandler {
	return ListItemsHandler{rest.HandlerBase{}, logger, usecase}
}

func NewWireListItemsHandler(logger *zap.Logger, usecase usecase.ListItemsImpl) ListItemsHandler {
	return ListItemsHandler{rest.HandlerBase{}, logger, usecase}
}

func (h ListItemsHandler) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		r, description, err := rest.BindRequest[contract.ListItemsRequest](c, rest.Form)
		if err != nil {
			c.JSON(http.StatusBadRequest, rest.Errf(CodeInvalidRequest, description))
			return
		}
		items, count, err := h.usecase.Execute(c, r.ToParams())
		if err != nil {
			var (
				status        int
				message, code string
			)
			switch err {
			default:
				status, code, message = rest.DefaultErrResponse()
			}
			h.logger.Error("failed to list items", zap.Any("request", r), zap.Error(err))
			c.JSON(status, rest.Errf(code, message))
			return
		}

		c.JSON(http.StatusOK, rest.Okf(contract.ToResponses(items, count)))
	}
}
