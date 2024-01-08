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

func NewMapelHandler(mapelService *service.MapelService) *MapelHandler {
	return &MapelHandler{MapelService: mapelService}
}

func (handler *MapelHandler) GetMapel(c *fiber.Ctx) error {
	mapel, err := handler.MapelService.GetMapel()

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error getting users", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Mapels retrieved successfully", "data": mapel})
}

func (handler *MapelHandler) GetMapelByID(c *fiber.Ctx) error {
	mapelID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Mapel retrieved successfully", "data": mapelID})
}

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
