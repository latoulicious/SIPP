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

// CountRelatedQuestionsHandler handles requests to get the total count of related questions for all Sumatif records
func (handler *SumatifHandler) CountRelatedQuestionsHandler(c *fiber.Ctx) error {
	// Log the start of the handler
	log.Println("CountRelatedQuestionsHandler started")

	// Call the service method to get the sumatifs with question counts
	sumatifsWithQuestionCounts, err := handler.SumatifService.CountRelatedQuestionsForAll()
	if err != nil {
		// Log the error
		log.Printf("Error in CountRelatedQuestionsHandler: %v\n", err)

		// Handle error, e.g., write an error response
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error getting sumatif question counts", "data": err})
	}

	// Log the successful retrieval of sumatifs with question counts
	log.Println("Successfully retrieved sumatifs with question counts")

	// Serialize the sumatifs with question counts to JSON and write the response
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Sumatif question counts retrieved successfully", "data": sumatifsWithQuestionCounts})
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
	return c.Status(405).JSON(fiber.Map{"status": "error", "message": "Updates are not supported for this resource", "data": nil})
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
