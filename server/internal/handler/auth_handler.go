// handler/auth_handler.go

package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/latoulicious/SIPP/internal/service"
)

type AuthHandler struct {
	UserService *service.UserService
}

// NewAuthHandler creates a new AuthHandler with the given UserService.

func NewAuthHandler(userService *service.UserService) *AuthHandler {
	return &AuthHandler{
		UserService: userService,
	}
}

// LoginHandler handles user login requests. It parses the username, password, name
// and role from the request body, calls the UserService to authenticate, and returns
// a JWT token on success or an error status code on failure.

func (handler *AuthHandler) LoginHandler(c *fiber.Ctx) error {
	// Print incoming request details for debugging purposes
	log.Printf("Incoming Request: %+v\n", c.Request())

	var request struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Name     string `json:"name"`
		Role     string `json:"role"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Invalid request format",
			"data":    nil,
		})
	}

	// Authenticate the user
	token, err := handler.UserService.Authenticate(request.Username, request.Password)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"status": "error", "message": "Authentication failed", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Authentication successful", "data": token})
}
