package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"github.com/latoulicious/SIPP/internal/service"
)

type MapelHandler struct {
	MapelService *service.MapelService
}

// NewMapelHandler creates a new MapelHandler with the given MapelService.
// It returns a pointer to the MapelHandler.

func NewMapelHandler(mapelService *service.MapelService) *MapelHandler {
	return &MapelHandler{
		MapelService: mapelService,
	}
}

// GetMapel retrieves all mapels from the database and returns them in a fiber.Map.
// It returns a 500 error if there is a problem retrieving the mapels. Otherwise
// it returns a 200 status with the mapels in the response body.

func (handler *MapelHandler) GetMapel(c *fiber.Ctx) error {
	mapel, err := handler.MapelService.GetMapel()

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error getting mapel", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Mapels retrieved successfully", "data": mapel})
}

// GetMapelByID retrieves a mapel by its ID.
// It parses the ID from the route parameter "id",
// returns 400 if the ID is invalid,
// and returns the mapel in the response on success.

func (handler *MapelHandler) GetMapelByID(c *fiber.Ctx) error {
	mapelID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Mapel retrieved successfully", "data": mapelID})
}

// GetMapelPublic fetches all mapel records without requiring authentication.
// It returns the mapel records if found, else returns error.

func (handler *MapelHandler) GetMapelPublic(c *fiber.Ctx) error {
	// Implement logic to fetch all mapel without requiring JWT authentication
	mapel, err := handler.MapelService.GetMapelPublic()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error getting mapel", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Mapels retrieved successfully", "data": mapel})
}

// CreateMapel creates a new mapel from the provided input.
// It parses the input from the request body into a mapel struct.
// Returns 500 if there is an error parsing the request body or creating the mapel.
// Otherwise returns 200 with the created mapel.

func (handler *MapelHandler) CreateMapel(c *fiber.Ctx) error {
	mapel := new(model.Mapel)

	err := c.BodyParser(mapel)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	err = handler.MapelService.CreateMapel(mapel)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create mapel", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Mapel created", "data": mapel})
}

// UpdateMapel updates an existing mapel record.
// It parses the mapel ID from the route parameter "id", returns 400 if invalid.
// It fetches the existing mapel record by ID, returns 404 if not found.
// It parses the updated mapel data from the request body, returns 500 on error.
// It calls the service layer to update the mapel by ID with the updated data.
// Returns 500 if the service update fails, 200 with updated mapel on success.

func (handler *MapelHandler) UpdateMapel(c *fiber.Ctx) error {
	mapelID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	// Call the service method to get the mapel by ID
	mapel, err := handler.MapelService.GetMapelByID(mapelID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Mapel not found", "data": nil})
	}

	// Store the request body in the mapel and return an error if encountered
	err = c.BodyParser(mapel)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	// Call the service method to update the mapel
	err = handler.MapelService.UpdateMapel(mapelID, mapel)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not update mapel", "data": err})
	}

	// Return the updated mapel
	return c.JSON(fiber.Map{"status": "success", "message": "Mapel updated", "data": mapel})
}

// DeleteMapel deletes a mapel by ID.
// It first parses the mapel ID from the request parameters.
// Returns 400 if the ID is invalid.
// Calls the service layer to delete the mapel by ID.
// Returns 500 if there is an error deleting the mapel.
// Otherwise returns a 200 success message that the mapel was deleted.

func (handler *MapelHandler) DeleteMapel(c *fiber.Ctx) error {
	mapelID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	// Call the service method to delete the mapel by ID
	err = handler.MapelService.DeleteMapel(mapelID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not delete mapel", "data": err})
	}

	// Return success message
	return c.JSON(fiber.Map{"status": "success", "message": "Mapel deleted"})
}
