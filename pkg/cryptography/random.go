package cryptography

import "crypto/rand"

func generateRandomBytes(bytes int) ([]byte, error) {
	out := make([]byte, bytes)
	_, err := rand.Read(out)
	return out, err
}
