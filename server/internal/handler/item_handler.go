package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/latoulicious/SIPP/internal/model"
	"github.com/latoulicious/SIPP/internal/service"
)

type ItemHandler struct {
	ItemService *service.ItemService
}

func NewItemHandler(itemService *service.ItemService) *ItemHandler {
	return &ItemHandler{
		ItemService: itemService,
	}
}

func (handler *ItemHandler) GetItem(c *fiber.Ctx) error {
	item, err := handler.ItemService.GetItem()

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error getting item", "data": err})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Item retrieved successfully", "data": item})
}

func (handler *ItemHandler) GetItemByID(c *fiber.Ctx) error {
	itemID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	item, err := handler.ItemService.GetItemByID(itemID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Error getting item", "data": err})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Item retrieved successfully", "data": item})
}

func (handler *ItemHandler) CreateItem(c *fiber.Ctx) error {
	item := &model.ItemSoal{}

	// Log request body for debugging purposes
	log.Printf("Request Body: %+v\n", item)

	err := c.BodyParser(item)
	if err != nil {
		// Log parsing error for debugging purposes
		log.Printf("Error parsing request body: %v\n", err)
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	err = handler.ItemService.CreateItem(item)
	if err != nil {
		// Log creation error for debugging purposes
		log.Printf("Error creating item: %v\n", err)
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create item", "data": err})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Item created", "data": item})
}

func (handler *ItemHandler) UpdateItem(c *fiber.Ctx) error {
	itemID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	// Retrieve the existing Item by ID
	item, err := handler.ItemService.GetItemByID(itemID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Item not found", "data": nil})
	}

	// Parse the JSON body into the existing Item struct
	err = c.BodyParser(item)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	// Update the Item in the database
	err = handler.ItemService.UpdateItem(item)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't update item", "data": err})
	}

	// Return the updated Item
	return c.JSON(fiber.Map{"status": "success", "message": "Item updated", "data": item})
}

func (handler *ItemHandler) DeleteItem(c *fiber.Ctx) error {
	itemID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid UUID", "data": nil})
	}

	// Delete the Item from the database
	err = handler.ItemService.DeleteItem(itemID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't delete item", "data": err})
	}

	// Return success message
	return c.JSON(fiber.Map{"status": "success", "message": "Item deleted", "data": nil})
}
