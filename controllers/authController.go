package controllers

import (
	"first-application/models"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	user := models.User{
		FirstName: "Jhon",
	}

	user.LastName = "Doe"
	user.Email = "dicky@gmail.com"

	return c.JSON(user)
}
