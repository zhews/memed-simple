package domain

import "github.com/google/uuid"

type User struct {
	Id           uuid.UUID
	Username     string
	Name         string
	Admin        bool
	PasswordHash []byte
	CreatedAt    int64
	UpdatedAt    int64
}

type UserRepository interface {
	GetByUsername(username string) (User, error)
	Insert(user User) error
}
