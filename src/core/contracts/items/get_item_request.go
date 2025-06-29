package items

import (
	usecase "item-detail-api/src/core/usecases/items"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type GetItemRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

func (lr GetItemRequest) ToParams(ctx *gin.Context) usecase.GetItemParams {
	id, err := uuid.Parse(lr.ID)
	if err != nil {
		id = uuid.Nil
	}
	return usecase.GetItemParams{
		ID: id,
	}
}
