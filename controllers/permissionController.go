package controllers

import (
	"first-application/database"
	"first-application/models"

	"github.com/gofiber/fiber/v2"
)

func AllPermissions(c *fiber.Ctx) error {
	var permission []models.Permission

	database.DB.Find(&permission)

	return c.JSON(permission)
}
