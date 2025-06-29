package items

import (
	contract "item-detail-api/src/core/contracts/items"
	usecase "item-detail-api/src/core/usecases/items"
	"item-detail-api/src/entrypoints/rest"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const (
	CodeInvalidRequest = "U0"
)

type GetItemHandler struct {
	rest.HandlerBase
	logger  *zap.Logger
	usecase usecase.GetItem
}

func NewGetItemHandler(logger *zap.Logger, usecase usecase.GetItem) GetItemHandler {
	return GetItemHandler{rest.HandlerBase{}, logger, usecase}
}

func NewWireGetItemHandler(logger *zap.Logger, usecase usecase.GetItemImpl) GetItemHandler {
	return GetItemHandler{rest.HandlerBase{}, logger, usecase}
}

func (h GetItemHandler) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		r, description, err := rest.BindRequest[contract.GetItemRequest](c, rest.Uri)
		if err != nil {
			c.JSON(http.StatusBadRequest, rest.Errf(CodeInvalidRequest, description))
			return
		}
		item, _, err := h.usecase.Execute(c, r.ToListConfig(c))
		if err != nil {
			var (
				status        int
				message, code string
			)
			switch err {
			case usecase.ErrItemIDRequired:
				status, code, message = http.StatusBadRequest, "err-item-id-required", "item ID is required"
			case usecase.ErrItemNotFound:
				status, code, message = http.StatusBadRequest, "err-item-not-found", "item not found"
			default:
				status, code, message = rest.DefaultErrResponse()
			}
			h.logger.Error("failed to get item", zap.Any("request", r), zap.Error(err))
			c.JSON(status, rest.Errf(code, message))
			return
		}

		c.JSON(http.StatusOK, rest.Okf(contract.ToResponse(item)))
	}
}
