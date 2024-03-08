package controllers

import (
	db "ServU/src/database"
	"ServU/src/models"
	"github.com/gofiber/fiber/v2"
)

func GetRequests (c *fiber.Ctx) error {
	
	var requests []models.Request
	db.DB.Find(&requests)
	if len(requests) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "No users found",
		})
	}
	return c.JSON(requests)
		
}

func GetRequest (c *fiber.Ctx) error {
	var request models.Request
	requestId := c.Params("id")

	if err := db.DB.Where("id = ?", requestId).First(&request).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Request not found",
		})
	}

	return c.JSON(request)
}

func GetUsers (c *fiber.Ctx) error {
	var users []models.User
	db.DB.Find(&users)
	if len(users) == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "No users found",
		})
	}
	return c.JSON(users)
}

func GetUser (c *fiber.Ctx) error {
	var user models.User
	userId := c.Params("id")

	if err := db.DB.Where("id = ?", userId).First(&user).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Request not found",
		})
	}

	return c.JSON(user)
}
