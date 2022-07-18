package util

import (
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/gofiber/fiber/v2"
)

var (
	Verbose = log.New(os.Stderr, color.New(color.FgYellow).Sprint("Verbose")+" ", log.Lmsgprefix|log.Lshortfile)
	Info    = log.New(os.Stderr, color.New(color.FgBlue).Sprint("Info")+" ", log.Lmsgprefix)
	Success = log.New(os.Stderr, color.New(color.FgGreen).Sprint("Success")+" ", log.Lmsgprefix)
	Panic   = log.New(os.Stderr, color.New(color.BgHiRed|color.Bold).Sprint("Panic")+" ", log.Lmsgprefix|log.Llongfile)
	Error   = log.New(os.Stderr, color.New(color.FgRed).Sprint("Error")+" ", log.Lmsgprefix|log.Llongfile)
)

func LogAPIConn(c *fiber.Ctx) {
	Info.Printf("%s => %s from User-Agent %s\n", c.Method(), c.Path(), c.Context().UserAgent())
}
