package domain

import "github.com/google/uuid"

type Meme struct {
	Id        uuid.UUID
	Title     string
	Image     string
	CreatedBy uuid.UUID
	CreatedAt int64
	UpdatedAt int64
}

type MemeRepository interface {
	GetAll() ([]Meme, error)
	GetById(id uuid.UUID) (Meme, error)
	Insert(meme Meme) error
	Update(meme Meme) error
	Delete(id uuid.UUID) error
}
