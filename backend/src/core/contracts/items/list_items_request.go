package items

import (
	usecase "item-detail-api/src/core/usecases/items"

	"github.com/google/uuid"
)

type ListItemsRequest struct {
	Limit        int    `form:"limit"`
	Offset       int    `form:"offset"`
	Count        bool   `form:"count"`
	ID           string `form:"id"`
	CategoryName string `form:"category_name"`
	ProductName  string `form:"product_name"`
	WithDetails  bool   `form:"with_details"`
}

func (r ListItemsRequest) ToParams() usecase.ListItemsConfig {
	id, err := uuid.Parse(r.ID)
	if err != nil {
		id = uuid.Nil
	}
	return usecase.ListItemsConfig{
		Limit:        r.Limit,
		Offset:       r.Offset,
		Count:        r.Count,
		ID:           id,
		CategoryName: r.CategoryName,
		ProductName:  r.ProductName,
		WithDetails:  r.WithDetails,
	}
}
