package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"github.com/latoulicious/SIPP/internal/service"
)

type CapaianHandler struct {
	CapaianService *service.CapaianService
}

// NewCapaianHandler returns a new CapaianHandler instance with the given CapaianService.
// This allows creating a CapaianHandler with a mock or configured CapaianService instance.

func NewCapaianHandler(capaianService *service.CapaianService) *CapaianHandler {
	return &CapaianHandler{
		CapaianService: capaianService,
	}
}

// GetCapaian retrieves all capaian records from the database and returns them.
// It returns a 500 error if there is a problem retrieving the records.
// Otherwise it returns a 200 status with the capaian records in the response body.

func (handler *CapaianHandler) GetCapaian(c *fiber.Ctx) error {
	capaian, err := handler.CapaianService.GetCapaian()

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error getting capaian", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Capaian retrieved successfully", "data": capaian})
}

// GetCapaianByID retrieves a Capaian by its ID.
// It parses the ID from the route parameter "id",
// validates it is a valid UUID, calls the CapaianService
// to retrieve the Capaian by that ID, and returns
// it in a JSON response with status code 200 if found.
// If there is an error parsing the ID, it returns a
// 400 error response. If there is an error retrieving
// the Capaian, it returns a 500 error response.

func (handler *CapaianHandler) GetCapaianByID(c *fiber.Ctx) error {
	capaianID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	capaian, err := handler.CapaianService.GetCapaianByID(capaianID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error getting capaian", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Capaian retrieved successfully", "data": capaian})
}

// CreateCapaian handles creating a new capaian record.
// It parses the request body into a Capaian struct, logs for debugging,
// calls the CapaianService to insert into the database, and returns
// a JSON response with the result.

func (handler *CapaianHandler) CreateCapaian(c *fiber.Ctx) error {

	capaian := new(model.Capaian)

	// Log request body for debugging purposes
	log.Printf("Request Body: %+v\n", capaian)
	log.Printf("Received Capaian object: %+v\n", capaian)

	err := c.BodyParser(capaian)
	if err != nil {
		// Log parsing error for debugging purposes
		log.Printf("Error parsing request body: %v\n", err)
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	err = handler.CapaianService.CreateCapaian(capaian)
	if err != nil {
		// Log creation error for debugging purposes
		log.Printf("Error creating capaian: %v\n", err)
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create capaian", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Capaian created", "data": capaian})
}

// UpdateCapaian handles updating an existing capaian record.
// It parses the capaian ID from the route parameter "id", validates it is a
// valid UUID, parses the updated capaian data from the request body, calls the
// CapaianService to update the database, and returns the updated capaian in the
// response. Returns a 400 error for invalid ID, 500 error for any service/db error,
// and 404 if the original capaian is not found.

func (handler *CapaianHandler) UpdateCapaian(c *fiber.Ctx) error {
	capaianID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	capaian := new(model.Capaian)
	capaian.ID = capaianID

	// Parse the JSON body into the existing Capaian struct
	err = c.BodyParser(capaian)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	log.Printf("Updated Capaian: %+v\n", capaian)
	log.Print(capaian.UserID)

	// Update the Capaian in the database
	err = handler.CapaianService.UpdateCapaian(capaian)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't update capaian", "data": err})
	}

	// Retrieve the existing Capaian by ID
	capaian, err = handler.CapaianService.GetCapaianByID(capaianID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Capaian not found", "data": nil})
	}

	// Return the updated Capaian
	return c.JSON(fiber.Map{"status": "success", "message": "Capaian updated", "data": capaian})
}

// DeleteCapaian deletes a Capaian record from the database.
// It takes a Fiber context and Capaian ID as parameters.
// It first validates the Capaian ID by parsing it as a UUID.
// If valid, it calls the CapaianService DeleteCapaian method to delete from the DB.
// It returns a 200 success response if deleted, or a 500 error if delete failed.

func (handler *CapaianHandler) DeleteCapaian(c *fiber.Ctx) error {
	capaianID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	// Delete the Capaian from the database
	err = handler.CapaianService.DeleteCapaian(capaianID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't delete capaian", "data": err})
	}

	// Return success message
	return c.JSON(fiber.Map{"status": "success", "message": "Capaian deleted", "data": nil})
}
