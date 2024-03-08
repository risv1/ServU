package middlewares

import (
	"ServU/src/controllers"
	"github.com/gofiber/fiber/v2"
)

func IsAdmin (c *fiber.Ctx) error {
	user, err := controllers.GetUserSession(c)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	if user["role"] != "admin" {
		return c.Status(403).JSON(fiber.Map{
			"message": "Forbidden",
		})
	}

	return c.Next()
}