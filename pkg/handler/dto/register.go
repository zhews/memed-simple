package dto

type RegisterRequest struct {
	Username string `json:"username,omitempty"`
	Name     string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
}
