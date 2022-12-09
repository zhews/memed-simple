package domain

import "github.com/google/uuid"

type User struct {
	Id           uuid.UUID
	Username     string
	Name         string
	PasswordHash string
}
