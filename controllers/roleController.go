package controllers

import (
	"first-application/database"
	"first-application/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func AllRoles(c *fiber.Ctx) error {
	var roles []models.Role

	database.DB.Find(&roles)

	return c.JSON(roles)
}

func CreateRole(c *fiber.Ctx) error {
	var role models.Role

	if err := c.BodyParser(&role); err != nil {
		return err
	}

	role.Id = uuid.New().String()

	database.DB.Create(&role)

	return c.JSON(role)
}

func GetRole(c *fiber.Ctx) error {
	id := c.Params("id")

	role := models.Role{
		Id: id,
	}

	database.DB.Find(&role)

	if role.Name == "" {
		c.Status(404)
		return c.JSON(fiber.Map{
			"message": "Data Role Not Found",
		})
	}

	return c.JSON(role)
}

func UpdateRole(c *fiber.Ctx) error {
	id := c.Params("id")

	role := models.Role{
		Id: id,
	}

	if err := c.BodyParser(&role); err != nil {
		return err
	}

	database.DB.Model(&role).Updates(role)

	return c.JSON(role)
}

func DeleteRole(c *fiber.Ctx) error {
	id := c.Params("id")

	role := models.Role{
		Id: id,
	}

	database.DB.Delete(role)

	return c.JSON(fiber.Map{
		"message": "Success Delete",
	})
}
