package user

import (
	"github.com/zhews/memed-simple/pkg/cryptography"
	"os"
	"strconv"
)

const (
	VariableBaseURI                     = "MEMED_BASE_URI"
	VariableAccessSecretKey             = "MEMED_ACCESS_SECRET_KEY"
	VariableRefreshSecretKey            = "MEMED_REFRESH_SECRET_KEY"
	VariableEncryptionKey               = "MEMED_ENCRYPTION_KEY"
	VariableArgon2IDParameterSaltSize   = "MEMED_ARGON2ID_PARAMETER_SALT_SIZE"
	VariableArgon2IDParameterIterations = "MEMED_ARGON2ID_PARAMETER_ITERATIONS"
	VariableArgon2IDParameterMemory     = "MEMED_ARGON2ID_PARAMETER_MEMORY"
	VariableArgon2IDParameterThreads    = "MEMED_ARGON2ID_PARAMETER_THREADS"
	VariableArgon2IDParameterKeyLength  = "MEMED_ARGON2ID_PARAMETER_KEY_LENGTH"
)

func FromEnvironmentalVariables() (Config, error) {
	baseURI := os.Getenv(VariableBaseURI)
	accessSecretKey := os.Getenv(VariableAccessSecretKey)
	refreshSecretKey := os.Getenv(VariableRefreshSecretKey)
	encryptionKey := os.Getenv(VariableEncryptionKey)
	argon2IDParameterSaltSizeString := os.Getenv(VariableArgon2IDParameterSaltSize)
	argon2IDParameterSaltSize, err := strconv.Atoi(argon2IDParameterSaltSizeString)
	if err != nil {
		return Config{}, err
	}
	argon2IDParameterIterationsString := os.Getenv(VariableArgon2IDParameterIterations)
	argon2IDParameterIterations, err := strconv.Atoi(argon2IDParameterIterationsString)
	if err != nil {
		return Config{}, err
	}
	argon2IDParameterMemoryString := os.Getenv(VariableArgon2IDParameterMemory)
	argon2IDParameterMemory, err := strconv.Atoi(argon2IDParameterMemoryString)
	if err != nil {
		return Config{}, err
	}
	argon2IDParameterThreadsString := os.Getenv(VariableArgon2IDParameterThreads)
	argon2IDParameterThreads, err := strconv.Atoi(argon2IDParameterThreadsString)
	if err != nil {
		return Config{}, err
	}
	argon2IDParameterKeyLengthString := os.Getenv(VariableArgon2IDParameterKeyLength)
	argon2IDParameterKeyLength, err := strconv.Atoi(argon2IDParameterKeyLengthString)
	if err != nil {
		return Config{}, err
	}
	argon2IDParameter := cryptography.Argon2IDParameter{
		SaltSize:   argon2IDParameterSaltSize,
		Iterations: uint32(argon2IDParameterIterations),
		Memory:     uint32(argon2IDParameterMemory),
		Threads:    uint8(argon2IDParameterThreads),
		KeyLength:  uint32(argon2IDParameterKeyLength),
	}
	config := Config{
		BaseURI:           baseURI,
		AccessSecretKey:   accessSecretKey,
		RefreshSecretKey:  refreshSecretKey,
		EncryptionKey:     encryptionKey,
		Argon2IDParameter: argon2IDParameter,
	}
	return config, nil
}
