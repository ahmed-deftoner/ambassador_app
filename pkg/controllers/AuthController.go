package controllers

import (
	"github.com/ahmed-deftoner/ambassador/pkg/models"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] != data["confirmPassword"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "password confirmation failed",
		})
	}

	user := models.User{
		FirstName:    data["firstName"],
		LastName:     data["lastName"],
		Email:        data["email"],
		IsAmbassador: false,
	}
	return c.JSON(fiber.Map{
		"message": "hello",
	})
}
