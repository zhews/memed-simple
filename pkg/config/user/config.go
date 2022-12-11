package user

import "github.com/zhews/memed-simple/pkg/cryptography"

type Config struct {
	Port              int                            `json:"port,omitempty"`
	DatabaseURL       string                         `json:"databaseURL,omitempty"`
	BaseURI           string                         `json:"baseURI,omitempty"`
	AccessSecretKey   string                         `json:"accessSecretKey,omitempty"`
	RefreshSecretKey  string                         `json:"refreshSecretKey,omitempty"`
	EncryptionKey     string                         `json:"encryptionKey,omitempty"`
	Argon2IDParameter cryptography.Argon2IDParameter `json:"argon2IDParameter,omitempty"`
}
