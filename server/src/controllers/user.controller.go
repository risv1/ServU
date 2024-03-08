package controllers

import (
	db "ServU/src/database"
	"ServU/src/models"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Printf("Error loading .env file")
	}
}

func GetUserSession(c *fiber.Ctx) (map[string]interface{}, error) {
	cookie := c.Cookies("token")

	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SECRET), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}
	return claims, nil
}

func GetUserSessionHandler(c *fiber.Ctx) error {
	userSession, err := GetUserSession(c)
	if err != nil {
		return err
	}

	return c.JSON(userSession)
}

func SendRequest(c *fiber.Ctx) error {

	var data map[string]string

	err := c.BodyParser(&data)
	if err != nil {
		return err
	}

	userSession, err := GetUserSession(c)
	if err != nil {
		return err
	}

	id := uuid.New().String()

	requestData := models.Request{
		ID:              id,
		Problem_Details: data["problem_details"],
		Since_When:      data["since_when"],
		Customer_Name:   userSession["name"].(string),
		Customer_Email:  userSession["email"].(string),
		Customer_Phone:  data["c_phone"],
		Severity:        data["severity"],
		Category:        data["category"],
	}

	db.DB.Create(&requestData)

	return c.JSON(requestData)

}
