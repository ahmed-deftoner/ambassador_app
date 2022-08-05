package controllers

import (
	"fmt"

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

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User
	//var x int64
	database.DB.First(&user)
	fmt.Println(user)

	if user.Id == 0 {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "invalid password",
		})
	}

	return c.JSON(user)
}
