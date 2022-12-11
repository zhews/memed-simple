package domain

import "github.com/google/uuid"

type Meme struct {
	Title     string
	Image     string
	CreatedBy uuid.UUID
	CreatedAt int64
	UpdatedAt int64
}
