package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Hash(c *fiber.Ctx) error {
	hash := c.Params("hash")

	return c.SendString(fmt.Sprintf("Hello, %s!", hash))
}
