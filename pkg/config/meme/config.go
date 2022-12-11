package meme

type Config struct {
	Port             int    `json:"port,omitempty"`
	DatabaseURL      string `json:"databaseURL,omitempty"`
	MemeDirectory    string `json:"memeDirectory,omitempty"`
	UserMicroservice string `json:"userMicroservice,omitempty"`
	UserEndpoint     string `json:"userEndpoint,omitempty"`
	AccessSecretKey  string `json:"accessSecretKey,omitempty"`
}
