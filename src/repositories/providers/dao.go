package providers

import "item-detail-api/src/repositories/basemodel"

type ProviderDAO struct {
	basemodel.BaseModel
	Name string `gorm:"column:name"`
}
