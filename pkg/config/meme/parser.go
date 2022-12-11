package meme

import (
	"os"
	"strconv"
)

const (
	VariablePort             = "MEMED_PORT"
	VariableDatabaseURL      = "MEMED_DATABASE_URL"
	VariableCorsAllowOrigins = "MEMED_ALLOW_ORIGINS"
	VariableMemeDirectory    = "MEMED_MEME_DIRECTORY"
	VariableUserMicroservice = "MEMED_USER_MICROSERVICE"
	VariableUserEndpoint     = "MEMED_USER_ENDPOINT"
	VariableAccessSecretKey  = "MEMED_ACCESS_SECRET_KEY"
)

func ParseFromEnvironmentalVariables() (Config, error) {
	portString := os.Getenv(VariablePort)
	port, err := strconv.Atoi(portString)
	if err != nil {
		return Config{}, err
	}
	allowOrigins := os.Getenv(VariableCorsAllowOrigins)
	databaseURL := os.Getenv(VariableDatabaseURL)
	memeDirectory := os.Getenv(VariableMemeDirectory)
	if _, err = os.Stat(memeDirectory); os.IsNotExist(err) {
		return Config{}, err
	}
	userMicroservice := os.Getenv(VariableUserMicroservice)
	userEndpoint := os.Getenv(VariableUserEndpoint)
	accessSecretKey := os.Getenv(VariableAccessSecretKey)
	config := Config{
		Port:             port,
		DatabaseURL:      databaseURL,
		CorsAllowOrigins: allowOrigins,
		MemeDirectory:    memeDirectory,
		UserMicroservice: userMicroservice,
		UserEndpoint:     userEndpoint,
		AccessSecretKey:  accessSecretKey,
	}
	return config, err
}
