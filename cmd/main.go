// ⚡️ Fiber is an Express inspired web framework written in Go with ☕️
// 🤖 Github Repository: https://github.com/gofiber/fiber
// 📌 API Documentation: https://docs.gofiber.io

package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/mauric/gofiber/database"
)

func main() {
	database.ConnecDb()
	// Fiber instance
	app := fiber.New()

	setupRoutes(app)

	// Start server
	log.Fatal(app.Listen(":3000"))
}
