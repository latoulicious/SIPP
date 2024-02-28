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

// NewItemSoalHandler creates a new ItemSoalHandler with the given ItemSoalService.
// It is used to instantiate the ItemSoalHandler with its required dependencies.

func NewItemSoalHandler(itemsoalService *service.ItemSoalService) *ItemSoalHandler {
	return &ItemSoalHandler{
		ItemSoalService: itemsoalService,
	}
}

// GetItemSoal retrieves all itemsoal from the database
// and returns them in a fiber.Map with status.

func (handler *ItemSoalHandler) GetItemSoal(c *fiber.Ctx) error {
	itemsoal, err := handler.ItemSoalService.GetItemSoal()

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error getting itemsoal", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "ItemSoal retrieved successfully", "data": itemsoal})
}

// GetItemSoalByID retrieves an itemsoal by its ID.
// It parses the ID from the route parameter "id".
// Returns a 404 if the ID is invalid.
// Returns a 500 if there is an error retrieving the itemsoal.
// Otherwise returns the itemsoal with a 200 status.

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

// CountRelatedQuestionsHandler handles requests to get the total count of
// related questions for all ItemSoal records. It calls the ItemSoalService
// to retrieve the ItemSoal records with their related question counts,
// logs and handles any errors, and responds with the question count data.

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
