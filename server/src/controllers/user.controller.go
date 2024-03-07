package controllers

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
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