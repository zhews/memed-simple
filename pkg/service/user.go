package service

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/zhews/memed-simple/pkg/cryptography"
	"github.com/zhews/memed-simple/pkg/domain"
	"time"
)

type UserService struct {
	Argon2IDParameters cryptography.Argon2IDParameters
	Repository         domain.UserRepository
}

var ErrorInvalidCredentials = errors.New("invalid credentials")

func (us *UserService) Register(username, name, password string) error {
	now := time.Now().Unix()
	passwordHash, err := cryptography.HashPassword(password, us.Argon2IDParameters)
	if err != nil {
		return err
	}
	encryptedPasswordHash, err := cryptography.Encrypt([]byte{}, []byte(passwordHash))
	if err != nil {
		return err
	}
	user := domain.User{
		Id:           uuid.New(),
		Username:     username,
		Name:         name,
		Admin:        false,
		PasswordHash: encryptedPasswordHash,
		CreatedAt:    now,
		UpdatedAt:    now,
	}
	err = us.Repository.Insert(user)
	return err
}

func (us *UserService) Login(username, password string) (string, string, error) {
	user, err := us.Repository.GetByUsername(username)
	if err != nil {
		return "", "", err
	}
	correctPasswordHash, err := cryptography.Decrypt([]byte{}, user.PasswordHash)
	if err != nil {
		return "", "", err
	}
	passwordHash, err := cryptography.HashPassword(password, us.Argon2IDParameters)
	if correctPasswordHash != passwordHash {
		return "", "", ErrorInvalidCredentials
	}
	claims := jwt.MapClaims{
		"iss":   "api.memed.io/user",
		"sub":   user.Id.String(),
		"exp":   time.Now().Add(time.Second * 10).Unix(),
		"iat":   time.Now().Unix(),
		"admin": user.Admin,
	}
	accessToken, err := cryptography.CreateJWT([]byte{}, claims)
	if err != nil {
		return "", "", err
	}
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	refreshToken, err := cryptography.CreateJWT([]byte{}, claims)
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}
