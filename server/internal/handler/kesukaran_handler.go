package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"github.com/latoulicious/SIPP/internal/service"
)

type KesukaranHandler struct {
	KesukaranService *service.KesukaranService
}

func NewKesukaranHandler(service *service.KesukaranService) *KesukaranHandler {
	return &KesukaranHandler{
		KesukaranService: service,
	}
}

func (handler *KesukaranHandler) GetKesukaran(c *fiber.Ctx) error {
	kesukaran, err := handler.KesukaranService.GetKesukaran()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error getting kesukaran", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Kesukaran retrieved successfully", "data": kesukaran})
}

func (handler *KesukaranHandler) GetKesukaranByID(c *fiber.Ctx) error {
	kesukaranID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	kesukaran, err := handler.KesukaranService.GetKesukaranByID(kesukaranID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error getting kesukaran", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Kesukaran retrieved successfully", "data": kesukaran})
}

func (handler *KesukaranHandler) CreateKesukaran(c *fiber.Ctx) error {
	rawBody := c.Body()
	log.Printf("Raw Request Body: %s\n", rawBody)

	kesukaran := new(model.Kesukaran)

	// Parse the request body into the kesukaran struct
	if err := c.BodyParser(kesukaran); err != nil {
		log.Printf("Error parsing request body: %v\n", err)
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	// Create the kesukaran in the database
	err := handler.KesukaranService.CreateKesukaran(kesukaran)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create kesukaran", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Kesukaran created successfully", "data": kesukaran})
}

func (handler *KesukaranHandler) UpdateKesukaran(c *fiber.Ctx) error {
	kesukaranID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	kesukaran := new(model.Kesukaran)
	kesukaran.ID = kesukaranID

	// Parse the request body into the kesukaran struct
	if err := c.BodyParser(kesukaran); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	// Update the kesukaran in the database
	err = handler.KesukaranService.UpdateKesukaran(kesukaran)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't update kesukaran", "data": err})
	}

	// Retrieve the existing kesukaran by ID
	kesukaran, err = handler.KesukaranService.GetKesukaranByID(kesukaranID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't retrieve kesukaran", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Kesukaran updated successfully", "data": kesukaran})
}

func (handler *KesukaranHandler) DeleteKesukaran(c *fiber.Ctx) error {
	kesukaranID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	// Delete the kesukaran in the database
	err = handler.KesukaranService.DeleteKesukaran(kesukaranID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't delete kesukaran", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Kesukaran deleted successfully", "data": nil})
}
