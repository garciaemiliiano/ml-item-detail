package basemodel

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uuid.UUID `gorm:"type:char(36);primary_key;"`
	CreatedAt string    `gorm:"column:created_at"`
	UpdatedAt string    `gorm:"column:updated_at"`
	DeletedAt *string   `gorm:"column:deleted_at"`
}

func (b *BaseModel) BeforeCreate(tx *gorm.DB) error {
	if b.ID == uuid.Nil {
		b.ID = uuid.New()
	}
	return nil
}
