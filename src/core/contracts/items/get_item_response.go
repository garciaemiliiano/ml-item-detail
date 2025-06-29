package items

import (
	entity "item-detail-api/src/core/entities/items"
)

type GetItemResponse struct {
	ID string `json:"id"`
}

func ToResponse(i entity.Item) GetItemResponse {
	return GetItemResponse{
		ID: i.ID.String(),
	}
}
