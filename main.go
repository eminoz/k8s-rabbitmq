package main

import (
	"fmt"

	"github.com/eminoz/logging/broker"
	"github.com/eminoz/logging/model"
	"github.com/gofiber/fiber/v2"
)

func main() {
	f := fiber.New()
	broker.Connect()
	f.Get("/", func(c *fiber.Ctx) error { return c.JSON("hello") })
	f.Post("/create", func(c *fiber.Ctx) error {
		user := new(model.User)
		c.BodyParser(user)
		b := broker.NewUserProducer()
		b.CreatedUser(*user)
		fmt.Print("user sent to rabbitmq")
		return c.JSON(user)
	})

	f.Listen(":3000")
}
