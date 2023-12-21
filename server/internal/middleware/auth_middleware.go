// auth_middleware.go

package middleware

import (
	"log"
	"strings"

	"github.com/casbin/casbin/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/latoulicious/SIPP/internal/config"
	"github.com/latoulicious/SIPP/internal/repository"
)

// AuthMiddleware is a middleware function to check JWT token on protected endpoints
func AuthMiddleware(e *casbin.Enforcer, userRepository *repository.UserRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString := c.Get("Authorization")
		if tokenString == "" {
			log.Println("No token provided")
			return c.Status(401).JSON(fiber.Map{"status": "error", "message": "Unauthorized", "data": nil})
		}

		// Check for the "Bearer " prefix
		const prefix = "Bearer "
		if strings.HasPrefix(tokenString, prefix) {
			tokenString = tokenString[len(prefix):]
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.JwtSecret), nil
		})
		if err != nil || !token.Valid {
			log.Println("Invalid token:", err)
			return c.Status(401).JSON(fiber.Map{"status": "error", "message": "Invalid token", "data": nil})
		}

		// Extract claims and user information as needed
		claims := token.Claims.(jwt.MapClaims)
		log.Println("Decoded Claims in AuthMiddleware:", claims)
		username, okUsername := claims["username"].(string)
		role, okRole := claims["role"].(string)

		// Check if username and role are present in claims
		if !okUsername || !okRole {
			return c.Status(401).JSON(fiber.Map{"status": "error", "message": "Invalid token claims", "data": nil})
		}

		log.Println("Username in AuthMiddleware:", username)
		log.Println("Role in AuthMiddleware:", role)

		// Add user information to the context
		c.Locals("user_id", username) // Assuming username can serve as user_id
		c.Locals("user_role", role)   // Set the user's role in the context

		// Log user information for debugging
		log.Println("User ID:", username)
		log.Println("User Role:", role)

		// Proceed to RBAC check middleware
		return RBACMiddleware(e, userRepository)(c)
	}
}
