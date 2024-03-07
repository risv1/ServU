package routes

import (
	"ServU/src/controllers"

	"github.com/gofiber/fiber/v2"
)


func UserRoutes(app *fiber.App) {

	app.Get("/session", controllers.GetUserSessionHandler)

	app.Post("/register", controllers.Register)
	app.Post("/login", controllers.Login)
	app.Post("/logout", controllers.Logout)
}