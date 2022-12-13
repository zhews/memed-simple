package cryptography

import (
	"github.com/golang-jwt/jwt/v4"
)

func CreateJWT(key []byte, claims jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	signedToken, err := token.SignedString(key)
	return signedToken, err
}

func ValidateJWT(key []byte, signedToken string) (jwt.MapClaims, error) {
	var claims jwt.MapClaims
	_, err := jwt.ParseWithClaims(signedToken, &claims, getKeyFunc(key))
	if err != nil {
		return nil, err
	}
	return claims, nil
}

func getKeyFunc(key []byte) func(token *jwt.Token) (interface{}, error) {
	return func(token *jwt.Token) (interface{}, error) {
		return key, nil
	}
}
