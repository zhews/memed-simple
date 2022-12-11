package handler

import "github.com/gofiber/fiber/v2"

func HandleHealth(ctx *fiber.Ctx) error {
	return ctx.SendStatus(fiber.StatusNoContent)
}
