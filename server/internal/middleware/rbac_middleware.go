// middleware/rbac_middleware.go

package middleware

import (
	"fmt"

	"github.com/casbin/casbin/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/repository"
)

// Exclude the /api/get-user-role endpoint from RBAC checks
var excludePaths = []string{"/api/get-user-role"}

// RBACMiddleware is a middleware function to check RBAC on protected endpoints
func RBACMiddleware(e *casbin.Enforcer, userRepository *repository.UserRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if c.Path() == "/api/get-user-role" {
			return c.Next()
		}

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
			fmt.Println("Error fetching user:", err)
			return c.Status(401).JSON(fiber.Map{"status": "error", "message": "Invalid user", "data": nil})
		}

		// Debugging statements
		fmt.Println("RBAC Check - User ID:", userIDStr)
		fmt.Println("RBAC Check - User Role:", user.Role)
		fmt.Println("RBAC Check - Path:", c.Path())
		fmt.Println("RBAC Check - Method:", c.Method())

		// Use the actual path of the request in the enforcement check
		if ok, err := e.Enforce(user.Role, c.Path(), c.Method()); !ok {
			fmt.Println("RBAC Check Failed:", err)
			return c.Status(403).JSON(fiber.Map{"status": "error", "message": "Forbidden", "data": nil})
		}

		// Proceed to the next middleware or handler
		return c.Next()
	}
}
