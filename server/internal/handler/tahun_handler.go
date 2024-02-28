package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"github.com/latoulicious/SIPP/internal/service"
)

type TahunHandler struct {
	TahunService *service.TahunService
}

// NewTahunHandler initializes a new TahunHandler with the given TahunService.
// It returns a pointer to the initialized TahunHandler.

func NewTahunHandler(tahunService *service.TahunService) *TahunHandler {
	return &TahunHandler{
		TahunService: tahunService,
	}
}

// GetTahun retrieves all tahun records from the database and returns them in a fiber response.
// It calls the TahunService's GetTahun method to get the tahun records, and returns a 200 status code
// with the tahun data if successful. If there is an error getting the tahun, it returns a 500 status
// code with an error message.

func (handler *TahunHandler) GetTahun(c *fiber.Ctx) error {
	tahun, err := handler.TahunService.GetTahun()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error getting tahun", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Tahun retrieved successfully", "data": tahun})
}

// GetTahunByID retrieves a tahun by its ID.
// It parses the ID from the URL parameter "id", validates it is a valid UUID,
// and returns the tahun with the matching ID if found.
// It returns a 400 status code if the ID is invalid.
// It returns a 500 status code if there is an error retrieving the tahun.
// Otherwise it returns a 200 status code with the tahun in the response body.

func (handler *TahunHandler) GetTahunByID(c *fiber.Ctx) error {
	tahunID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Tahun retrieved successfully", "data": tahunID})
}

// GetTahunPublic fetches all tahun records without requiring JWT
// authentication. It calls the TahunService's GetTahunPublic method
// to retrieve the tahun, and returns a 200 status with the tahun
// data on success, or a 500 status with an error message on failure.

func (handler *TahunHandler) GetTahunPublic(c *fiber.Ctx) error {
	// Implement logic to fetch all tahun without requiring JWT authentication
	tahun, err := handler.TahunService.GetTahunPublic()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error getting tahun", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Tahun retrieved successfully", "data": tahun})
}

// CreateTahun creates a new tahun record in the database.
// It parses the tahun data from the request body, calls the TahunService's
// CreateTahun method to insert into the database, and returns the created
// tahun in the response if successful. Returns a 500 error if there is an
// error parsing the request or inserting into the database.

func (handler *TahunHandler) CreateTahun(c *fiber.Ctx) error {
	tahun := new(model.TahunAjar)

	err := c.BodyParser(tahun)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	err = handler.TahunService.CreateTahun(tahun)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create tahun", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Tahun created", "data": tahun})
}

// UpdateTahun updates an existing tahun record in the database.
// It takes a tahun ID from the URL parameter "id", validates it is a valid UUID.
// It fetches the existing tahun record by ID using the TahunService.
// It parses the updated tahun data from the request body and validates it.
// It calls the TahunService to update the tahun in the database.
// It returns a 400 status code if the ID is invalid.
// It returns a 404 status code if the tahun is not found.
// It returns a 500 status code if there is any error updating the tahun.
// Otherwise it returns a 200 status code with the updated tahun data.

func (handler *TahunHandler) UpdateTahun(c *fiber.Ctx) error {
	tahunID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	tahun, err := handler.TahunService.GetTahunByID(tahunID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Tahun not found", "data": nil})
	}

	err = c.BodyParser(tahun)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	err = handler.TahunService.UpdateTahun(tahun)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't update tahun", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Tahun updated", "data": tahun})
}

// DeleteTahun deletes a tahun record from the database by ID.
// It parses the tahun ID from the request parameters, calls the TahunService's
// DeleteTahun method to delete from the database, and returns a 200 response
// if successful. Returns a 400 error if the ID is invalid or a 500 error if
// there is a problem deleting from the database.

func (handler *TahunHandler) DeleteTahun(c *fiber.Ctx) error {
	tahunID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	err = handler.TahunService.DeleteTahun(tahunID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't delete tahun", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Tahun deleted"})
}
