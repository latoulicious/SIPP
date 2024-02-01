package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"github.com/latoulicious/SIPP/internal/service"
)

type SoalHandler struct {
	SoalService *service.SoalService
}

func NewSoalHandler(soalService *service.SoalService) *SoalHandler {
	return &SoalHandler{
		SoalService: soalService,
	}
}

func (handler *SoalHandler) GetSoal(c *fiber.Ctx) error {
	soal, err := handler.SoalService.GetSoal()

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error getting soal", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Soal retrieved successfully", "data": soal})
}

func (handler *SoalHandler) GetSoalByID(c *fiber.Ctx) error {
	soalID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	soal, err := handler.SoalService.GetSoalByID(soalID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error getting soal", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Soal retrieved successfully", "data": soal})
}

func (handler *SoalHandler) CreateSoal(c *fiber.Ctx) error {

	soal := new(model.Soal)

	// Log request body for debugging purposes
	log.Printf("Request Body: %+v\n", soal)
	log.Printf("Received Soal object: %+v\n", soal)

	err := c.BodyParser(soal)
	if err != nil {
		// Log parsing error for debugging purposes
		log.Printf("Error parsing request body: %v\n", err)
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	err = handler.SoalService.CreateSoal(soal)
	if err != nil {
		// Log creation error for debugging purposes
		log.Printf("Error creating soal: %v\n", err)
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create soal", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Soal created", "data": soal})
}

func (handler *SoalHandler) UpdateSoal(c *fiber.Ctx) error {
	soalID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	soal := new(model.Soal)
	soal.ID = soalID

	// Parse the JSON body into the existing Soal struct
	err = c.BodyParser(soal)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	log.Printf("Updated Soal: %+v\n", soal)

	// Update the Soal in the database
	err = handler.SoalService.UpdateSoal(soal)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't update soal", "data": err})
	}

	// Retrieve the existing Soal by ID
	soal, err = handler.SoalService.GetSoalByID(soalID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Soal not found", "data": nil})
	}

	// Return the updated Soal
	return c.JSON(fiber.Map{"status": "success", "message": "Soal updated", "data": soal})
}

func (handler *SoalHandler) DeleteSoal(c *fiber.Ctx) error {
	soalID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	// Delete the Soal from the database
	err = handler.SoalService.DeleteSoal(soalID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't delete soal", "data": err})
	}

	// Return success message
	return c.JSON(fiber.Map{"status": "success", "message": "Soal deleted", "data": nil})
}
