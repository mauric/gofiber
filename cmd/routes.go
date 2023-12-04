package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mauric/gofiber/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.ListFacts)

	app.Post("/facts", handlers.CreateFact)
}
