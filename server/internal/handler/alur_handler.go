package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"github.com/latoulicious/SIPP/internal/service"
)

type AlurHandler struct {
	AlurService *service.AlurService
}

func NewAlurHandler(alurService *service.AlurService) *AlurHandler {
	return &AlurHandler{
		AlurService: alurService,
	}
}

func (handler *AlurHandler) GetAlur(c *fiber.Ctx) error {
	alur, err := handler.AlurService.GetAlur()

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error getting alur", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Alur retrieved successfully", "data": alur})
}

func (handler *AlurHandler) GetAlurByID(c *fiber.Ctx) error {
	alurID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	alur, err := handler.AlurService.GetAlurByID(alurID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error getting alur", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Alur retrieved successfully", "data": alur})
}

func (handler *AlurHandler) CreateAlur(c *fiber.Ctx) error {
	alur := new(model.AlurTP)

	// Log request body for debugging purposes
	log.Printf("Request Body: %+v\n", alur)

	err := c.BodyParser(alur)
	if err != nil {
		// Log parsing error for debugging purposes
		log.Printf("Error parsing request body: %v\n", err)
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	err = handler.AlurService.CreateAlur(alur)
	if err != nil {
		// Log creation error for debugging purposes
		log.Printf("Error creating alur: %v\n", err)
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create alur", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Alur created", "data": alur})

}

func (handler *AlurHandler) UpdateAlur(c *fiber.Ctx) error {
	alurID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	// Retrieve the existing Alur by ID
	alur, err := handler.AlurService.GetAlurByID(alurID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Alur not found", "data": nil})
	}

	// Parse the JSON body into the existing Alur struct
	err = c.BodyParser(alur)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	// Update the Alur in the database
	err = handler.AlurService.UpdateAlur(alur)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't update alur", "data": err})
	}

	// Return the updated Alur
	return c.JSON(fiber.Map{"status": "success", "message": "Alur updated", "data": alur})
}

func (handler *AlurHandler) DeleteAlur(c *fiber.Ctx) error {
	alurID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	// Delete the Alur from the database
	err = handler.AlurService.DeleteAlur(alurID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't delete alur", "data": err})
	}

	// Return success message
	return c.JSON(fiber.Map{"status": "success", "message": "Alur deleted", "data": nil})
}
