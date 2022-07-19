package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Base(c *fiber.Ctx) error {
	return c.SendString("ok its up")
}
