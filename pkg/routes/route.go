package routes

import (
	"github.com/ahmed-deftoner/ambassador/pkg/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app fiber.App) {
	api := app.Group("api")
	admin := api.Group("admin")
	admin.Post("register", controllers.Register)
}
