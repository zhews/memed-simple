package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func AdminOnly(ctx *fiber.Ctx) error {
	token, ok := ctx.Locals("user").(*jwt.Token)
	if !ok {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	claims := token.Claims.(jwt.MapClaims)
	admin := claims["admin"].(bool)
	if !admin {
		return ctx.SendStatus(fiber.StatusForbidden)
	}
	return ctx.Next()
}
