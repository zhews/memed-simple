package dto

import "github.com/google/uuid"

type MemeResponse struct {
	Id        uuid.UUID    `json:"id,omitempty"`
	Title     string       `json:"title,omitempty"`
	Image     string       `json:"image,omitempty"`
	Creator   UserResponse `json:"creator"`
	CreatedAt int64        `json:"createdAt,omitempty"`
	UpdatedAt int64        `json:"updatedAt,omitempty"`
}
