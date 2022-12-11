package cryptography

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"golang.org/x/crypto/argon2"
	"regexp"
	"strconv"
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

func HashPassword(password string, parameter Argon2IDParameter) (string, error) {
	salt, err := generateRandomBytes(parameter.SaltSize)
	if err != nil {
		return "", err
	}
	hash := argon2.IDKey([]byte(password), salt, parameter.Iterations, parameter.Memory, parameter.Threads, parameter.KeyLength)
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)
	encodedHash := encodeHash(argon2.Version, parameter, b64Salt, b64Hash)
	return encodedHash, nil
}

func encodeHash(version int, parameter Argon2IDParameter, salt, hash string) string {
	return fmt.Sprintf(hashFormat, version, parameter.Memory, parameter.Iterations, parameter.Threads, salt, hash)
}

var (
	RegexEncodedHash          = regexp.MustCompile(`\$argon2id\$v=(\d+)\$m=(\d+),t=(\d+),p=(\d+)\$([A-Za-z0-9+/=]+)\$([A-Za-z0-9+/=]+)`)
	ErrorInvalidHashEncoding  = errors.New("hash does not have a valid encoding")
	ErrorInvalidArgon2Version = errors.New("invalid argon2 version")
)

func decodeHash(encodedHash string) (Argon2IDParameter, []byte, []byte, error) {
	matches := RegexEncodedHash.FindAllStringSubmatch(encodedHash, -1)
	if len(matches) != 1 {
		return Argon2IDParameter{}, nil, nil, ErrorInvalidHashEncoding
	}
	version, err := strconv.Atoi(matches[0][1])
	if err != nil {
		return Argon2IDParameter{}, nil, nil, err
	}
	if argon2.Version != version {
		return Argon2IDParameter{}, nil, nil, ErrorInvalidArgon2Version
	}
	memory, err := strconv.Atoi(matches[0][2])
	if err != nil {
		return Argon2IDParameter{}, nil, nil, err
	}
	iterations, err := strconv.Atoi(matches[0][3])
	if err != nil {
		return Argon2IDParameter{}, nil, nil, err
	}
	threads, err := strconv.Atoi(matches[0][4])
	if err != nil {
		return Argon2IDParameter{}, nil, nil, err
	}
	salt, err := base64.RawStdEncoding.DecodeString(matches[0][5])
	if err != nil {
		return Argon2IDParameter{}, nil, nil, err
	}
	hash, err := base64.RawStdEncoding.DecodeString(matches[0][6])
	if err != nil {
		return Argon2IDParameter{}, nil, nil, err
	}
	argon2IDParameter := Argon2IDParameter{
		SaltSize:   len(salt),
		Iterations: uint32(iterations),
		Memory:     uint32(memory),
		Threads:    uint8(threads),
		KeyLength:  uint32(len(hash)),
	}
	return argon2IDParameter, salt, hash, nil
}

var ErrorHashAndPasswordDoNotMatch = errors.New("hash and password do not match")

func CompareHashAndPassword(encodedHash, password string) error {
	argon2IDParameter, salt, hash, err := decodeHash(encodedHash)
	if err != nil {
		return err
	}
	hashedPassword := argon2.IDKey([]byte(password), salt, argon2IDParameter.Iterations, argon2IDParameter.Memory, argon2IDParameter.Threads, argon2IDParameter.KeyLength)
	if bytes.Compare(hash, hashedPassword) != 0 {
		return ErrorHashAndPasswordDoNotMatch
	}
	return nil
}
