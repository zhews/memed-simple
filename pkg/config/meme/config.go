package meme

type Config struct {
	MemeDirectory    string `json:"memeDirectory,omitempty"`
	UserMicroservice string
	UserEndpoint     string
}
