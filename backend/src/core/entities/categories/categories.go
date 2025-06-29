package categories

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Name      string
}
