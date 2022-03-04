package controllers

import (
	"first-application/database"
	"first-application/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func AllUsers(c *fiber.Ctx) error {
	var users []models.User

	database.DB.Preload("Role").Find(&users)

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

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")

	user := models.User{
		Id: id,
	}

	database.DB.Preload("Role").Find(&user)

	if user.FirstName == "" {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "Data User Not Found",
		})
	}

	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")

	user := models.User{
		Id: id,
	}

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	database.DB.Model(&user).Updates(user)

	return c.JSON(user)
}

func Deleteuser(c *fiber.Ctx) error {
	id := c.Params("id")

	user := models.User{
		Id: id,
	}

	database.DB.Delete(user)

	return c.JSON(fiber.Map{
		"message": "Success Delete",
	})
}
