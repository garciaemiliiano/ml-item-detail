package reviews

import "github.com/google/uuid"

type Review struct {
	ID          uuid.UUID
	CreatedAt   string
	UpdatedAt   string
	DeletedAt   *string
	Description string
	Rating      int
	ProductID   uuid.UUID
	UserID      uuid.UUID
}
