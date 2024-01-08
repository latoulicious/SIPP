package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"github.com/latoulicious/SIPP/internal/service"
)

type ModulHandler struct {
	ModulService *service.ModulService
}

func NewModulHandler(modulService *service.ModulService) *ModulHandler {
	return &ModulHandler{
		ModulService: modulService,
	}
}

func (handler *ModulHandler) GetModul(c *fiber.Ctx) error {
	modul, err := handler.ModulService.GetModul()

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error getting modul", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Modul retrieved successfully", "data": modul})
}

func (handler *ModulHandler) GetModulByID(c *fiber.Ctx) error {
	modulID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	modul, err := handler.ModulService.GetModulByID(modulID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error getting modul", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Modul retrieved successfully", "data": modul})
}

func (handler *ModulHandler) CreateModul(c *fiber.Ctx) error {
	modul := new(model.ModulAjar)

	// Log request body for debugging purposes
	log.Printf("Request Body: %+v\n", modul)

	err := c.BodyParser(modul)
	if err != nil {
		// Log parsing error for debugging purposes
		log.Printf("Error parsing request body: %v\n", err)
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	err = handler.ModulService.CreateModul(modul)
	if err != nil {
		// Log creation error for debugging purposes
		log.Printf("Error creating modul: %v\n", err)
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create modul", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Modul created", "data": modul})
}

func (handler *ModulHandler) UpdateModul(c *fiber.Ctx) error {
	modulID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	// Retrieve the existing Modul by ID
	modul, err := handler.ModulService.GetModulByID(modulID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Modul not found", "data": nil})
	}

	// Parse the JSON body into the existing Modul struct
	err = c.BodyParser(modul)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	// Update the Modul in the database
	err = handler.ModulService.UpdateModul(&modul)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't update modul", "data": err})
	}

	// Return the updated Modul
	return c.JSON(fiber.Map{"status": "success", "message": "Modul updated", "data": modul})
}

func (handler *ModulHandler) DeleteModul(c *fiber.Ctx) error {
	modulID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	// Delete the Modul from the database
	err = handler.ModulService.DeleteModul(modulID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't delete modul", "data": err})
	}

	// Return success message
	return c.JSON(fiber.Map{"status": "success", "message": "Modul deleted", "data": nil})
}
