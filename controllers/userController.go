package controllers

import (
	"first-application/database"
	"first-application/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func AllUsers(c *fiber.Ctx) error {
	var users []models.User

	database.DB.Find(&users)

	return c.JSON(users)
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User
	var data models.User

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	database.DB.Where("Email = ?", user.Email).First(&data)

	if data.Email != "" {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Email Already Exists",
		})
	}

	user.Id = uuid.New().String()
	user.SetPassword("1234")

	database.DB.Create(&user)

	return c.JSON(user)
}
