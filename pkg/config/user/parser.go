package user

import (
	"github.com/zhews/memed-simple/pkg/cryptography"
	"os"
	"strconv"
)

const (
	VariablePort                        = "MEMED_PORT"
	VariableCorsAllowOrigins            = "MEMED_CORS_ALLOW_ORIGINS"
	VariableDatabaseURL                 = "MEMED_DATABASE_URL"
	VariableBaseURI                     = "MEMED_BASE_URI"
	VariableAccessSecretKey             = "MEMED_ACCESS_SECRET_KEY"
	VariableAccessTokenValidSeconds     = "MEMED_ACCESS_TOKEN_VALID_MINUTES"
	VariableRefreshSecretKey            = "MEMED_REFRESH_SECRET_KEY"
	VariableRefreshTokenValidHours      = "MEMED_REFRESH_TOKEN_VALID_HOURS"
	VariableEncryptionKey               = "MEMED_ENCRYPTION_KEY"
	VariableArgon2IDParameterSaltSize   = "MEMED_ARGON2ID_PARAMETER_SALT_SIZE"
	VariableArgon2IDParameterIterations = "MEMED_ARGON2ID_PARAMETER_ITERATIONS"
	VariableArgon2IDParameterMemory     = "MEMED_ARGON2ID_PARAMETER_MEMORY"
	VariableArgon2IDParameterThreads    = "MEMED_ARGON2ID_PARAMETER_THREADS"
	VariableArgon2IDParameterKeyLength  = "MEMED_ARGON2ID_PARAMETER_KEY_LENGTH"
)

func FromEnvironmentalVariables() (Config, error) {
	portString := os.Getenv(VariablePort)
	port, err := strconv.Atoi(portString)
	if err != nil {
		return Config{}, err
	}
	corsAllowOrigins := os.Getenv(VariableCorsAllowOrigins)
	databaseURL := os.Getenv(VariableDatabaseURL)
	baseURI := os.Getenv(VariableBaseURI)
	accessSecretKey := os.Getenv(VariableAccessSecretKey)
	accessTokenValidSecondsString := os.Getenv(VariableAccessTokenValidSeconds)
	accessTokenValidSeconds, err := strconv.Atoi(accessTokenValidSecondsString)
	if err != nil {
		return Config{}, err
	}
	refreshSecretKey := os.Getenv(VariableRefreshSecretKey)
	refreshTokenValidHoursString := os.Getenv(VariableRefreshTokenValidHours)
	refreshTokenValidHours, err := strconv.Atoi(refreshTokenValidHoursString)
	if err != nil {
		return Config{}, err
	}
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
		Port:                    port,
		CorsAllowOrigins:        corsAllowOrigins,
		DatabaseURL:             databaseURL,
		BaseURI:                 baseURI,
		AccessSecretKey:         accessSecretKey,
		AccessTokenValidSeconds: accessTokenValidSeconds,
		RefreshSecretKey:        refreshSecretKey,
		RefreshTokenValidHours:  refreshTokenValidHours,
		EncryptionKey:           encryptionKey,
		Argon2IDParameter:       argon2IDParameter,
	}
	return config, nil
}
