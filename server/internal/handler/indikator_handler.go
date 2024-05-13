package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"github.com/latoulicious/SIPP/internal/service"
)

type IndikatorHandler struct {
	IndikatorService *service.IndikatorService
}

func NewIndikatorHandler(service *service.IndikatorService) *IndikatorHandler {
	return &IndikatorHandler{
		IndikatorService: service,
	}
}

func (handler *IndikatorHandler) GetIndikator(c *fiber.Ctx) error {
	indikator, err := handler.IndikatorService.GetIndikator()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error getting indikator", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Indikator retrieved successfully", "data": indikator})
}

func (handler *IndikatorHandler) GetIndikatorByID(c *fiber.Ctx) error {
	indikatorID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	indikator, err := handler.IndikatorService.GetIndikatorByID(indikatorID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error getting indikator", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Indikator retrieved successfully", "data": indikator})
}

func (handler *IndikatorHandler) CreateIndikator(c *fiber.Ctx) error {
	rawBody := c.Body()
	log.Printf("Raw Request Body: %s\n", rawBody)

	indikator := new(model.Indikator)

	// Parse the request body into the indikator struct
	if err := c.BodyParser(indikator); err != nil {
		log.Printf("Error parsing request body: %v\n", err)
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	// Create the indikator in the database
	err := handler.IndikatorService.CreateIndikator(indikator)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create indikator", "data": err})
	}

	// Return success message
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Indikator created successfully", "data": indikator})
}

func (handler *IndikatorHandler) UpdateIndikator(c *fiber.Ctx) error {
	indikatorID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	indikator := new(model.Indikator)
	indikator.ID = indikatorID

	// Parse the JSON body into the existing Indikator struct
	if err := c.BodyParser(indikator); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	// Update the indikator in the database
	err = handler.IndikatorService.UpdateIndikator(indikator)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't update indikator", "data": err})
	}

	// Retrieve the existing Indikator by ID
	indikator, err = handler.IndikatorService.GetIndikatorByID(indikatorID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Indikator not found", "data": err})
	}

	// Return success message
	return c.JSON(fiber.Map{"status": "success", "message": "Indikator updated successfully", "data": indikator})
}

func (handler *IndikatorHandler) DeleteIndikator(c *fiber.Ctx) error {
	indikatorID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	// Delete the Indikator from the database
	err = handler.IndikatorService.DeleteIndikator(indikatorID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't delete indikator", "data": err})
	}

	// Return success message
	return c.JSON(fiber.Map{"status": "success", "message": "Indikator deleted successfully", "data": err})
}
