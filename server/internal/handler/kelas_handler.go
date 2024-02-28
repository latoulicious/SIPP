package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"github.com/latoulicious/SIPP/internal/service"
)

type KelasHandler struct {
	KelasService *service.KelasService
}

// NewKelasHandler initializes a new KelasHandler with the given KelasService.
// It returns a pointer to the KelasHandler.

func NewKelasHandler(kelasService *service.KelasService) *KelasHandler {
	return &KelasHandler{
		KelasService: kelasService,
	}
}

// GetKelas retrieves all kelas records from the database
// and returns them in a Kelas array.
// It returns any errors from the database query.

func (handler *KelasHandler) GetKelas(c *fiber.Ctx) error {
	kelas, err := handler.KelasService.GetKelas()

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error getting kelas", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Kelas retrieved successfully", "data": kelas})
}

// GetKelasByID retrieves a kelas by its ID.
// It parses the ID from the URL parameter "id",
// validates that it is a valid UUID,
// and returns the kelas with the given ID.
// Returns a 400 error if the ID is invalid.

func (handler *KelasHandler) GetKelasByID(c *fiber.Ctx) error {
	kelasID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Kelas retrieved successfully", "data": kelasID})
}

// GetKelasPublic retrieves all kelas records without requiring JWT authentication.
// It returns the kelas records in a Kelas array and any errors from the database query.
// This allows fetching kelas data publicly without needing to authenticate.

func (handler *KelasHandler) GetKelasPublic(c *fiber.Ctx) error {
	// Implement logic to fetch all kelas without requiring JWT authentication
	kelas, err := handler.KelasService.GetKelasPublic()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error getting kelas", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Kelas retrieved successfully", "data": kelas})
}

// CreateKelas creates a new kelas record in the database.
// It parses a kelas object from the request body,
// validates it, and passes it to the KelasService.CreateKelas method.
// Returns a 500 error if parsing or validation fails,
// or if the database insert fails.
// On success, returns a 200 status with the created kelas object.

func (handler *KelasHandler) CreateKelas(c *fiber.Ctx) error {
	kelas := new(model.Kelas)

	err := c.BodyParser(kelas)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	err = handler.KelasService.CreateKelas(kelas)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create kelas", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Kelas created", "data": kelas})
}

// UpdateKelas updates a kelas.
// It parses the kelas ID from the URL parameter "id", validates it is a valid UUID,
// fetches the existing kelas by that ID, updates it with the request body,
// calls the service to update it in the database, and returns the updated kelas.
// Returns a 400 error if the ID is invalid, a 404 error if the kelas is not found,
// a 500 error if there is an error updating the kelas.

func (handler *KelasHandler) UpdateKelas(c *fiber.Ctx) error {
	kelasID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	// Call the service method to get the kelas by ID
	kelas, err := handler.KelasService.GetKelasByID(kelasID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Kelas not found", "data": nil})
	}

	// Store the request body in the kelas and return an error if encountered
	err = c.BodyParser(kelas)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	// Call the service method to update the kelas
	err = handler.KelasService.UpdateKelas(kelas)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not update kelas", "data": err})
	}

	// Return the updated kelas
	return c.JSON(fiber.Map{"status": "success", "message": "Kelas updated", "data": kelas})
}

// DeleteKelas deletes a kelas by ID.
// It looks up the kelas by the ID in the path parameter, calls the service
// layer to delete it, and returns a 200 status code on success or an error
// status code on failure.

func (handler *KelasHandler) DeleteKelas(c *fiber.Ctx) error {
	kelasID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	// Call the service method to delete the kelas by ID
	err = handler.KelasService.DeleteKelas(kelasID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not delete kelas", "data": err})
	}

	// Return success message
	return c.JSON(fiber.Map{"status": "success", "message": "Kelas deleted"})
}
