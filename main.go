package main

import (
	"github.com/ahmed-deftoner/ambassador/pkg/database"
	"github.com/ahmed-deftoner/ambassador/pkg/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()
	database.AddMigration()
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	routes.Setup(app)

	app.Listen(":8000")
}
