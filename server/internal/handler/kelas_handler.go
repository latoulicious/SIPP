package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"github.com/latoulicious/SIPP/internal/service"
)

type KelasHandler struct {
	KelasService *service.KelasService
}

func NewKelasHandler(kelasService *service.KelasService) *KelasHandler {
	return &KelasHandler{
		KelasService: kelasService,
	}
}

func (handler *KelasHandler) GetKelas(c *fiber.Ctx) error {
	kelas, err := handler.KelasService.GetKelas()

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error getting kelas", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Kelas retrieved successfully", "data": kelas})
}

func (handler *KelasHandler) GetKelasByID(c *fiber.Ctx) error {
	kelasID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Kelas retrieved successfully", "data": kelasID})
}

func (handler *KelasHandler) GetKelasPublic(c *fiber.Ctx) error {
	// Implement logic to fetch all kelas without requiring JWT authentication
	kelas, err := handler.KelasService.GetKelasPublic()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error getting kelas", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Kelas retrieved successfully", "data": kelas})
}

func (handler *KelasHandler) CreateKelas(c *fiber.Ctx) error {
	kelas := new(model.Kelas)

	err := c.BodyParser(kelas)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	err = handler.KelasService.CreateKelas(kelas)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create kelas", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Kelas created", "data": kelas})
}

func (handler *KelasHandler) UpdateKelas(c *fiber.Ctx) error {
	kelasID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	// Call the service method to get the kelas by ID
	kelas, err := handler.KelasService.GetKelasByID(kelasID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Kelas not found", "data": nil})
	}

	// Store the request body in the kelas and return an error if encountered
	err = c.BodyParser(kelas)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	// Call the service method to update the kelas
	err = handler.KelasService.UpdateKelas(kelas)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not update kelas", "data": err})
	}

	// Return the updated kelas
	return c.JSON(fiber.Map{"status": "success", "message": "Kelas updated", "data": kelas})
}

func (handler *KelasHandler) DeleteKelas(c *fiber.Ctx) error {
	kelasID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	// Call the service method to delete the kelas by ID
	err = handler.KelasService.DeleteKelas(kelasID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not delete kelas", "data": err})
	}

	// Return success message
	return c.JSON(fiber.Map{"status": "success", "message": "Kelas deleted"})
}
