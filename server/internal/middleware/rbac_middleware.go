// middleware/rbac_middleware.go

package middleware

import (
	"log"
	"os"
	"strings"

	"github.com/casbin/casbin/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/repository"
)

// RBACMiddleware is a middleware function to check RBAC on protected endpoints
func RBACMiddleware(e *casbin.Enforcer, userRepository *repository.UserRepository) fiber.Handler {
	// Exclude certain paths from RBAC checks
	excludePaths := []string{"/api/user/:id/change-password"}

	return func(c *fiber.Ctx) error {
		// Check if the current path should be excluded from RBAC checks
		for _, path := range excludePaths {
			if strings.HasPrefix(c.Path(), path) && c.Method() == "POST" {
				return c.Next()
			}
		}

		// Extract user ID from Bearer token
		userIDStr, ok := c.Locals("user_id").(string)
		if !ok {
			return c.Status(401).JSON(fiber.Map{"status": "error", "message": "User ID not found in context", "data": nil})
		}

		// Convert the userID to UUID
		userID, err := uuid.Parse(userIDStr)
		if err != nil {
			return c.Status(401).JSON(fiber.Map{"status": "error", "message": "Invalid user ID", "data": nil})
		}

		// Fetch the user by user ID
		user, err := userRepository.GetUserByID(userID)
		if err != nil {
			log.Println("Error fetching user:", err)
			return c.Status(401).JSON(fiber.Map{"status": "error", "message": "Invalid user", "data": nil})
		}

		if os.Getenv("ENVIRONMENT") == "development" {
			// Debugging statements
			log.Println("RBAC Check - User ID:", userIDStr)
			log.Println("RBAC Check - User Role:", user.Role)
			log.Println("RBAC Check - Path:", c.Path())
			log.Println("RBAC Check - Method:", c.Method())
		}

		// Use user role in the enforcement check
		if ok, err := e.Enforce(user.Role, c.Path(), c.Method()); !ok {
			log.Println("RBAC Check Failed:", err)
			return c.Status(403).JSON(fiber.Map{"status": "error", "message": "Forbidden", "data": nil})
		}

		// Proceed to the next middleware or handler
		return c.Next()
	}
}
