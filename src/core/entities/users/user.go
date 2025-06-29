package users

import "github.com/google/uuid"

type User struct {
	ID        uuid.UUID
	CreatedAt string
	UpdatedAt string
	DeletedAt *string
	Username  string
}
