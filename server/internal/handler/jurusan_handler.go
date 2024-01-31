package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"github.com/latoulicious/SIPP/internal/service"
)

type JurusanHandler struct {
	JurusanService *service.JurusanService
}

func NewJurusanHandler(jurusanService *service.JurusanService) *JurusanHandler {
	return &JurusanHandler{
		JurusanService: jurusanService,
	}
}

func (handler *JurusanHandler) GetJurusan(c *fiber.Ctx) error {
	jurusan, err := handler.JurusanService.GetJurusan()

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error getting jurusan", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Jurusan retrieved successfully", "data": jurusan})
}

func (handler *JurusanHandler) GetJurusanByID(c *fiber.Ctx) error {
	jurusanID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Jurusan retrieved successfully", "data": jurusanID})
}

func (handler *JurusanHandler) GetJurusanPublic(c *fiber.Ctx) error {
	// Implement logic to fetch all jurusan without requiring JWT authentication
	jurusan, err := handler.JurusanService.GetJurusanPublic()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error getting jurusan", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Jurusan retrieved successfully", "data": jurusan})
}

func (handler *JurusanHandler) CreateJurusan(c *fiber.Ctx) error {
	jurusan := new(model.Jurusan)

	err := c.BodyParser(jurusan)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	// If the ID field was not included in the request body, generate a new UUID
	if jurusan.ID == uuid.Nil {
		jurusan.ID = uuid.New()
	}

	err = handler.JurusanService.CreateJurusan(jurusan)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create jurusan", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Jurusan created", "data": jurusan})
}

func (handler *JurusanHandler) UpdateJurusan(c *fiber.Ctx) error {
	jurusanID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	// Call the service method to get the jurusan by ID
	jurusan, err := handler.JurusanService.GetJurusanByID(jurusanID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Jurusan not found", "data": nil})
	}

	// Store the request body in the jurusan and return an error if encountered
	err = c.BodyParser(jurusan)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	// Call the service method to update the jurusan
	err = handler.JurusanService.UpdateJurusan(jurusan)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not update jurusan", "data": err})
	}

	// Return the updated jurusan
	return c.JSON(fiber.Map{"status": "success", "message": "Jurusan updated", "data": jurusan})
}

func (handler *JurusanHandler) DeleteJurusan(c *fiber.Ctx) error {
	jurusanID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	// Call the service method to delete the jurusan by ID
	err = handler.JurusanService.DeleteJurusan(jurusanID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not delete jurusan", "data": err})
	}

	// Return success message
	return c.JSON(fiber.Map{"status": "success", "message": "Jurusan deleted"})
}
