package main

import (
	"github.com/ahmed-deftoner/ambassador/pkg/database"
	"github.com/ahmed-deftoner/ambassador/pkg/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()
	database.AddMigration()
	app := fiber.New()
	routes.Setup(app)

	app.Listen(":8000")
}
