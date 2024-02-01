package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"github.com/latoulicious/SIPP/internal/service"
)

type PengayaanHandler struct {
	PengayaanService *service.PengayaanService
}

func NewPengayaanHandler(pengayaanService *service.PengayaanService) *PengayaanHandler {
	return &PengayaanHandler{
		PengayaanService: pengayaanService,
	}
}

func (handler *PengayaanHandler) GetPengayaan(c *fiber.Ctx) error {
	pengayaan, err := handler.PengayaanService.GetPengayaan()

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error getting pengayaan", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Pengayaan retrieved successfully", "data": pengayaan})
}

func (handler *PengayaanHandler) GetPengayaanByID(c *fiber.Ctx) error {
	pengayaanID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	pengayaan, err := handler.PengayaanService.GetPengayaanByID(pengayaanID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error getting pengayaan", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Pengayaan retrieved successfully", "data": pengayaan})
}

func (handler *PengayaanHandler) CreatePengayaan(c *fiber.Ctx) error {

	pengayaan := new(model.Pengayaan)

	// Log request body for debugging purposes
	log.Printf("Request Body: %+v\n", pengayaan)
	log.Printf("Received Pengayaan object: %+v\n", pengayaan)

	err := c.BodyParser(pengayaan)
	if err != nil {
		// Log parsing error for debugging purposes
		log.Printf("Error parsing request body: %v\n", err)
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	err = handler.PengayaanService.CreatePengayaan(pengayaan)
	if err != nil {
		// Log creation error for debugging purposes
		log.Printf("Error creating pengayaan: %v\n", err)
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create pengayaan", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Pengayaan created", "data": pengayaan})
}

func (handler *PengayaanHandler) UpdatePengayaan(c *fiber.Ctx) error {
	pengayaanID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	pengayaan := new(model.Pengayaan)
	pengayaan.ID = pengayaanID

	// Parse the JSON body into the existing Pengayaan struct
	err = c.BodyParser(pengayaan)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	log.Printf("Updated Pengayaan: %+v\n", pengayaan)
	log.Print(pengayaan.UserID)

	// Update the Pengayaan in the database
	err = handler.PengayaanService.UpdatePengayaan(pengayaan)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't update pengayaan", "data": err})
	}

	// Retrieve the existing Pengayaan by ID
	pengayaan, err = handler.PengayaanService.GetPengayaanByID(pengayaanID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Pengayaan not found", "data": nil})
	}

	// Return the updated Pengayaan
	return c.JSON(fiber.Map{"status": "success", "message": "Pengayaan updated", "data": pengayaan})
}

func (handler *PengayaanHandler) DeletePengayaan(c *fiber.Ctx) error {
	pengayaanID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	// Delete the Pengayaan from the database
	err = handler.PengayaanService.DeletePengayaan(pengayaanID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't delete pengayaan", "data": err})
	}

	// Return success message
	return c.JSON(fiber.Map{"status": "success", "message": "Pengayaan deleted", "data": nil})
}
