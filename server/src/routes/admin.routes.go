package routes

import (
	"ServU/src/controllers"
	"ServU/src/middlewares"
	"github.com/gofiber/fiber/v2"
)

func AdminRoutes(app *fiber.App) {

	admin := app.Group("/admin").Use(middlewares.IsAdmin)

	admin.Get("/requests", controllers.GetRequests)
	admin.Get("/requests/:id", controllers.GetRequest)
	admin.Get("/users", controllers.GetUsers)
	admin.Get("/users/:id", controllers.GetUser)

}