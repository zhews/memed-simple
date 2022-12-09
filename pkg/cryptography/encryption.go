package cryptography

import (
	"crypto/aes"
	"crypto/cipher"
)

func Encrypt(key, data []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonce, err := generateRandomBytes(gcm.NonceSize())
	if err != nil {
		return nil, err
	}
	cipherText := gcm.Seal(nonce, nonce, data, nil)
	return cipherText, nil
}
