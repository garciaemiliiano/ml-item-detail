package providers

import (
	"time"

	"github.com/google/uuid"
)

type Provider struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Name      string
}
