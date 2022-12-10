package cryptography

import (
	"errors"
	"github.com/golang-jwt/jwt"
)

var (
	ErrorInvalidToken = errors.New("token is invalid")
)

func CreateJWT(key []byte, claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodEdDSA, claims)
	signedToken, err := token.SignedString(key)
	return signedToken, err
}

func ValidateJWT(key []byte, signedToken string) (jwt.MapClaims, error) {
	var claims jwt.MapClaims
	token, err := jwt.ParseWithClaims(signedToken, &claims, getKeyFunc(key))
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, ErrorInvalidToken
	}
	return claims, nil
}

func getKeyFunc(key []byte) func(token *jwt.Token) (interface{}, error) {
	return func(token *jwt.Token) (interface{}, error) {
		return key, nil
	}
}
