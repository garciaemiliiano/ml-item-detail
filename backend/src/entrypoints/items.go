package entrypoints

import "item-detail-api/src/entrypoints/rest/items"

type ItemContainer struct {
	GetProduct items.GetItemHandler
}

func NewItemContainer(GetProduct items.GetItemHandler) ItemContainer {
	return ItemContainer{GetProduct}
}
