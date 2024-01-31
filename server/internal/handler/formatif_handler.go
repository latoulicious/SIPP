package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"github.com/latoulicious/SIPP/internal/service"
)

type FormatifHandler struct {
	FormatifService *service.FormatifService
}

func NewFormatifHandler(formatifService *service.FormatifService) *FormatifHandler {
	return &FormatifHandler{
		FormatifService: formatifService,
	}
}

func (handler *FormatifHandler) GetFormatif(c *fiber.Ctx) error {
	formatif, err := handler.FormatifService.GetFormatif()

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error getting formatif", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Formatif retrieved successfully", "data": formatif})
}

func (handler *FormatifHandler) GetFormatifByID(c *fiber.Ctx) error {
	formatifID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	formatif, err := handler.FormatifService.GetFormatifByID(formatifID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error getting formatif", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Formatif retrieved successfully", "data": formatif})
}

func (handler *FormatifHandler) CreateFormatif(c *fiber.Ctx) error {

	formatif := new(model.Formatif)

	// Log request body for debugging purposes
	log.Printf("Request Body: %+v\n", formatif)
	log.Printf("Received Formatif object: %+v\n", formatif)

	err := c.BodyParser(formatif)
	if err != nil {
		// Log parsing error for debugging purposes
		log.Printf("Error parsing request body: %v\n", err)
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	err = handler.FormatifService.CreateFormatif(formatif)
	if err != nil {
		// Log creation error for debugging purposes
		log.Printf("Error creating formatif: %v\n", err)
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create formatif", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Formatif created", "data": formatif})
}

func (handler *FormatifHandler) UpdateFormatif(c *fiber.Ctx) error {
	formatifID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	formatif := new(model.Formatif)
	formatif.ID = formatifID

	// Parse the JSON body into the existing Formatif struct
	err = c.BodyParser(formatif)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	log.Printf("Updated Formatif: %+v\n", formatif)
	log.Print(formatif.UserID)

	// Update the Formatif in the database
	err = handler.FormatifService.UpdateFormatif(formatif)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't update formatif", "data": err})
	}

	// Retrieve the existing Formatif by ID
	formatif, err = handler.FormatifService.GetFormatifByID(formatifID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Formatif not found", "data": nil})
	}

	// Return the updated Formatif
	return c.JSON(fiber.Map{"status": "success", "message": "Formatif updated", "data": formatif})
}

func (handler *FormatifHandler) DeleteFormatif(c *fiber.Ctx) error {
	formatifID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	// Delete the Formatif from the database
	err = handler.FormatifService.DeleteFormatif(formatifID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't delete formatif", "data": err})
	}

	// Return success message
	return c.JSON(fiber.Map{"status": "success", "message": "Formatif deleted", "data": nil})
}
