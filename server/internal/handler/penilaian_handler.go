package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"github.com/latoulicious/SIPP/internal/service"
)

type PenilaianHandler struct {
	PenilaianService *service.PenilaianService
}

func NewPenilaianHandler(penilaianService *service.PenilaianService) *PenilaianHandler {
	return &PenilaianHandler{
		PenilaianService: penilaianService,
	}
}

func (handler *PenilaianHandler) GetPenilaian(c *fiber.Ctx) error {
	penilaian, err := handler.PenilaianService.GetPenilaian()

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error getting penilaian", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Penilaian retrieved successfully", "data": penilaian})
}

func (handler *PenilaianHandler) GetPenilaianByID(c *fiber.Ctx) error {
	penilaianID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	penilaian, err := handler.PenilaianService.GetPenilaianByID(penilaianID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error getting penilaian", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Penilaian retrieved successfully", "data": penilaian})
}

func (handler *PenilaianHandler) CreatePenilaian(c *fiber.Ctx) error {
	penilaian := new(model.Penilaian)

	// Log request body for debugging purposes
	log.Printf("Request Body: %+v\n", penilaian)

	err := c.BodyParser(penilaian)
	if err != nil {
		// Log parsing error for debugging purposes
		log.Printf("Error parsing request body: %v\n", err)
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	err = handler.PenilaianService.CreatePenilaian(penilaian)
	if err != nil {
		// Log creation error for debugging purposes
		log.Printf("Error creating penilaian: %v\n", err)
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create penilaian", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Penilaian created", "data": penilaian})
}

func (handler *PenilaianHandler) UpdatePenilaian(c *fiber.Ctx) error {
	penilaianID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	// Retrieve the existing Penilaian by ID
	penilaian, err := handler.PenilaianService.GetPenilaianByID(penilaianID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error getting penilaian", "data": err})
	}

	// Parse the JSON body into the existing Penilaian struct
	err = c.BodyParser(penilaian)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	// Update the Penilaian in the database
	err = handler.PenilaianService.UpdatePenilaian(penilaian)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't update penilaian", "data": err})
	}

	// Return the updated Penilaian
	return c.JSON(fiber.Map{"status": "success", "message": "Penilaian updated", "data": penilaian})
}

func (handler *PenilaianHandler) DeletePenilaian(c *fiber.Ctx) error {
	penilaianID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	// Delete the Penilaian from the database
	err = handler.PenilaianService.DeletePenilaian(penilaianID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error deleting penilaian", "data": err})
	}

	// Return success message
	return c.JSON(fiber.Map{"status": "success", "message": "Success", "data": nil})
}
