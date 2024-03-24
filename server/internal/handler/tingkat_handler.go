package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/service"
)

type TingkatHandler struct {
	TingkatService *service.TingkatService
}

// NewTingkatHandler creates a new TingkatHandler with the given TingkatService
func NewTingkatHandler(tingkatService *service.TingkatService) *TingkatHandler {
	return &TingkatHandler{
		TingkatService: tingkatService,
	}
}

// GetTingkat retrieves the tingkat details from the database.
func (handler *TingkatHandler) GetIndikatorTingkat(c *fiber.Ctx) error {
	indikatorTingkat, err := handler.TingkatService.GetIndikatorTingkat()

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error getting tingkat", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Tingkat retrieved successfully", "data": indikatorTingkat})
}

// GetIndikatorTingkatByID retrieves a tingkat by its ID.
func (handler *TingkatHandler) GetIndikatorTingkatByID(c *fiber.Ctx) error {
	indikatorTingkatID, err := uuid.Parse(c.Params("id"))

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	indikatorTingkat, err := handler.TingkatService.GetIndikatorTingkatByID(indikatorTingkatID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error getting tingkat", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Tingkat retrieved successfully", "data": indikatorTingkat})
}

// CreateIndikatorTingkat creates a new tingkat in the database.
func (handler *TingkatHandler) CreateIndikatorTingkat(c *fiber.Ctx) error {

	// Define a struct to hold the body of the request.
	type RequestBody struct {
		IndikatorID uuid.UUID `json:"indikator_id"`
		KesukaranID uuid.UUID `json:"kesukaran_id"`
	}

	// Parse the request body.
	var body RequestBody
	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	// Log the parsed IDs for debugging purposes.
	log.Printf("Parsed Indikator ID: %s\n", body.IndikatorID)
	log.Printf("Parsed Kesukaran ID: %s\n", body.KesukaranID)

	// Call the service with the parsed IDs.
	err = handler.TingkatService.CreateIndikatorTingkat(body.IndikatorID, body.KesukaranID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error creating tingkat", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Tingkat created successfully", "data": nil})
}
