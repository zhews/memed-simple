package user

import "github.com/zhews/memed-simple/pkg/cryptography"

type Config struct {
	Port                    int                            `json:"port,omitempty"`
	CorsAllowOrigins        string                         `json:"corsAllowOrigins,omitempty"`
	DatabaseURL             string                         `json:"databaseURL,omitempty"`
	BaseURI                 string                         `json:"baseURI,omitempty"`
	AccessSecretKey         string                         `json:"accessSecretKey,omitempty"`
	AccessTokenValidSeconds int                            `json:"accessTokenValidSeconds,omitempty"`
	RefreshSecretKey        string                         `json:"refreshSecretKey,omitempty"`
	RefreshTokenValidHours  int                            `json:"refreshTokenValidHours,omitempty"`
	EncryptionKey           string                         `json:"encryptionKey,omitempty"`
	Argon2IDParameter       cryptography.Argon2IDParameter `json:"argon2IDParameter,omitempty"`
}
