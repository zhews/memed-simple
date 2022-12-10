package dto

type CheckUsernameRequest struct {
	Username string `json:"username,omitempty"`
}

type CheckUsernameResponse struct {
	Valid bool `json:"valid,omitempty"`
}
