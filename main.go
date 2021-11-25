package main

import (
	"first-application/database"
	"first-application/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {

	database.Connect()

	app := fiber.New()

	routes.Setup(app)

	app.Listen(":8000")
	//go run github.com/oxequa/realize start
}
