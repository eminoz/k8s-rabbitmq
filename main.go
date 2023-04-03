package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	f := fiber.New()
	f.Get("/", func(c *fiber.Ctx) error { return c.JSON("hello") })

	f.Listen(":3000")
}
