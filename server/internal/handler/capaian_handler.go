package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"github.com/latoulicious/SIPP/internal/service"
)

type CapaianHandler struct {
	CapaianService *service.CapaianService
}

func NewCapaianHandler(capaianService *service.CapaianService) *CapaianHandler {
	return &CapaianHandler{
		CapaianService: capaianService,
	}
}

func (handler *CapaianHandler) GetCapaian(c *fiber.Ctx) error {
	capaian, err := handler.CapaianService.GetCapaian()

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error getting capaian", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Capaian retrieved successfully", "data": capaian})
}

func (handler *CapaianHandler) GetCapaianByID(c *fiber.Ctx) error {
	capaianID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	capaian, err := handler.CapaianService.GetCapaianByID(capaianID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error getting capaian", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Capaian retrieved successfully", "data": capaian})
}

func (handler *CapaianHandler) CreateCapaian(c *fiber.Ctx) error {

	capaian := new(model.Capaian)

	// Log request body for debugging purposes
	log.Printf("Request Body: %+v\n", capaian)
	log.Printf("Received Capaian object: %+v\n", capaian)

	err := c.BodyParser(capaian)
	if err != nil {
		// Log parsing error for debugging purposes
		log.Printf("Error parsing request body: %v\n", err)
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	err = handler.CapaianService.CreateCapaian(capaian)
	if err != nil {
		// Log creation error for debugging purposes
		log.Printf("Error creating capaian: %v\n", err)
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create capaian", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Capaian created", "data": capaian})
}

func (handler *CapaianHandler) UpdateCapaian(c *fiber.Ctx) error {
	capaianID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	capaian := new(model.Capaian)
	capaian.ID = capaianID

	// Parse the JSON body into the existing Capaian struct
	err = c.BodyParser(capaian)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	log.Printf("Updated Capaian: %+v\n", capaian)
	log.Print(capaian.UserID)

	// Update the Capaian in the database
	err = handler.CapaianService.UpdateCapaian(capaian)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't update capaian", "data": err})
	}

	// Retrieve the existing Capaian by ID
	capaian, err = handler.CapaianService.GetCapaianByID(capaianID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Capaian not found", "data": nil})
	}

	// Return the updated Capaian
	return c.JSON(fiber.Map{"status": "success", "message": "Capaian updated", "data": capaian})
}

func (handler *CapaianHandler) DeleteCapaian(c *fiber.Ctx) error {
	capaianID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	// Delete the Capaian from the database
	err = handler.CapaianService.DeleteCapaian(capaianID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't delete capaian", "data": err})
	}

	// Return success message
	return c.JSON(fiber.Map{"status": "success", "message": "Capaian deleted", "data": nil})
}
