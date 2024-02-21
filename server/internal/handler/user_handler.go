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

// NewUserHandler returns a new UserHandler instance with the given UserService.

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

// GetUsers retrieves all users from the database
// and returns them in a fiber response.
// It returns a 500 error if there is a problem retrieving the users.
// Otherwise it returns a 200 OK with the users in the response body.

func (handler *UserHandler) GetUsers(c *fiber.Ctx) error {
	users, err := handler.UserService.GetUsers()

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error getting users", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Users retrieved successfully", "data": users})
}

// GetUserByID retrieves a user by ID from the database.
// It parses the ID from the route parameter "id".
// Returns 400 if the ID is invalid.
// Calls the UserService to get the user by ID.
// Returns 404 if the user is not found.
// Otherwise returns 200 with the user in the response body.

func (handler *UserHandler) GetUserByID(c *fiber.Ctx) error {
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

// GetAllUsersPublic retrieves all users without requiring authentication.
// It returns a 500 error if there is a problem retrieving the users.
// Otherwise it returns a 200 OK with the users in the response body.
func (handler *UserHandler) GetAllUsersPublic(c *fiber.Ctx) error {
	// Implement logic to fetch all users without requiring JWT authentication
	users, err := handler.UserService.GetUsersPublic()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error getting users", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Users retrieved successfully", "data": users})
}

// CreateUser handles creating a new user.
// It takes a user object from the request body,
// passes it to the UserService to create the user,
// and returns the created user object in the response.
// Returns 500 error if there is an error creating the user.
// Returns 200 with the created user object on success.
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

// UpdateUser updates a user in the database.
// It takes a user ID as a parameter.
// It returns a 400 status if the ID is invalid.
// It calls the UserService to get the user by ID.
// It returns a 404 status if the user is not found.
// It parses the request body into the user struct.
// It returns a 500 status if parsing fails.
// It calls the UserService to update the user.
// It returns a 500 status if the update fails.
// If successful, it returns a 200 status with the updated user.

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

// DeleteUser deletes a user by ID.
// It takes a user ID from the request parameters,
// calls the UserService to delete the user by that ID,
// and returns a 200 response with a success message on success.
// Returns a 400 response if the ID is invalid.
// Returns a 404 response if the user is not found.
// Returns a 500 response if there is an error deleting the user.

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

// ChangePassword changes the password for the user with the given ID.
// It parses the user ID from the URL parameters and returns 400 if invalid.
// It parses the new password from the request body and returns 400 if invalid.
// It calls the UserService to change the password.
// It returns 500 if the service call fails.
// On success it returns 200 with a success message.

func (handler *UserHandler) ChangePassword(c *fiber.Ctx) error {
	userID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	var requestBody struct {
		NewPassword string `json:"newPassword"`
	}

	if err := c.BodyParser(&requestBody); err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid request body", "data": nil})
	}

	err = handler.UserService.ChangePassword(userID, requestBody.NewPassword)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not change password", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Password changed successfully", "data": nil})
}
