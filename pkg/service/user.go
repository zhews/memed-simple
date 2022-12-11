package service

import (
	"errors"
	"github.com/google/uuid"
	userConfig "github.com/zhews/memed-simple/pkg/config/user"
	"github.com/zhews/memed-simple/pkg/cryptography"
	"github.com/zhews/memed-simple/pkg/domain"
	"github.com/zhews/memed-simple/pkg/repository"
	"time"
)

type UserService struct {
	Config     userConfig.Config
	Repository domain.UserRepository
}

func (us *UserService) Register(username, name, password string) error {
	passwordHash, err := cryptography.HashPassword(password, us.Config.Argon2IDParameter)
	if err != nil {
		return err
	}
	encryptedPasswordHash, err := cryptography.Encrypt([]byte(us.Config.EncryptionKey), []byte(passwordHash))
	if err != nil {
		return err
	}
	now := time.Now().Unix()
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

func (us *UserService) Login(username, password string) (domain.User, error) {
	user, err := us.Repository.GetByUsername(username)
	if err != nil {
		if errors.Is(err, repository.ErrorNoRows) {
			return domain.User{}, ErrorUserNotFound
		}
		return domain.User{}, err
	}
	correctPasswordHash, err := cryptography.Decrypt([]byte(us.Config.EncryptionKey), user.PasswordHash)
	if err != nil {
		return domain.User{}, err
	}
	err = cryptography.CompareHashAndPassword(correctPasswordHash, password)
	if err != nil {
		return domain.User{}, ErrorInvalidCredentials
	}
	return user, nil
}

func (us *UserService) CheckUsername(username string) (bool, error) {
	_, err := us.Repository.GetByUsername(username)
	if err != nil && errors.Is(err, repository.ErrorNoRows) {
		return true, nil
	}
	return false, err
}
