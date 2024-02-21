package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/service"
)

type ItemSoalHandler struct {
	ItemSoalService *service.ItemSoalService
}

func NewItemSoalHandler(itemsoalService *service.ItemSoalService) *ItemSoalHandler {
	return &ItemSoalHandler{
		ItemSoalService: itemsoalService,
	}
}

func (handler *ItemSoalHandler) GetItemSoal(c *fiber.Ctx) error {
	itemsoal, err := handler.ItemSoalService.GetItemSoal()

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error getting itemsoal", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "ItemSoal retrieved successfully", "data": itemsoal})
}

func (handler *ItemSoalHandler) GetItemSoalByID(c *fiber.Ctx) error {
	itemsoalID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	itemsoal, err := handler.ItemSoalService.GetItemSoalByID(itemsoalID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error getting itemsoal", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "ItemSoal retrieved successfully", "data": itemsoal})
}

// CountRelatedQuestionsHandler handles requests to get the total count of related questions for all ItemSoal records
func (handler *ItemSoalHandler) CountRelatedQuestionsHandler(c *fiber.Ctx) error {
	// Log the start of the handler
	log.Println("CountRelatedQuestionsHandler started")

	// Call the service method to get the itemsoals with question counts
	itemsoalsWithQuestionCounts, err := handler.ItemSoalService.CountRelatedQuestionsForAll()
	if err != nil {
		// Log the error
		log.Printf("Error in CountRelatedQuestionsHandler: %v\n", err)

		// Handle error, e.g., write an error response
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error getting itemsoal question counts", "data": err})
	}

	// Log the successful retrieval of itemsoals with question counts
	log.Println("Successfully retrieved itemsoals with question counts")

	// Serialize the itemsoals with question counts to JSON and write the response
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "ItemSoal question counts retrieved successfully", "data": itemsoalsWithQuestionCounts})
}
