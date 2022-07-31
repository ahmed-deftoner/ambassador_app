package main

import (
	"github.com/ahmed-deftoner/ambassador/pkg/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()
	database.AddMigration()
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Listen(":8000")
}
