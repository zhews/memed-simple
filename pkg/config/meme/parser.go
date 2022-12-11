package meme

import (
	"os"
	"strconv"
)

const (
	VariablePort             = "MEMED_PORT"
	VariableDatabaseURL      = "MEMED_DATABASE_URL"
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
	databaseURL := os.Getenv(VariableDatabaseURL)
	memeDirectory := os.Getenv(VariableMemeDirectory)
	userMicroservice := os.Getenv(VariableUserMicroservice)
	userEndpoint := os.Getenv(VariableUserEndpoint)
	accessSecretKey := os.Getenv(VariableAccessSecretKey)
	config := Config{
		Port:             port,
		DatabaseURL:      databaseURL,
		MemeDirectory:    memeDirectory,
		UserMicroservice: userMicroservice,
		UserEndpoint:     userEndpoint,
		AccessSecretKey:  accessSecretKey,
	}
	return config, err
}
