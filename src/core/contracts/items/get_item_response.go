package items

import (
	entity "item-detail-api/src/core/entities/items"
)

type GetItemResponse struct {
	ID           string  `json:"id"`
	ProductID    string  `json:"product_id"`
	ProviderID   string  `json:"provider_id"`
	SerialNumber string  `json:"serial_number"`
	Stock        int     `json:"stock"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
	DeletedAt    *string `json:"deleted_at"`
}

func ToResponse(i entity.Item) GetItemResponse {
	return GetItemResponse{
		ID:           i.ID.String(),
		ProductID:    i.ProductID.String(),
		ProviderID:   i.ProviderID.String(),
		SerialNumber: i.SerialNumber,
		Stock:        i.Stock,
		CreatedAt:    i.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:    i.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}
