package routes

import (
	"first-application/controllers"
	"first-application/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)

	app.Use(middlewares.IsAuthenticated)

	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)

	app.Get("/api/user/list", controllers.AllUsers)
	app.Post("/api/user/create", controllers.CreateUser)
	app.Put("/api/user/update/:id", controllers.UpdateUser)
	app.Get("/api/user/by/:id", controllers.GetUser)
	app.Delete("/api/user/delete/:id", controllers.Deleteuser)

	app.Get("/api/role/list", controllers.AllRoles)
	app.Post("/api/role/create", controllers.CreateRole)
	app.Put("/api/role/update/:id", controllers.UpdateRole)
	app.Get("/api/role/by/:id", controllers.GetRole)
	app.Delete("/api/role/delete/:id", controllers.DeleteRole)

	app.Get("/api/permission/list", controllers.AllPermissions)
}
