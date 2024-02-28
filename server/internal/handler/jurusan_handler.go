package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"github.com/latoulicious/SIPP/internal/service"
)

type JurusanHandler struct {
	JurusanService *service.JurusanService
}

// NewJurusanHandler creates a new JurusanHandler with the given JurusanService.
// It is used to instantiate a JurusanHandler for handling jurusan API requests.

func NewJurusanHandler(jurusanService *service.JurusanService) *JurusanHandler {
	return &JurusanHandler{
		JurusanService: jurusanService,
	}
}

// GetJurusan retrieves all jurusan from the database
// and returns them in a fiber.Map response.
// It returns a 500 error if there is a problem retrieving the jurusan.

func (handler *JurusanHandler) GetJurusan(c *fiber.Ctx) error {
	jurusan, err := handler.JurusanService.GetJurusan()

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error getting jurusan", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Jurusan retrieved successfully", "data": jurusan})
}

// GetJurusanByID retrieves a jurusan by its ID.
// It parses the jurusan ID from the URL parameter "id".
// Returns a 400 status code if the ID is invalid.
// Returns the jurusan data in the response on success.

func (handler *JurusanHandler) GetJurusanByID(c *fiber.Ctx) error {
	jurusanID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Jurusan retrieved successfully", "data": jurusanID})
}

// GetJurusanPublic fetches all jurusan from the database without requiring JWT
// authentication. It calls the JurusanService's GetJurusanPublic method to
// retrieve the jurusan and returns them in a JSON response.
func (handler *JurusanHandler) GetJurusanPublic(c *fiber.Ctx) error {
	// Implement logic to fetch all jurusan without requiring JWT authentication
	jurusan, err := handler.JurusanService.GetJurusanPublic()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error getting jurusan", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Jurusan retrieved successfully", "data": jurusan})
}

// CreateJurusan handles creating a new jurusan.
// It parses a jurusan from the request body, generates a new UUID if one is not provided,
// calls the JurusanService to insert into the database,
// and returns the created jurusan in the response.
// Returns 500 error if there is an error parsing the request or inserting into the database.

func (handler *JurusanHandler) CreateJurusan(c *fiber.Ctx) error {
	jurusan := new(model.Jurusan)

	err := c.BodyParser(jurusan)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	// If the ID field was not included in the request body, generate a new UUID
	if jurusan.ID == uuid.Nil {
		jurusan.ID = uuid.New()
	}

	err = handler.JurusanService.CreateJurusan(jurusan)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create jurusan", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Jurusan created", "data": jurusan})
}

// UpdateJurusan updates an existing jurusan. It takes the jurusan ID from the
// request parameters, retrieves the existing jurusan by calling the service,
// updates the jurusan fields from the request body, calls the service to
// update in the database, and returns the updated jurusan. Returns a 400 error
// if the jurusan ID is invalid, 404 if not found, and 500 if there is any
// error updating.

func (handler *JurusanHandler) UpdateJurusan(c *fiber.Ctx) error {
	jurusanID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	// Call the service method to get the jurusan by ID
	jurusan, err := handler.JurusanService.GetJurusanByID(jurusanID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Jurusan not found", "data": nil})
	}

	// Store the request body in the jurusan and return an error if encountered
	err = c.BodyParser(jurusan)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	// Call the service method to update the jurusan
	err = handler.JurusanService.UpdateJurusan(jurusan)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not update jurusan", "data": err})
	}

	// Return the updated jurusan
	return c.JSON(fiber.Map{"status": "success", "message": "Jurusan updated", "data": jurusan})
}

// DeleteJurusan deletes a jurusan record from the database by ID.
// It parses the jurusan ID from the request parameters, calls the service
// layer to delete the record, and returns a 200 status code on success or
// 500 if there is an error deleting the record.

func (handler *JurusanHandler) DeleteJurusan(c *fiber.Ctx) error {
	jurusanID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	// Call the service method to delete the jurusan by ID
	err = handler.JurusanService.DeleteJurusan(jurusanID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not delete jurusan", "data": err})
	}

	// Return success message
	return c.JSON(fiber.Map{"status": "success", "message": "Jurusan deleted"})
}
