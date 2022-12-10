package cryptography

import (
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/argon2"
)

const (
	hashFormat = "$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s"
)

type Argon2IDParameter struct {
	SaltSize   int
	Iterations uint32
	Memory     uint32
	Threads    uint8
	KeyLength  uint32
}

func HashPassword(password string, parameters Argon2IDParameter) (string, error) {
	salt, err := generateRandomBytes(parameters.SaltSize)
	if err != nil {
		return "", err
	}
	hash := argon2.IDKey([]byte(password), salt, parameters.Iterations, parameters.Memory, parameters.Threads, parameters.KeyLength)
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)
	encodedHash := encodeHash(argon2.Version, parameters, b64Salt, b64Hash)
	return encodedHash, nil
}

func encodeHash(version int, parameters Argon2IDParameter, salt, hash string) string {
	return fmt.Sprintf(hashFormat, version, parameters.Memory, parameters.Iterations, parameters.Threads, salt, hash)
}
