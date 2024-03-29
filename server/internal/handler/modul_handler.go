package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"github.com/latoulicious/SIPP/internal/service"
)

type ModulHandler struct {
	ModulService *service.ModulService
}

// NewModulHandler creates a new ModulHandler with the given ModulService.
// It is used to instantiate a ModulHandler to handle modul operations.

func NewModulHandler(modulService *service.ModulService) *ModulHandler {
	return &ModulHandler{
		ModulService: modulService,
	}
}

// GetModul retrieves details about the modul from the database
// and returns it in a Modul struct.
// It returns any errors from the database layer.

func (handler *ModulHandler) GetModul(c *fiber.Ctx) error {
	modul, err := handler.ModulService.GetModul()

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error getting modul", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Modul retrieved successfully", "data": modul})
}

// GetModulByID retrieves a modul by its ID.
// It takes a Fiber context and parses the "id" parameter as a UUID.
// It calls the ModulService to get the modul by that ID.
// If there is an error parsing the ID or getting the modul, it returns a 500 status with an error message.
// Otherwise it returns a 200 status with the retrieved modul data.

func (handler *ModulHandler) GetModulByID(c *fiber.Ctx) error {
	modulID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	modul, err := handler.ModulService.GetModulByID(modulID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error getting modul", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Modul retrieved successfully", "data": modul})
}

// CreateModul handles creating a new modul.
// It parses the request body into a Modul struct, calls the ModulService
// to insert into the database, and returns the created modul.
// It handles any errors from input validation or the database layer.

func (handler *ModulHandler) CreateModul(c *fiber.Ctx) error {
	modul := new(model.ModulAjar)

	// Log request body for debugging purposes
	log.Printf("Request Body: %+v\n", modul)

	err := c.BodyParser(modul)
	if err != nil {
		// Log parsing error for debugging purposes
		log.Printf("Error parsing request body: %v\n", err)
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	err = handler.ModulService.CreateModul(modul)
	if err != nil {
		// Log creation error for debugging purposes
		log.Printf("Error creating modul: %v\n", err)
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create modul", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Modul created", "data": modul})
}

// UpdateModul updates an existing modul in the database.
// It takes a modul ID parameter and parses it as a UUID.
// It parses the request body into the modul struct to update.
// It calls the ModulService to update the modul in the database.
// It handles any errors from parsing, updating, or retrieving the updated modul.
// Finally it returns the updated modul on success.

func (handler *ModulHandler) UpdateModul(c *fiber.Ctx) error {
	modulID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	modul := model.ModulAjar{}
	modul.ID = modulID

	// Parse the JSON body into the existing Modul struct
	err = c.BodyParser(&modul)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	log.Printf("Updated modul: %+v\n", modul)
	log.Print(modul.UserID)

	// Update the Modul in the database
	err = handler.ModulService.UpdateModul(&modul)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't update modul", "data": err})
	}

	// Retrieve the existing Capaian by ID
	modul, err = handler.ModulService.GetModulByID(modulID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Modul not found", "data": nil})
	}

	// Return the updated Modul
	return c.JSON(fiber.Map{"status": "success", "message": "Modul updated", "data": modul})
}

// DeleteModul deletes a modul by ID from the database.
// It returns a 400 status code if the ID is invalid.
// It returns a 500 status code if there is a database error.
// On success it returns a 200 status code.

func (handler *ModulHandler) DeleteModul(c *fiber.Ctx) error {
	modulID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	// Delete the Modul from the database
	err = handler.ModulService.DeleteModul(modulID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't delete modul", "data": err})
	}

	// Return success message
	return c.JSON(fiber.Map{"status": "success", "message": "Modul deleted", "data": nil})
}
