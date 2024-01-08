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

func NewAuthHandler(userService *service.UserService) *AuthHandler {
	return &AuthHandler{
		UserService: userService,
	}
}

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
