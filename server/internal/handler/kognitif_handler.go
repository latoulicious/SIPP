package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"github.com/latoulicious/SIPP/internal/service"
)

type KognitifHandler struct {
	KognitifService *service.KognitifService
}

func NewKognitifHandler(kognitifService *service.KognitifService) *KognitifHandler {
	return &KognitifHandler{
		KognitifService: kognitifService,
	}
}

func (handler *KognitifHandler) GetKognitif(c *fiber.Ctx) error {
	kognitif, err := handler.KognitifService.GetKognitif()

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error getting kognitif", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Kognitif retrieved successfully", "data": kognitif})
}

func (handler *KognitifHandler) GetKognitifByID(c *fiber.Ctx) error {
	kognitifID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	kognitif, err := handler.KognitifService.GetKognitifByID(kognitifID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error getting kognitif", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Kognitif retrieved successfully", "data": kognitif})
}

// CountRelatedQuestionsHandler handles requests to get the total count of related questions for all Kognitif records
func (handler *KognitifHandler) CountRelatedQuestionsHandler(c *fiber.Ctx) error {
	// Log the start of the handler
	log.Println("CountRelatedQuestionsHandler started")

	// Call the service method to get the kognitifs with question counts
	kognitifsWithQuestionCounts, err := handler.KognitifService.CountRelatedQuestionsForAll()
	if err != nil {
		// Log the error
		log.Printf("Error in CountRelatedQuestionsHandler: %v\n", err)

		// Handle error, e.g., write an error response
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error getting kognitif question counts", "data": err})
	}

	// Log the successful retrieval of kognitifs with question counts
	log.Println("Successfully retrieved kognitifs with question counts")

	// Serialize the kognitifs with question counts to JSON and write the response
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Kognitif question counts retrieved successfully", "data": kognitifsWithQuestionCounts})
}

func (handler *KognitifHandler) CreateKognitif(c *fiber.Ctx) error {

	kognitif := new(model.Kognitif)

	// Log request body for debugging purposes
	log.Printf("Request Body: %+v\n", kognitif)
	log.Printf("Received Kognitif object: %+v\n", kognitif)

	err := c.BodyParser(kognitif)
	if err != nil {
		// Log parsing error for debugging purposes
		log.Printf("Error parsing request body: %v\n", err)
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	err = handler.KognitifService.CreateKognitif(kognitif)
	if err != nil {
		// Log creation error for debugging purposes
		log.Printf("Error creating kognitif: %v\n", err)
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create kognitif", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Kognitif created", "data": kognitif})
}

func (handler *KognitifHandler) UpdateKognitif(c *fiber.Ctx) error {
	return c.Status(405).JSON(fiber.Map{"status": "error", "message": "Updates are not supported for this resource", "data": nil})
}

func (handler *KognitifHandler) DeleteKognitif(c *fiber.Ctx) error {
	kognitifID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	// Delete the Kognitif from the database
	err = handler.KognitifService.DeleteKognitif(kognitifID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't delete kognitif", "data": err})
	}

	// Return success message
	return c.JSON(fiber.Map{"status": "success", "message": "Kognitif deleted", "data": nil})
}
