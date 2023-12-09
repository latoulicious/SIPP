package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"github.com/latoulicious/SIPP/internal/service"
)

type UserHandler struct {
	UserService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

func (handler *UserHandler) GetUsers(c *fiber.Ctx) error {
	users, err := handler.UserService.GetUsers()

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error getting users", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Users retrieved successfully", "data": users})
}

func (handler *UserHandler) GetUser(c *fiber.Ctx) error {
	userID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	// Call the service method to get the user by ID
	user, err := handler.UserService.GetUserByID(userID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found", "data": nil})
	}

	// Return the user
	return c.JSON(fiber.Map{"status": "success", "message": "User Found", "data": user})
}

func (handler *UserHandler) CreateUser(c *fiber.Ctx) error {
	user := new(model.Users)

	err := c.BodyParser(user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	err = handler.UserService.CreateUser(user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create user", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "User created", "data": user})

}

func (handler *UserHandler) UpdateUser(c *fiber.Ctx) error {
	userID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	// Call the service method to get the user by ID
	user, err := handler.UserService.GetUserByID(userID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found", "data": nil})
	}

	// Store the request body in the user and return an error if encountered
	err = c.BodyParser(user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	// Call the service method to update the user
	err = handler.UserService.UpdateUser(user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not update user", "data": err})
	}

	// Return the updated user
	return c.JSON(fiber.Map{"status": "success", "message": "Updated User", "data": user})
}

func (handler *UserHandler) DeleteUser(c *fiber.Ctx) error {
	userID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	// Call the service method to delete the user by ID
	err = handler.UserService.DeleteUser(userID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not delete user", "data": err})
	}

	// Return success message
	return c.JSON(fiber.Map{"status": "success", "message": "Deleted User"})
}
