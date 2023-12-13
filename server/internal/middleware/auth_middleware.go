// auth_middleware.go

package middleware

import (
	"fmt"
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
			fmt.Println("No token provided")
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
			fmt.Println("Invalid token:", err)
			return c.Status(401).JSON(fiber.Map{"status": "error", "message": "Invalid token", "data": nil})
		}

		// Extract claims and user information as needed
		claims := token.Claims.(jwt.MapClaims)
		username := claims["username"].(string)

		// Fetch the user by username
		user, err := userRepository.FindByUsername(username)
		if err != nil {
			fmt.Println("Error fetching user:", err)
			return c.Status(401).JSON(fiber.Map{"status": "error", "message": "Invalid credentials", "data": nil})
		}

		// Add user information to the context
		c.Locals("user_id", user.ID.String())

		// Set claims
		claims["role"] = user.Role // Include the user's role in the claims

		// Log user information for debugging
		fmt.Println("User ID:", user.ID.String())

		// Proceed to RBAC check middleware
		return RBACMiddleware(e, userRepository)(c)

	}
}
