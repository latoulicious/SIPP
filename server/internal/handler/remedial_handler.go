package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"github.com/latoulicious/SIPP/internal/service"
)

type RemedialHandler struct {
	RemedialService *service.RemedialService
}

func NewRemedialHandler(remedialService *service.RemedialService) *RemedialHandler {
	return &RemedialHandler{
		RemedialService: remedialService,
	}
}

func (handler *RemedialHandler) GetRemedial(c *fiber.Ctx) error {
	remedial, err := handler.RemedialService.GetRemedial()

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error getting remedial", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Remedial retrieved successfully", "data": remedial})
}

func (handler *RemedialHandler) GetRemedialByID(c *fiber.Ctx) error {
	remedialID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	remedial, err := handler.RemedialService.GetRemedialByID(remedialID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error getting remedial", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Remedial retrieved successfully", "data": remedial})
}

// CountRelatedQuestionsHandler handles requests to get the total count of related questions for all Remedial records
func (handler *RemedialHandler) CountRelatedQuestionsHandler(c *fiber.Ctx) error {
	// Log the start of the handler
	log.Println("CountRelatedQuestionsHandler started")

	// Call the service method to get the remedials with question counts
	remedialsWithQuestionCounts, err := handler.RemedialService.CountRelatedQuestionsForAll()
	if err != nil {
		// Log the error
		log.Printf("Error in CountRelatedQuestionsHandler: %v\n", err)

		// Handle error, e.g., write an error response
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error getting remedial question counts", "data": err})
	}

	// Log the successful retrieval of remedials with question counts
	log.Println("Successfully retrieved remedials with question counts")

	// Serialize the remedials with question counts to JSON and write the response
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Remedial question counts retrieved successfully", "data": remedialsWithQuestionCounts})
}

func (handler *RemedialHandler) CreateRemedial(c *fiber.Ctx) error {
	rawBody := c.Body()
	log.Printf("Raw Request Body: %s\n", rawBody)

	remedial := new(model.Remedial)

	// Parse the request body into the remedial struct
	if err := c.BodyParser(remedial); err != nil {
		log.Printf("Error parsing request body: %v\n", err)
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	// Log the parsed remedial object
	log.Printf("Parsed Remedial object: %+v\n", remedial)

	// Call the repository method to create the remedial
	if err := handler.RemedialService.CreateRemedial(remedial); err != nil {
		log.Printf("Error creating remedial: %v\n", err)
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create remedial", "data": err})
	}

	// Return the created remedial as the response
	return c.JSON(fiber.Map{"status": "success", "message": "Remedial created", "data": remedial})
}

func (handler *RemedialHandler) UpdateRemedial(c *fiber.Ctx) error {
	remedialID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	remedial := new(model.Remedial)
	remedial.ID = remedialID

	// Parse the JSON body into the existing Remedial struct
	err = c.BodyParser(remedial)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	log.Printf("Updated Remedial: %+v\n", remedial)
	log.Print(remedial.UserID)

	// Update the Remedial in the database
	err = handler.RemedialService.UpdateRemedial(remedial)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't update remedial", "data": err})
	}

	// Retrieve the existing Remedial by ID
	remedial, err = handler.RemedialService.GetRemedialByID(remedialID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Remedial not found", "data": nil})
	}

	// Return the updated Remedial
	return c.JSON(fiber.Map{"status": "success", "message": "Remedial updated", "data": remedial})
}

func (handler *RemedialHandler) DeleteRemedial(c *fiber.Ctx) error {
	remedialID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	// Delete the Remedial from the database
	err = handler.RemedialService.DeleteRemedial(remedialID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't delete remedial", "data": err})
	}

	// Return success message
	return c.JSON(fiber.Map{"status": "success", "message": "Remedial deleted", "data": nil})
}
