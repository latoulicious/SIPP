// middleware/auth_middleware.go

package middleware

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/latoulicious/SIPP/internal/config"
)

// AuthMiddleware is a middleware function to check JWT token on protected endpoints
func AuthMiddleware(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		fmt.Println("No token provided")
		return c.Status(401).JSON(fiber.Map{"status": "error", "message": "Unauthorized", "data": nil})
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.JwtSecret), nil
	})
	if err != nil || !token.Valid {
		fmt.Println("Invalid token:", err)
		return c.Status(401).JSON(fiber.Map{"status": "error", "message": "Invalid token", "data": nil})
	}

	// Extract claims and user information as needed
	userID := token.Claims.(jwt.MapClaims)["username"].(string)

	// Add user information to the context
	c.Locals("user_id", userID)

	// Proceed to the next middleware or handler
	return c.Next()
}
