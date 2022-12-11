package dto

import "github.com/google/uuid"

type User struct {
	Id        uuid.UUID `json:"id,omitempty"`
	Username  string    `json:"username,omitempty" json:"username,omitempty"`
	Name      string    `json:"name,omitempty" json:"name,omitempty"`
	CreatedAt int64     `json:"createdAt,omitempty"`
	UpdatedAt int64     `json:"updatedAt,omitempty"`
}
