package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/insyri/pfs/backend/util"
)

func Hash(c *fiber.Ctx) error {
	util.LogAPIConn(c)
	hash := c.Params("hash")

	return c.SendString(fmt.Sprintf("Hello, %s!", hash))
}
