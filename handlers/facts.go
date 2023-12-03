package handlers

import "github.com/gofiber/fiber/v2"

func Home(c *fiber.Ctx) error {
	return c.SendString("hello, mauric here")
}

// Handler old version
func hello(c *fiber.Ctx) error {
	return c.SendString("Hello, Argentina 👋!")
}
