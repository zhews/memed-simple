package config

import "github.com/zhews/memed-simple/pkg/cryptography"

type UserConfig struct {
	BaseURI           string                         `json:"baseURI,omitempty"`
	AccessSecretKey   string                         `json:"accessSecretKey,omitempty" json:"accessSecretKey,omitempty"`
	RefreshSecretKey  string                         `json:"refreshSecretKey,omitempty" json:"refreshSecretKey,omitempty"`
	EncryptionKey     string                         `json:"encryptionKey,omitempty" json:"encryptionKey,omitempty"`
	Argon2IDParameter cryptography.Argon2IDParameter `json:"argon2IDParameter,omitempty" json:"argon2IDParameter"`
}
