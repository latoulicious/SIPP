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

// NewKognitifHandler creates a new KognitifHandler instance with the provided KognitifService.
// It is used to initialize a KognitifHandler with its service dependency.

func NewKognitifHandler(kognitifService *service.KognitifService) *KognitifHandler {
	return &KognitifHandler{
		KognitifService: kognitifService,
	}
}

// GetKognitif retrieves kognitif data from the KognitifService and returns
// it in a JSON response. If there is an error getting the kognitif data,
// it returns a 500 status with the error details.

func (handler *KognitifHandler) GetKognitif(c *fiber.Ctx) error {
	kognitif, err := handler.KognitifService.GetKognitif()

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error getting kognitif", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Kognitif retrieved successfully", "data": kognitif})
}

// GetKognitifByID retrieves a Kognitif by its ID.
// It parses the ID from the route parameter "id",
// calls the KognitifService to get the Kognitif,
// and returns the Kognitif with a 200 status.
// If there is an error parsing the ID or retrieving the Kognitif,
// it returns a 400 or 500 status with an error message.

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

// CountRelatedQuestionsHandler handles requests to get the total count of related questions for all Kognitif records.
// It calls the KognitifService to retrieve the kognitifs with their related question counts, logs the process, handles any errors,
// and returns the kognitifs with question counts in a JSON response if successful.
// This is an exported handler method so the documentation comment describes its overall purpose at a high level.

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

// CreateKognitif handles requests to create a new Kognitif record.
// It parses the Kognitif from the request body, calls the KognitifService
// to create it in the database, and returns the created Kognitif with a
// 201 status. If there are any errors parsing or creating, it returns
// a 500 status with an error message.

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

// UpdateKognitif handles requests to update an existing kognitif record.
// It returns a 405 status code indicating that updates are not supported for kognitifs.

func (handler *KognitifHandler) UpdateKognitif(c *fiber.Ctx) error {
	return c.Status(405).JSON(fiber.Map{"status": "error", "message": "Updates are not supported for this resource", "data": nil})
}

// DeleteKognitif handles requests to delete a Kognitif record by ID.
// It parses the Kognitif ID from the request parameters, calls the
// KognitifService to delete it from the database, and returns a
// success/error response.

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
