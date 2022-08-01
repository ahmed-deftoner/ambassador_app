package controllers

import (
	"github.com/ahmed-deftoner/ambassador/pkg/database"
	"github.com/ahmed-deftoner/ambassador/pkg/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
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

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 4)

	user := models.User{
		FirstName:    data["firstName"],
		LastName:     data["lastName"],
		Email:        data["email"],
		Password:     string(password),
		IsAmbassador: false,
	}

	database.DB.Create(&user)

	return c.JSON(user)
}
