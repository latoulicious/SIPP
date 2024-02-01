package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"github.com/latoulicious/SIPP/internal/service"
)

type SumatifHandler struct {
	SumatifService *service.SumatifService
}

func NewSumatifHandler(sumatifService *service.SumatifService) *SumatifHandler {
	return &SumatifHandler{
		SumatifService: sumatifService,
	}
}

func (handler *SumatifHandler) GetSumatif(c *fiber.Ctx) error {
	sumatif, err := handler.SumatifService.GetSumatif()

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error getting sumatif", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Sumatif retrieved successfully", "data": sumatif})
}

func (handler *SumatifHandler) GetSumatifByID(c *fiber.Ctx) error {
	sumatifID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	sumatif, err := handler.SumatifService.GetSumatifByID(sumatifID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error getting sumatif", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Sumatif retrieved successfully", "data": sumatif})
}

func (handler *SumatifHandler) CreateSumatif(c *fiber.Ctx) error {

	sumatif := new(model.Sumatif)

	// Log request body for debugging purposes
	log.Printf("Request Body: %+v\n", sumatif)
	log.Printf("Received Sumatif object: %+v\n", sumatif)

	err := c.BodyParser(sumatif)
	if err != nil {
		// Log parsing error for debugging purposes
		log.Printf("Error parsing request body: %v\n", err)
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	err = handler.SumatifService.CreateSumatif(sumatif)
	if err != nil {
		// Log creation error for debugging purposes
		log.Printf("Error creating sumatif: %v\n", err)
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create sumatif", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Sumatif created", "data": sumatif})
}

func (handler *SumatifHandler) UpdateSumatif(c *fiber.Ctx) error {
	sumatifID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	sumatif := new(model.Sumatif)
	sumatif.ID = sumatifID

	// Parse the JSON body into the existing Sumatif struct
	err = c.BodyParser(sumatif)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	log.Printf("Updated Sumatif: %+v\n", sumatif)
	log.Print(sumatif.UserID)

	// Update the Sumatif in the database
	err = handler.SumatifService.UpdateSumatif(sumatif)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't update sumatif", "data": err})
	}

	// Retrieve the existing Sumatif by ID
	sumatif, err = handler.SumatifService.GetSumatifByID(sumatifID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Sumatif not found", "data": nil})
	}

	// Return the updated Sumatif
	return c.JSON(fiber.Map{"status": "success", "message": "Sumatif updated", "data": sumatif})
}

func (handler *SumatifHandler) DeleteSumatif(c *fiber.Ctx) error {
	sumatifID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	// Delete the Sumatif from the database
	err = handler.SumatifService.DeleteSumatif(sumatifID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't delete sumatif", "data": err})
	}

	// Return success message
	return c.JSON(fiber.Map{"status": "success", "message": "Sumatif deleted", "data": nil})
}
