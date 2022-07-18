package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/insyri/pfs/backend/util"
)

func Base(c *fiber.Ctx) error {
	util.LogAPIConn(c)
	return c.SendString("ok its up")
}
