package handler

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	userConfig "github.com/zhews/memed-simple/pkg/config/user"
	"github.com/zhews/memed-simple/pkg/cryptography"
	"github.com/zhews/memed-simple/pkg/handler/dto"
	"github.com/zhews/memed-simple/pkg/repository"
	"github.com/zhews/memed-simple/pkg/service"
	"time"
)

type UserHandler struct {
	Config  userConfig.Config
	Service service.UserService
}

func (uh *UserHandler) HandleGetById(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	user, err := uh.Service.GetById(id)
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	response := dto.UserResponse{
		Id:        user.Id,
		Username:  user.Username,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (uh *UserHandler) HandleRegister(ctx *fiber.Ctx) error {
	var request dto.RegisterRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	err := uh.Service.Register(request.Username, request.Name, request.Password)
	if err != nil {
		if errors.Is(err, repository.ErrorUsernameAlreadyTaken) {
			return ctx.SendStatus(fiber.StatusBadRequest)
		}
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	return ctx.SendStatus(fiber.StatusNoContent)
}

func (uh *UserHandler) HandleLogin(ctx *fiber.Ctx) error {
	var request dto.LoginRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.SendStatus(fiber.StatusBadRequest)
	}
	user, err := uh.Service.Login(request.Username, request.Password)
	if err != nil {
		if errors.Is(err, service.ErrorUserNotFound) || errors.Is(err, service.ErrorInvalidCredentials) {
			return ctx.SendStatus(fiber.StatusUnauthorized)
		}
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	claims := jwt.MapClaims{
		"iss":   fmt.Sprintf("%s/user", uh.Config.BaseURI),
		"sub":   user.Id.String(),
		"exp":   time.Now().Add(time.Second * time.Duration(uh.Config.AccessTokenValidSeconds)).Unix(),
		"iat":   time.Now().Unix(),
		"admin": user.Admin,
	}
	accessToken, err := cryptography.CreateJWT([]byte(uh.Config.AccessSecretKey), claims)
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(uh.Config.RefreshTokenValidHours)).Unix()
	refreshToken, err := cryptography.CreateJWT([]byte(uh.Config.RefreshSecretKey), claims)
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	response := dto.LoginResponse{
		AccessToken: accessToken,
	}
	cookie := &fiber.Cookie{
		Name:     "refreshToken",
		Value:    refreshToken,
		Secure:   true,
		HTTPOnly: true,
	}
	ctx.Cookie(cookie)
	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (uh *UserHandler) HandleCheckUsername(ctx *fiber.Ctx) error {
	username := ctx.Params("username")
	valid, err := uh.Service.CheckUsername(username)
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	response := dto.CheckUsernameResponse{
		Valid: valid,
	}
	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (uh *UserHandler) HandleRefresh(ctx *fiber.Ctx) error {
	refreshToken := ctx.Cookies("refreshToken")
	claims, err := cryptography.ValidateJWT([]byte(uh.Config.RefreshSecretKey), refreshToken)
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	renewedClaims := jwt.MapClaims{
		"iss":   fmt.Sprintf("%s/user", uh.Config.BaseURI),
		"sub":   claims["sub"],
		"exp":   time.Now().Add(time.Second * time.Duration(uh.Config.AccessTokenValidSeconds)).Unix(),
		"iat":   time.Now().Unix(),
		"admin": claims["admin"],
	}
	accessToken, err := cryptography.CreateJWT([]byte(uh.Config.AccessSecretKey), renewedClaims)
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	response := dto.RefreshResponse{
		AccessToken: accessToken,
	}
	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (uh *UserHandler) HandleLogout(ctx *fiber.Ctx) error {
	ctx.ClearCookie("refreshToken")
	return ctx.SendStatus(fiber.StatusNoContent)
}
