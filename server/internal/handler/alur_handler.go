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

// NewAlurHandler creates a new AlurHandler instance with the given AlurService.
// It is used to instantiate the AlurHandler with its service dependency.

func NewAlurHandler(alurService *service.AlurService) *AlurHandler {
	return &AlurHandler{
		AlurService: alurService,
	}
}

// GetAlur retrieves the current alur from the AlurService and returns
// it in the response. It returns a 500 error if there is an error
// retrieving the alur.

func (handler *AlurHandler) GetAlur(c *fiber.Ctx) error {
	alur, err := handler.AlurService.GetAlur()

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error getting alur", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Alur retrieved successfully", "data": alur})
}

// GetAlurByID retrieves an alur by ID.
// It parses the ID from the request parameter "id", calls the AlurService
// to get the alur by that ID, and returns the alur in the response if found.
// If the ID is invalid or an error occurs getting the alur, it returns a
// 400 or 500 error respectively.

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

// CreateAlur handles creating a new alur. It parses an alur from the request body,
// passes it to the AlurService to create it, and returns the created alur in the response.
// If there are any errors parsing the request or creating the alur, it logs the error
// and returns a 400 or 500 error.

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

// UpdateAlur updates an existing alur. It takes the alur ID from the URL parameter,
// retrieves the existing alur by that ID, parses the updated alur data from the
// request body, calls the AlurService to update it in the database, and returns
// the updated alur in the response. Returns 400 if the ID is invalid, 404 if the
// alur doesn't exist, or 500 if there are any errors parsing the request or
// updating the alur.

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

// DeleteAlur deletes an Alur record from the database by ID.
// It parses the alurID from the request parameters, calls the AlurService
// to delete the record, and returns a JSON response indicating success or failure.
// Returns a 400 if the alurID is invalid, 404 if the alur is not found,
// or 500 if there is a database error.

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
