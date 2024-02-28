package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"github.com/latoulicious/SIPP/internal/service"
)

type FormatifHandler struct {
	FormatifService *service.FormatifService
}

// NewFormatifHandler creates a new FormatifHandler with the given FormatifService.
func NewFormatifHandler(formatifService *service.FormatifService) *FormatifHandler {
	return &FormatifHandler{
		FormatifService: formatifService,
	}
}

// GetFormatif retrieves the formatif data from the database
// via the FormatifService and returns it in the response.
// It returns a 500 error if there is a problem retrieving the data.
// Otherwise it returns a 200 with the formatif data.

func (handler *FormatifHandler) GetFormatif(c *fiber.Ctx) error {
	formatif, err := handler.FormatifService.GetFormatif()

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error getting formatif", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Formatif retrieved successfully", "data": formatif})
}

// GetFormatifByID retrieves a formatif by its ID.
// It parses the ID from the route parameter "id".
// Returns a 404 if the ID is invalid.
// Returns a 500 if there is an error retrieving the formatif.
// Otherwise returns a 200 with the formatif data.

func (handler *FormatifHandler) GetFormatifByID(c *fiber.Ctx) error {
	formatifID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	formatif, err := handler.FormatifService.GetFormatifByID(formatifID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error getting formatif", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Formatif retrieved successfully", "data": formatif})
}

// CountRelatedQuestionsHandler handles requests to get the total count of related questions for all Formatif records.
// It logs the start and end of the request, calls the FormatifService to retrieve the data, handles any errors,
// logs the successful retrieval, serializes the response to JSON, and writes the JSON response.
// CountRelatedQuestionsHandler handles requests to get the total count of related questions for all Formatif records

func (handler *FormatifHandler) CountRelatedQuestionsHandler(c *fiber.Ctx) error {
	// Log the start of the handler
	log.Println("CountRelatedQuestionsHandler started")

	// Call the service method to get the formatifs with question counts
	formatifsWithQuestionCounts, err := handler.FormatifService.CountRelatedQuestionsForAll()
	if err != nil {
		// Log the error
		log.Printf("Error in CountRelatedQuestionsHandler: %v\n", err)

		// Handle error, e.g., write an error response
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error getting formatif question counts", "data": err})
	}

	// Log the successful retrieval of formatifs with question counts
	log.Println("Successfully retrieved formatifs with question counts")

	// Serialize the formatifs with question counts to JSON and write the response
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Formatif question counts retrieved successfully", "data": formatifsWithQuestionCounts})
}

// CreateFormatif handles requests to create a new formatif.
// It parses the request body into a formatif struct.
// Returns 500 if there is an error parsing the request body.
// Calls the FormatifService to create the formatif in the database.
// Returns 500 if there is an error creating the formatif.
// Returns 200 with the created formatif on success.

func (handler *FormatifHandler) CreateFormatif(c *fiber.Ctx) error {
	rawBody := c.Body()
	log.Printf("Raw Request Body: %s\n", rawBody)

	formatif := new(model.Formatif)

	// Parse the request body into the formatif struct
	if err := c.BodyParser(formatif); err != nil {
		log.Printf("Error parsing request body: %v\n", err)
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	// Log the parsed formatif object
	log.Printf("Parsed Formatif object: %+v\n", formatif)

	// Call the repository method to create the formatif
	if err := handler.FormatifService.CreateFormatif(formatif); err != nil {
		log.Printf("Error creating formatif: %v\n", err)
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create formatif", "data": err})
	}

	// Return the created formatif as the response
	return c.JSON(fiber.Map{"status": "success", "message": "Formatif created", "data": formatif})
}

// UpdateFormatif handles requests to update an existing Formatif record.
// It parses the Formatif ID from the request parameters, the updated Formatif fields from the request body,
// calls the FormatifService to update the database, handles any errors, and returns the updated Formatif on success.

func (handler *FormatifHandler) UpdateFormatif(c *fiber.Ctx) error {
	formatifID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	formatif := new(model.Formatif)
	formatif.ID = formatifID

	// Parse the JSON body into the existing Formatif struct
	err = c.BodyParser(formatif)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	log.Printf("Updated Formatif: %+v\n", formatif)
	log.Print(formatif.UserID)

	// Update the Formatif in the database
	err = handler.FormatifService.UpdateFormatif(formatif)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't update formatif", "data": err})
	}

	// Retrieve the existing Formatif by ID
	formatif, err = handler.FormatifService.GetFormatifByID(formatifID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Formatif not found", "data": nil})
	}

	// Return the updated Formatif
	return c.JSON(fiber.Map{"status": "success", "message": "Formatif updated", "data": formatif})
}

// DeleteFormatif deletes a Formatif by ID from the database.
// It parses the formatif ID from the request parameters, calls the service
// to delete the Formatif, and returns a JSON response indicating success or failure.
// The status code will be 400 if the ID is invalid, 500 if there is a database error,
// or 200 on success.

func (handler *FormatifHandler) DeleteFormatif(c *fiber.Ctx) error {
	formatifID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	// Delete the Formatif from the database
	err = handler.FormatifService.DeleteFormatif(formatifID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't delete formatif", "data": err})
	}

	// Return success message
	return c.JSON(fiber.Map{"status": "success", "message": "Formatif deleted", "data": nil})
}
