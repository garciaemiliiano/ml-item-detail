package entrypoints

import "item-detail-api/src/entrypoints/rest/items"

type ItemContainer struct {
	GetProduct items.GetItemHandler
	ListItems  items.ListItemsHandler
}

func NewItemContainer(GetProduct items.GetItemHandler, ListItems items.ListItemsHandler) ItemContainer {
	return ItemContainer{GetProduct, ListItems}
}
