package items

import "item-detail-api/src/core/entities/items"

type ListItemsResponses struct {
	Items []GetItemResponse `json:"items"`
	Count int               `json:"count"`
}

func ToResponses(items []items.Item, count int) ListItemsResponses {
	itemResponses := make([]GetItemResponse, 0)
	for _, r := range items {
		itemResponses = append(itemResponses, ToResponse(r))
	}
	return ListItemsResponses{
		Items: itemResponses,
		Count: count,
	}

}
