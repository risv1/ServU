package main

import (
	db "ServU/src/database"
	"ServU/src/middlewares"
	"ServU/src/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {

	db.Connect()

	app := fiber.New()

	app.Use(middlewares.CorsConfig)

	routes.UserRoutes(app)
	routes.AdminRoutes(app)

	app.Listen(":8000")
}
