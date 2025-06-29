package providers

import (
	"item-detail-api/src/core/entities/providers"
	"item-detail-api/src/core/utils"
	"item-detail-api/src/repositories/basemodel"
)

type ProviderDAO struct {
	basemodel.BaseModel
	Name string `gorm:"column:name"`
}

func (ProviderDAO) TableName() string {
	return "providers"
}

func (p ProviderDAO) ToEntity() providers.Provider {
	return providers.Provider{
		ID:        p.ID,
		Name:      p.Name,
		CreatedAt: utils.ParseDateStr(p.CreatedAt),
		UpdatedAt: utils.ParseDateStr(p.UpdatedAt),
	}
}
