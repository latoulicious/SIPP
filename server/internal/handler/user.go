// handler/users.go

package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/database"
	models "github.com/latoulicious/SIPP/internal/model/user"
)

// GetUsers retrieves users from the database.
//
// Parameters:
// - c: The fiber.Ctx object representing the current HTTP request.
//
// Returns:
// - An error, if there was an error retrieving the users.
func GetUsers(c *fiber.Ctx) error {
	db := database.DB
	var users []models.Users

	db.Find(&users)

	if len(users) == 0 {
		return c.Status(404).JSON(fiber.Map{"Status": "error", "Message": "No users found", "data": nil})
	}

	return c.JSON(fiber.Map{"Status": "success", "Message": "Users retrieved successfully", "data": users})
}

func CreateUsers(c *fiber.Ctx) error {
	db := database.DB
	user := new(models.Users)

	err := c.BodyParser(user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	user.ID = uuid.New()
	err = db.Create(&user).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create user", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "User created successfully", "data": user})

}

func GetUsersById(c *fiber.Ctx) error {
	db := database.DB
	var user models.Users

	id := c.Params("usersId")

	db.Find(&user, "id = ?", id)

	if user.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No user found with ID", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "User found", "data": user})

}

func UpdateUsers(c *fiber.Ctx) error {
	type UpdateUsers struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Name     string `json:"name"`
		Mapel    string `json:"mapel"`
		Role     string `json:"role"`
	}
	db := database.DB
	var user models.Users

	id := c.Params("usersId")

	db.Find(&user, "id = ?", id)

	if user.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No user found with ID", "data": nil})
	}

	var UpdateUsersData UpdateUsers
	err := c.BodyParser(&UpdateUsersData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	user.Username = UpdateUsersData.Username
	user.Password = UpdateUsersData.Password
	user.Name = UpdateUsersData.Name
	user.Mapel = UpdateUsersData.Mapel
	user.Role = UpdateUsersData.Role

	db.Save(&user)

	return c.JSON(fiber.Map{"status": "success", "message": "User updated successfully", "data": user})
}

func DeleteUsers(c *fiber.Ctx) error {
	db := database.DB
	var user models.Users

	id := c.Params("usersId")

	db.Find(&user, "id = ?", id)

	if user.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No user found with ID", "data": nil})
	}

	err := db.Delete(&user, "id = ?", id).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't delete user", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "User deleted successfully", "data": nil})

}
