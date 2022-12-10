package dto

type LoginRequest struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type LoginResponse struct {
	AccessToken string `json:"accessToken,omitempty"`
}
