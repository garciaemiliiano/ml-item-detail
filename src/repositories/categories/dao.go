package categories

import "item-detail-api/src/repositories/basemodel"

type CategoryDAO struct {
	basemodel.BaseModel
	Name string `gorm:"column:name"`
}
